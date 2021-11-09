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

package platform

import (
	"github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
)

type MysqlEKS struct {
	Spec *v1alpha1.MysqlCluster
}

func (m *MysqlEKS) SetAllowedIPs(ips []string) error {
	m.Spec.Spec.MasterServiceSpec.AllowedSourceRanges = ips
	m.Spec.Spec.ReplicaServiceSpec.AllowedSourceRanges = ips
	return nil
}