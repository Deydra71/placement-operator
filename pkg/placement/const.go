/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package placement

const (
	// ServiceName -
	ServiceName = "placement"
	// ServiceAccount -
	ServiceAccount = "placement-operator-placement"
	// DatabaseName -
	DatabaseName = "placement"
	// PlacementDatabasePassword - ref name to get the placement db password from osp-secret
	PlacementDatabasePassword = "PlacementDatabasePassword"
	// PlacementPassword - ref name to get the placement password from osp-secret
	PlacementPassword = "PlacementPassword"

	// InputHashName -Name of the hash of hashes of all resources used to indentify an input change
	InputHashName = "input"

	// PlacementAdminPort -
	PlacementAdminPort int32 = 8778
	// PlacementPublicPort -
	PlacementPublicPort int32 = 8778
	// PlacementInternalPort -
	PlacementInternalPort int32 = 8778

	// AppSelector -
	AppSelector = "app"

	// KollaConfig -
	KollaConfig = "/var/lib/config-data/merged/placement-api-config.json"
	// DebugCommand -
	DebugCommand = "/usr/local/bin/kolla_set_configs && /bin/sleep infinity"
)
