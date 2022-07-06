/*
 * CloudLinux Software Inc 2019-2021 All Rights Reserved
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kuberlogicservice_env

import (
	"context"
	kuberlogiccomv1alpha1 "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/api/v1alpha1"
	config "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/cfg"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	v13 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type EnvironmentManager struct {
	client.Client

	NamespaceName string

	kls *kuberlogiccomv1alpha1.KuberLogicService
	cfg *config.Config
}

//+kubebuilder:rbac:groups=networking.k8s.io,resources=networkpolicies,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=namespaces;services;,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=resourcequotas;,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=secrets;,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=pods;,verbs=deletecollection

// SetupEnv checks if KLS environment is present and creates it if it is not
func SetupEnv(kls *kuberlogiccomv1alpha1.KuberLogicService, c client.Client, cfg *config.Config, ctx context.Context) (*EnvironmentManager, error) {
	// Namespace ns will contain all managed resources
	ns := getNamespace(kls)
	if _, err := controllerruntime.CreateOrUpdate(ctx, c, ns, func() error {
		return controllerruntime.SetControllerReference(kls, ns, c.Scheme())
	}); err != nil {
		return nil, errors.Wrap(err, "error setting up kls namespace")
	}

	// NetworkPolicy netpol prevents all cross namespace network communication for service pods
	netpol := getNetworkPolicy(kls, ns)
	if _, err := controllerruntime.CreateOrUpdate(ctx, c, netpol, func() error {
		return controllerruntime.SetControllerReference(kls, netpol, c.Scheme())
	}); err != nil {
		return nil, errors.Wrap(err, "error setting up kls networkpolicy")
	}

	// Prevent service pods from starting when requested
	resourceQuota := getResourceQuota(kls, ns)
	if _, err := controllerruntime.CreateOrUpdate(ctx, c, resourceQuota, func() error {
		resourceQuota.Spec.Hard = make(map[v1.ResourceName]resource.Quantity, 0)
		if kls.Paused() {
			resourceQuota.Spec.Hard["pods"] = resource.MustParse("0")
		}
		return controllerruntime.SetControllerReference(kls, resourceQuota, c.Scheme())
	}); err != nil {
		return nil, errors.Wrap(err, "error syncing kls resource quota")
	}

	// Sync TLS secret when defined in config
	tlsSecret := &v1.Secret{
		ObjectMeta: v12.ObjectMeta{
			Name:      cfg.SvcOpts.TLSSecretName,
			Namespace: ns.GetName(),
		},
	}

	if cfg.SvcOpts.TLSSecretName != "" {
		if _, err := controllerruntime.CreateOrUpdate(ctx, c, tlsSecret, func() error {
			srcSecret := &v1.Secret{
				ObjectMeta: v12.ObjectMeta{
					Name:      cfg.SvcOpts.TLSSecretName,
					Namespace: cfg.Namespace,
				},
			}
			if err := c.Get(ctx, client.ObjectKeyFromObject(srcSecret), srcSecret); err != nil {
				return errors.Wrap(err, "error getting source TLS secret")
			}
			tlsSecret.Data = srcSecret.Data
			return controllerruntime.SetControllerReference(kls, tlsSecret, c.Scheme())
		}); err != nil {
			return nil, errors.Wrap(err, "error syncing TLS Secret")
		}
	}

	// set namespace status field
	kls.Status.Namespace = ns.GetName()
	return &EnvironmentManager{
		Client: c,

		NamespaceName: ns.GetName(),

		cfg: cfg,
		kls: kls,
	}, nil
}

// PauseService deletes all pods in a service namespace.
// In addition to this resourceQuota in SetupEnv sets the hard limit of non-exited pods to 0.
func (e *EnvironmentManager) PauseService(ctx context.Context) error {
	if err := e.DeleteAllOf(ctx, &v1.Pod{}, &client.DeleteAllOfOptions{
		ListOptions: client.ListOptions{
			Namespace: e.NamespaceName,
		},
		DeleteOptions: client.DeleteOptions{},
	}); err != nil {
		return errors.Wrap(err, "error deleting all pods in namespace")
	}
	return nil
}

func getNamespace(kls *kuberlogiccomv1alpha1.KuberLogicService) *v1.Namespace {
	return &v1.Namespace{
		ObjectMeta: v12.ObjectMeta{
			Name:   kls.GetName(),
			Labels: envLabels(kls),
		},
	}
}

func getNetworkPolicy(kls *kuberlogiccomv1alpha1.KuberLogicService, ns *v1.Namespace) *v13.NetworkPolicy {
	return &v13.NetworkPolicy{
		ObjectMeta: v12.ObjectMeta{
			Name:      "namespace-only-egress",
			Namespace: ns.Name,
			Labels:    envLabels(kls),
		},
		Spec: v13.NetworkPolicySpec{
			PolicyTypes: []v13.PolicyType{
				v13.PolicyTypeEgress,
			},
			Egress: []v13.NetworkPolicyEgressRule{
				{
					To: []v13.NetworkPolicyPeer{
						{
							NamespaceSelector: &v12.LabelSelector{
								MatchLabels: envLabels(kls),
							},
						},
					},
				},
			},
		},
	}
}

func getResourceQuota(kls *kuberlogiccomv1alpha1.KuberLogicService, ns *v1.Namespace) *v1.ResourceQuota {
	q := &v1.ResourceQuota{
		ObjectMeta: v12.ObjectMeta{
			Name:      kls.GetName(),
			Namespace: ns.GetName(),
			Labels:    envLabels(kls),
		},
	}
	return q
}

func envLabels(kls *kuberlogiccomv1alpha1.KuberLogicService) map[string]string {
	return map[string]string{
		"kuberlogic.com/env": kls.Name,
	}
}
