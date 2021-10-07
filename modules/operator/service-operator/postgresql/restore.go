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

package postgresql

import (
	"fmt"
	kuberlogicv1 "github.com/kuberlogic/kuberlogic/modules/operator/api/v1"
	"github.com/kuberlogic/kuberlogic/modules/operator/service-operator/base"
	"github.com/kuberlogic/kuberlogic/modules/operator/util"
	v12 "k8s.io/api/core/v1"
)

const (
	restoreImage = "backup-restore-postgresql"
	restoreTag   = "latest"
)

type Restore struct {
	base.BaseRestore
	Cluster *Postgres
}

func (p *Restore) SetRestoreImage() {
	p.Image = util.GetKuberlogicImage(restoreImage, restoreTag)
}

func (p *Restore) SetRestoreEnv(cm *kuberlogicv1.KuberLogicBackupRestore) {
	pgDataSecret := fmt.Sprintf("%s.%s.credentials", postgreSuperUser,
		p.Cluster.Operator.Name)

	env := []v12.EnvVar{
		{
			Name:  "SCOPE",
			Value: p.Cluster.Operator.Name,
		},
		{
			Name:  "CLUSTER_NAME_LABEL",
			Value: postgresPodLabelKey,
		},
		{
			Name: "POD_NAMESPACE",
			ValueFrom: &v12.EnvVarSource{
				FieldRef: &v12.ObjectFieldSelector{
					APIVersion: "v1",
					FieldPath:  "metadata.namespace",
				},
			},
		},
		// Postgres env vars
		{
			Name:  "PG_VERSION",
			Value: p.Cluster.Operator.Spec.PostgresqlParam.PgVersion,
		},
		{
			Name:  "PGPORT",
			Value: "5432",
		},
		{
			Name:      "PGUSER",
			ValueFrom: util.FromSecret(pgDataSecret, "username"),
		},
		{
			Name:  "PGDATABASE",
			Value: postgreSuperUser,
		},
		{
			Name:  "PGSSLMODE",
			Value: "require",
		},
		{
			Name:      "PGPASSWORD",
			ValueFrom: util.FromSecret(pgDataSecret, passwordField),
		},
		{
			Name:  "PATH_TO_BACKUP",
			Value: cm.Spec.Backup,
		},
		{
			Name:  "DATABASE",
			Value: cm.Spec.Database,
		},
	}
	env = append(env, util.BucketVariables(cm.Spec.SecretName)...)
	env = append(env, util.S3Credentials(cm.Spec.SecretName)...)
	env = append(env, util.SentryEnv()...)
	p.EnvVar = env
}
