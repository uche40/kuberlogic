package keycloak

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-oidc"
	"github.com/kuberlogic/operator/modules/apiserver/internal/cache"
	"github.com/kuberlogic/operator/modules/apiserver/internal/logging"
	"github.com/kuberlogic/operator/modules/apiserver/internal/security/auth/policy"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type keycloakAuthProvider struct {
	oidcVerifier *oidc.IDTokenVerifier
	oauthConfig  oauth2.Config

	ctx                context.Context
	cache              cache.Cache
	permissionEnforcer policy.Enforcer

	log logging.Logger
}

type userPermissions []struct {
	Scopes []string `json:"scopes"`
	Rsid   string   `json:"rsid"`
	Rsname string   `json:"rsname"`
}

const (
	umaGrantType = "urn:ietf:params:oauth:grant-type:uma-ticket"
)

func (k *keycloakAuthProvider) GetAuthenticationSecret(username, password string) (string, error) {
	k.log.Debugf("getting authentication secret for user %s", username)

	oauth2token, err := k.oauthConfig.PasswordCredentialsToken(k.ctx, username, password)
	if err != nil {
		k.log.Errorf("error getting token for username %s: %s", username, err.Error())
		return "", fmt.Errorf("Failed to get token" + err.Error())
	}
	rawIDToken, ok := oauth2token.Extra("id_token").(string)
	if !ok {
		k.log.Debugf("no id_token found in oauth2 token: %v", oauth2token)
		return "", fmt.Errorf("No id_token found in oauth2 token")
	}
	idToken, err := k.oidcVerifier.Verify(k.ctx, rawIDToken)
	if err != nil {
		k.log.Errorf("failed to verify ID token: %s", err.Error())
		return "", fmt.Errorf("failed to verify ID token")
	}

	resp := struct {
		Oauth2token   *oauth2.Token
		IDTokenClaims *json.RawMessage
	}{oauth2token, new(json.RawMessage)}

	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		k.log.Errorf("error extracting id_token claims from id_token %v: %s", err.Error(), idToken)
		return "", fmt.Errorf("error extracting id_token claims")
	}
	return resp.Oauth2token.AccessToken, nil
}

func (k *keycloakAuthProvider) Authenticate(token string) (string, string, error) {
	k.log.Debugf("authenticating new user with token")

	p := strings.Split(token, " ")
	if len(p) != 2 {
		k.log.Errorf("error extracting authentication token from %s", token)
		return "", "", fmt.Errorf("error extracting authentication token")
	}

	idToken, err := k.oidcVerifier.Verify(k.ctx, p[1])
	if err != nil {
		k.log.Errorf("error veryfying authentication token %s", err.Error())
		return "", "", fmt.Errorf("error veryfying authentication token")
	}

	var userInfo struct {
		Username string `json:"preferred_username"`
	}
	if err := idToken.Claims(&userInfo); err != nil {
		k.log.Errorf("error getting username from authentication token: %s", err.Error())
		return "", "", fmt.Errorf("error getting username from authentication token")
	}

	if userInfo.Username == "" {
		return "", "", fmt.Errorf("empty username")
	}

	return userInfo.Username, p[1], nil
}

func (k *keycloakAuthProvider) Authorize(token, action, object string) (bool, error) {
	// check cache first
	if permissions, found := k.cache.Get(token); found {
		k.log.Debugf("permissions for action %s on object %s found in cache", action, object)
		authorized, err := k.permissionEnforcer.IsAuthorized(permissions.(policy.Permissions), token, object, action)
		return authorized, err
	}

	// get permissions from keycloak
	kPermissions, err := k.getUserPermissions(token)
	if err != nil {
		k.log.Errorf("error getting permissions from keycloak: %s", err.Error())
		return false, err
	}

	permissions := policy.Permissions{}
	for _, p := range *kPermissions {
		for _, s := range p.Scopes {
			permissions.Rules = append(permissions.Rules, policy.PermissionRule{
				Subject:  token,
				Resource: p.Rsname,
				Action:   s,
			})
		}
	}
	k.cache.Set(token, permissions, 60)

	authorized, err := k.permissionEnforcer.IsAuthorized(permissions, token, object, action)
	return authorized, err
}

func (k *keycloakAuthProvider) getUserPermissions(token string) (*userPermissions, error) {
	data := url.Values{}

	data.Set("grant_type", umaGrantType)
	data.Set("audience", k.oauthConfig.ClientID)
	data.Set("response_mode", "permissions")

	req, err := http.NewRequest("POST", k.oauthConfig.Endpoint.TokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		k.log.Errorf("error building a client for Keycloak authorization services: %s", err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		k.log.Errorf("error requesting Keycloak permissions: %s", err.Error())
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	k.log.Debugf("Keycloak authorization services response: %s", string(bodyBytes))

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting user permissions, status code is %v message is %s", res.StatusCode, string(bodyBytes))
	}

	permissions := &userPermissions{}
	if err := json.Unmarshal(bodyBytes, permissions); err != nil {
		return nil, fmt.Errorf("error unmarshalling keycloak response: " + err.Error())
	}
	return permissions, nil
}

func NewKeycloakAuthProvider(clientId, clientSecret, realmName, keycloakUrl string, cache cache.Cache, log logging.Logger) (*keycloakAuthProvider, error) {
	configUrl := fmt.Sprintf("%s/auth/realms/%s", keycloakUrl, realmName)
	ctx := context.Background()

	log.Debugf("initializing oidc provider with url %s", configUrl)
	provider, err := oidc.NewProvider(ctx, configUrl)
	if err != nil {
		return nil, fmt.Errorf("error initializing keycloak oidc config: " + err.Error())
	}

	log.Debugf("initializing oauth2 config with client_id %s", clientId)
	oauth2Config := oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  "",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	oidcConfig := &oidc.Config{
		ClientID: clientId,
	}

	// we use external rule enforcer because of limited Ketcloak enforcer
	log.Debugf("initializing permission policy enforcer")
	enforcer := policy.NewEnforcer(cache, log)

	return &keycloakAuthProvider{
		oidcVerifier: provider.Verifier(oidcConfig),
		oauthConfig:  oauth2Config,

		permissionEnforcer: enforcer,

		ctx:   ctx,
		cache: cache,
		log:   log,
	}, nil
}