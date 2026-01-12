# Changes

This document describes the relevant changes between releases of the API model.

## 0.0.441 Jan 12 2026
- Add `cloud_provider` field to CS `dns_domains` type

## 0.0.440 Dec 09 2025
- Align log forwarding config endpoints with API

## 0.0.439 Nov 25 2025
- Add patch operation on version gates

## 0.0.438 Nov 18 2025

- Add fields `LastClusterImagesetSync`, `LastHibernationCheck`, and
  `LastLimitedSupportOverrideCheck` to `Environment` resource

## 0.0.437 Nov 10 2025

- Add `LogForwarders` APIs/resources to `clusters_mgmt`
- Remove DisableUserWorkloadMonitoring field from ARO-HCP

## 0.0.436 Oct 30 2025

- Add ImageType enum to the NodePool model
- Add Preference enum field to the AwsCapacityReservation model
- Add Features sub-object to the MachineType model

## 0.0.435 Oct 21 2025

- Update deprecation message for DisableUserWorkloadMonitoring field in
  model/clusters_mgmt
- update some code comments mentioning clusters_mgmt to aro_hcp in
  model/aro_hcp

## 0.0.434 Oct 20 2025

- Split/Duplicate ARO-HCP models. From this point on, the ARO-HCP models
  become independent from the clusters_mgmt models

## 0.0.433 Sep 19 2025

- [OCM-689] Added AccountGroup and AccountGroupAssignment APIs for RBAC

## 0.0.432 Aug 28 2025

- Add resource-scoped permissions for WIF configurations

## 0.0.431 Aug 25 2025

- [OCM-17787] Add `ImageMirror` resource and type for ROSA HCP
- Add API endpoints for image mirros for ROSA HCP

## 0.0.430 Aug 19 2025

- update MarketType enum to match the camelcase values expected by CS
- deprecate os disk configuration old design
- Add API endpoints for cluster autoscaler in ARO HCP

## 0.0.427 Aug 06 2025

- Remove `osDiskSseEncryptionSetResourceId` field from type `AzureNodePoolOsDisk`
- Add methods to enable/disable backups of control planes
- Fix aws capacity reservation class to struct
- Adds the `federatedProjectId` and `federatedProjectNumber` fields to support federated project

## 0.0.426 Jul 23 2025

- no changes

## 0.0.425 Jul 23 2025

- [ARO-19918] Restructure OS disk attributes under a single section for ARO-HCP NodePool
- [OCM-15598] Add fields for federated project (WIF)
- Update metamodel version

## 0.0.424 Jul 20 2025

- [OCM-17005] Fix AWS capability handling

## 0.0.423 Jul 15 2025

- Update the model to use the latest metamodel generator
- Update metamodel generator makefile to also run `go mod tidy`
- Bump metamodel version to v0.0.71
- [ARO-19120] Cluster image registry - improve documentation
- [OCM-17005] Define capacity reservation parameters in the `AWSNodePool` type
- [ARO-19122] Remove `ClusterCapabilities` field from cluster spec

## 0.0.422 Jul 12 2025

- Fix documentation
- Bump metamodel to latest v0.0.70

## 0.0.421 Jul 10 2025

- added `ManagementCluster` field to `AWSBackupConfig` type and changed `BackupConfigs` field from map to list in `AWSShard`

## 0.0.420 Jun 25 2025

- Add `AWSBackupConfig` type
- Add `BackupConfigs` to `AWSShard`

## 0.0.419 Jun 05 2025

- Add `LogType` to `NotificationDetailsRequest`

## 0.0.418 Jun 04 2025

- Add arohcp `Version` types and endpoints

## 0.0.417 Apr 08 2025

- Add `RhRegionID` field to `Subscription` and `ClusterAuthorizationRequest` type

## 0.0.416 Mar 31 2025

- Added annotation to nodepool state

## 0.0.415 Mar 24 2025

- Added `ClusterCapabilities` type
- Added `Capabilities` field to `Cluster` type

## 0.0.414 Mar 17 2025

- Added annotations to async methods
- Added node pool status endpoint to ARO HCP service
- Added node pool endpoints to the ARO HCP service
- Update metamodel version to v0.0.67

## 0.0.413 Mar 9 2025

- Added inflight checks endpoints to the ARO HCP service
- Update metamodel version to v0.0.66

## 0.0.412 Feb 26 2025

- Added `ClusterCondition` field to `VersionGate` type

## 0.0.411 Feb 5 2025

- Update json annotations for SdnToOvnClusterMigration properties

## 0.0.410 Jan 31 2025

- Add `ClusterStatus` type to aro_hcp
- Add `ClusterStatus` resources to aro_hcp
- Update metamodel version to v0.0.65

## 0.0.409 Jan 22 2025

- Update `ClusterMigrationState` value to be a struct

## 0.0.408 Jan 21 2025

- Update SdnToOvn json value in `ClusterMigrationType` enum

## 0.0.407 Jan 17 2025

- Add `ClusterMigration` type
- Add `ClusterMigration` resources

## 0.0.406 Dec 18 2024

- Expose the GCP Shielded VM secure boot setting at the machine pool level

## 0.0.405 Dec 15 2024

- Add aro_hcp v1alpha1 Root resource

## 0.0.404 Dec 13 2024

- Add WifConfig patch endpoint

## 0.0.403 Dec 13 2024

- Add `NodesOutboundConnectivity` in azure_type.model to the `Azure` model
- Add `AzureNodesOutboundConnectivity` model

## 0.0.402 Nov 15 2024

- Add WifConfig status endpoint
- Add `WifConfigStatus` model

## 0.0.401 Oct 31 2024

- Add settings for HCP Shared VPC to `Cluster` model
- Add support for cluster arch to `DNS` model

## 0.0.400 Oct 24 2024

- Update `Parameters` attribute in addon_installation_type.model to the List of `AddonInstallationParameter` type.
- Update `Parameters` attribute in addon_type.model to the  List of `AddonParameter` type.

## 0.0.399 Oct 22 2024

- Improve Clusters Azure API type documentation
- Add Azure Operators Authentication related resources
- Add `AzureOperatorsAuthentication` resource to `Azure` resource

## 0.0.398 Oct 04 2024

- Add updates for GCP Private Service Connect feature

## 0.0.397 Sep 26 2024

- Add `FlapDetection` field to `status_board.statuses` model

## 0.0.396 Sep 24 2024

- Add `FlapDetection` field to `status_board.status_updates` model

## 0.0.395 Sep 13 2024

- Add `WifEnabled` to `Version` model
- Add `AWSShard` to the `ProvisionShard` model

## 0.0.394 Sep 10 2024

- Add `WifTemplates` field to `WifConfig` model

## 0.0.393 Sep 09 2024

- Include missing field to hypershift model
- Include new fields to manifest model
- Remove v2alpha1 resources

## 0.0.392 Sep 03 2024

- Defined "support" field in WifConfig structure
- Added "vm" WIF access method

## 0.0.391 Aug 28 2024

- Add `RegistryConfig` attribute to `Cluster` model
- Add `RegistryAllowlist` resource and endpoints

## 0.0.390 Aug 26 2024

- Add `RolePrefix` field to `WifGcp` model

## 0.0.389 Aug 13 2024

- Add state struct to node pool
- Add limited support reason override

## 0.0.388 Aug 06 2024

- Add 'OidcThumbprint' type model to v1 and v2alpha1
- Add 'OidcThumbprintInput' type model to v1 and v2alpha1
- Add 'OidcThumbprint' resource model to v1 and v2alpha1

## 0.0.387 Aug 05 2024

- Add `ProjectNumber` field to `WifConfig` model

## 0.0.386 Jul 29 2024

- Add `RootVolume` attribute to `AWSNodePool` model

## 0.0.385 Jul 29 2024

- Update WIF endpoint path
- Remove WIF templates endpoints

## 0.0.384 Jul 29 2024

- Add clusters_mgmt v2 API model

## 0.0.383 Jul 23 2024

- Add `Kind` and `Id` field to GCP `Authentication` structure

## 0.0.382 Jul 15 2024

- Add `Authentication` field to GCP model

## 0.0.381 Jul 11 2024

- Add WIF endpoints and resources

## 0.0.380 Jul 5 2024

- Add `Architecture` attribute to `MachineType` model
- Add `ReleaseImages` attribute to `Version` model

## O.0.379 Jun 28 2024

- Add `Ec2MetadataHttpTokens` attribute to `AWSNodePool` model

## O.0.378 Jun 24 2024

- Add `MultiArchEnabled` attribute to `Cluster` model

## 0.0.377 Jun 12 2024

- Change type of the `OSDiskSizeGibibytes` attribute in the `AzureNodePool` resource from String to Integer.

## 0.0.376 May 31 2024

- Add `AzureNodePool` resource to `NodePool` resource.

## 0.0.375 May 30 2024

- Add `ManagementUpgrade` parameters to the `NodePool` model to support additional upgrade configuration.
- Rename file for access_request_status_type to .model
- Add `AdditionalAllowedPrincipals` to `AWS` type to support additional allowed principal ARNs to be added to the hosted control plane's VPC Endpoint Service.

## 0.0.374 May 23 2024

- Add `CreationTimestamp` and `LastUpdateTimestamp` to `RolePolicyBinding` type
- Add `access_transparecy` Service and it's resources

## 0.0.373 May 13 2024

- Add `subnet_resource_id` to `Azure` resource
- Add `network_security_group_resource_id to`Azure` resource

## 0.0.372 May 2 2024

- Exposed `/api/clusters_mgmt/v1/clusters/{id}/kubelet_configs` for managed of KubeletConfig on HCP clusters
- Added `kubelet_configs` field to `NodePool` API Resource
- Added `name` field to `KubeletConfig` API resource

## 0.0.371 Apr 24 2024

- Add `Tags` to the `AWSMachinePool` model to support custom AWS tags for day 2 creation of machine pools

## 0.0.370 Apr 24 2024

- Add `RolePolicyBindings` to the `AWS` resource model to support STS Arbitrary Policies feature.

## 0.0.369 Apr 10 2024

- Fix spacing in description of Azure's ManagedResourceGroupName
- Fix CHANGES.md formatting

## 0.0.368 Apr 10 2024

- Add `Azure` resource to `Cluster` resource.

## 0.0.367 Apr 10 2024

- Update metamodel version to 0.0.60
- [OCM-6294] add /load_balancer_quota_values endpoint
- [OCM-7027] feat: document pagination and ordering support for break glass
- [OCM-7144] Add  /storage_quota_values endpoint

## 0.0.366 Mar 18 2024

- Fix default capabilities

## 0.0.365 Mar 17 2024

- [OCM-6763] Add default capability resource to SDK

## 0.0.364 Mar 12 2024

- Add `BreakGlassCredentials` to the `Cluster` resource model.

## 0.0.363 Mar 07 2024

- Add `NodeDrainGracePeriod` to the `NodePool` model.

## 0.0.362 Mar 06 2024

- Changed `UserName` attribute for TokenClaimMappings to `Username`.

## 0.0.361 Mar 06 2024

- Add `Scope` attribute to `ReservedResource`.
- Add `Scope` attribute to `ClusterAuthorizationRequest`.

## 0.0.360 Mar 04 2024

- Add `ComponentRoutes` attribute to `Ingress`.

## 0.0.359 Feb 28 2024

- Add `ExternalAuthConfig` resource to `Cluster`.

## 0.0.358 Feb 28 2024

- Add `DomainPrefix` to `Cluster` model.

## 0.0.357 Feb 26 2024

- [OCM-5822] Add `ExternalAuth` to `ExternalAuthConfig` model.

## 0.0.356 Feb 19 2024

- Reverting change to remove DELETE `provision_shards/{id}`

## 0.0.355 Feb 16 2024

- [OCM-5976] Removed undefined api calls

## 0.0.354 Feb 12 2024

- Add `AdditionalSecurityGroupIds` to `AWS Node Pool` type.

## 0.0.353 Feb 08 2024

- Add support for `PackageImage` in `clusters_mgmt`

## 0.0.352 Feb 05 2024

- [OSDEV-1296] Remove StatusBoard `fullname` search parameter.

## 0.0.351 Jan 30 2024

- Deprecate Notify resource

## 0.0.350 Jan 29 2024

- [OCM-5802] Add `ExternalAuthConfig` to `Cluster` model.

## 0.0.349 Jan 28 2024

- [OCM-5561] Remove redundant fields from /notify_details

## 0.0.348 Jan 26 2024

- Add `SubnetOutposts` and `AvailabilityZoneTypes` to `aws_node_pool_type` and `aws_machine_pool_type` models.

## 0.0.347 Jan 23 2024

- Add `HostedControlPlaneDefault` boolean to `Version` Type model.

## 0.0.346 Jan 11 2024

- Modify notify_details response

## 0.0.345 Jan 03 2024

- Add `validate_credentials` resource to `AwsInquiries`
  
## 0.0.344 Jan 03 2024

- [OCM-5426] Add the /notify_details endpoint to the SDK

## 0.0.343 Dec 15 2023

- Add `Platform` to `subnet_network_verification_type` resource

## 0.0.342 Dec 12 2023

- Add `Search` and `Order` methods to List `/api/clusters_mgmt/v1/clusters/{id}/node_pools`

## 0.0.341 Dec 04 2023

- Add DELETE `/api/addons_mgmt/v1/clusters/{id}/addons` endpoint

## 0.0.340 Dec 01 2023

- Add `Platform` type
- Add `Platform` to `network_verification_type` resource

## 0.0.339 November 14 2023

- Add GCP inquiries resource machinetypes

## 0.0.338 November 08 2023

- Add `ProductTechnologyPreviews` and `ProductMinimalVersions` endpoints

## 0.0.337 November 08 2023

- Updated `post` method definition on `KubeletConfig` resource for consistency with `update`

## 0.0.336 November 07 2023

- Added `security` field to Cluster Service GCP structure.

## 0.0.335 November 03 2023

- Add `doc_references` field in `LogEntry`

## 0.0.334 November 02 2023

- Add `Search` method to `status_board` `status_updates` model

## 0.0.333 November 02 2023

- Added `/api/clusters_mgmt/v1/clusters/{id}/kubelet_config` endpoint
- Added `KubeletConfig` struct to support requests to the `kubelet_config` endpoint
- Updated `Cluster` struct to be able to have the embedded `KubeletConfig`

## 0.0.332 November 01 2023

- Add `AdditionalInfraSecurityGroupIds` to `AWS` type
- Add `AdditionalControlPlaneSecurityGroupIds` to `AWS` type

## 0.0.331 October 27 2023

- Add `Search` method to `status_board` `products_resource`, `applications_resource`, and `services_resource` models

## 0.0.330 October 26 2023

- Add `Update` method to `HypershiftConfig` resource

## 0.0.329 October 23 2023

- Add get `ClusterId` to `network_verification_type` resource

## 0.0.328 October 19 2023

- Add get `VPC` to `Cluster` resource

## 0.0.327 October 18 2023

- Add `BestEffort` to method `Delete` in `Cluster` resource

## 0.0.326 October 13 2023

- Add `BackplaneURL` to `Environment` type

## 0.0.325 September 28 2023

- Add `OrganizationId` to `feature_review_request` type

## 0.0.324 September 26 2023

- Add `CreatedAt` to `cluster_log` type
- Add `CreatedBy` to `cluster_log` type

## 0.0.323 September 25 2023

- Add `GCPMarketplaceEnabled` to `version` type

## 0.0.322 September 25 2023

- Add `AdditionalComputeSecurityGroupIds` to `AWS` type
- Add `AdditionalSecurityGroupIds` to `AWS Machine Pool` type
- Add `AwsSecurityGroups` to `VPC` type

## 0.0.321 September 21 2023

- Revert the addition of `BackplaneURL` to `Environment` type

## 0.0.320 September 20 2023

- Exposed the `/api/clusters_mgmt/v1/aws_inquiries/sts_account_roles` endpoint
- Added the `AWSSTSAccountRole` type
- Added the `AWSSTSRole` type

## 0.0.319 September 20 2023

- Add `BackplaneURL` to `Environment` type

## 0.0.318 September 11 2023

- Add `ImageOverrides` to `Version` type

## 0.0.317 September 11 2023

- Add new resources and type for `TrustedIps`

## 0.0.316 September 6 2023

- Add `CIDRBlockp` field to `Subnetwork` and `CloudVPC`
- Add `ProvisionShardTopology` field to `ServerConfig`

## 0.0.315 September 5 2023

- Add DisplayName and Description properties to `BillingModelItem`

## 0.0.314 August 29 2023

- Revert "adding name field to Account type"

## 0.0.313 August 29 2023

- Add new resources and a type for `BillingModelItem`

## 0.0.312 August 25 2023

- Updated existing API's and added support for `Addon Installations` endpoints

## 0.0.311 August 21 2023

- Add a new resource to OSL clusters/cluster_logs

## 0.0.310 August 17 2023

- Modify SelfAccessReview to return IsOCMInternal field

## 0.0.309 August 13 2023

- Add the remainder of autoscaler params (#813)

## 0.0.308 August 08 2023

- Modify access review response to include is_ocm_internal field
- gcp: Add support for Shared VPC

## 0.0.307 August 01 2023

- Move `PrivateHostedZoneID` and `PrivateHostedZoneRoleARN` to `aws_type` resource

## 0.0.306 July 27 2023

- Fix upgrade related constants JSON output to align with existing values

## 0.0.305 July 26 2023

- Add `PrivateHostedZoneID` and `PrivateHostedZoneRoleARN` to `cluster_type` resource

## 0.0.304 July 26 2023

- Add upgrade related constants also for `NodePoolUpgradePolicy`.
- Change DNS domain field names.

## 0.0.303 July 24 2023

- Add upgrade related constants.
- Fix CHANGES.md formatting.

## 0.0.302 July 20 2023

- Add property `MarketplaceGCP` to `billing_model_type` in `clusters_mgmt` and `accounts_mgmt`
- Document `GovCloud`, `KMSLocationID` and `KMSLocationName` fields to `CloudRegion`
- Document `fetchRegions=true` to `cloud_providers`

## 0.0.301 July 12 2023

- Update name for `ClusterStsSupportRole` resource and type to `StsSupportJumpRole`

## 0.0.300 July 11 2023

- Add `UserDefined` in dns domain resource

## 0.0.299 July 11 2023

- Add autoscaler locator in cluster resource (#789)

## 0.0.298 July 9 2023

- Add cluster autoscaler API resources (#781)

## 0.0.297 July 6th 2023

- Add managed ingress attributes
- Fix `fetchLabels` and `fetchAccounts` url parameter names
- Add `ClusterStsSupportRole` resource and type

## 0.0.296 June 28 2023

- Add json annotation to `DeleteAssociatedResources` locator in account resource

## 0.0.295 June 26 2023

- Update `ReserveAt` to `ReserveAtTimestamp` in dns domain type

## 0.0.294 June 26 2023

- Add `DeleteAssociatedResources` locator to account resource

## 0.0.293 June 21 2023

- Add label list to OSDFM cluster request payloads
- Replace references to labels in OSDFM cluster structs with the labels themselves
- Fix typos in OSDFM cluster Label struct fields
- Add HashedPassword field to clusters_mgmt to provide encrypted value

## 0.0.292 June 21 2023

- Add cluster autoscaler structs (#747)

## 0.0.291 June 15 2023

- Add Reason to access review responses
- Enable users to provide both hashed and plain-text passwords
- API model for network verification

## 0.0.290 June 12 2023

- Rename `MachineTypeRootVolume` to `RootVolume`
- Put `RootVolume` in `ClusterNodes`
- add contracts to cloud accounts (#765)

## 0.0.289 June 7 2023

- Add Load balancer type to Ingress model
- remove unused API endpoints

## 0.0.288 June 5 2023

- Complete OSD FM api for SDK usage

## 0.0.287 June 1 2023

- Add `Htpasswd` to `Cluster`.

## 0.0.286 June 1 2023

- Add `MachineTypeRootVolume` to `MachinePool`.

## 0.0.285 May 25 2023

- Changed DNS Domain from Class to a Struct.
- Change dns domain type to class and remove ID.

## 0.0.284 May 24 2023

- Add 'add, update & delete' to cloud region/regions resource model

## 0.0.283 May 18 2023

- Fix Ec2MetadataHttpTokens field OpenAPI formatting
- Add node pool upgrade structs and resources

## 0.0.282 May 18 2023

- Rename HttpTokensState to Ec2MetadataHttpTokens

## 0.0.281 May 15 2023

- Add `RootVolume` of type `MachineTypeRootVolume` to `MachineType` type.

## 0.0.280 May 4 2023

- Add `HttpTokensState` to `AWS` resource.

## 0.0.279 Apr 26 2023

- Add `AuditLog` to `AWS` resource.
- Add `RoleArn` attribute to the `AuditLog` model.

## 0.0.278 Apr 23 2023

- Add InflightChecks locator to cluster resource

## 0.0.277 Apr 20 2023

- Add `BillingAccountID` to AWS model

## 0.0.276 Apr 18 2023

- Add delete method to `Account` resource.
- Add `tuning_configs` endpoints.
- Add `tuning_configs` field to Node Pools.

## 0.0.275 Apr 17 2023

- Add pending delete clusters API.

## 0.0.274 Apr 16 2023

- Add `Subnets` property to the CloudProviderData model.

## 0.0.273 Apr 14 2023

- update metamodel version 0.0.57
- remove circular dependencies from clusters mgmt

## 0.0.272 Apr 3 2023

- adding quota version
-

## 0.0.271 Apr 3 2023

- Adding `version_inquiry` endpoint to Managed Services.

## 0.0.270 Mar 31 2023

- adding quota auth to root resource model

## 0.0.269 Mar 30 2023

- Add `DeleteProtection` resource to `Cluster` resource.
- adding quota auth models

## 0.0.268 Mar 24 2023

- Replace `OidcConfigId` for `OidcConfig` in `STS` resource.

## 0.0.267 Mar 24 2023

- Add `OidcConfigId` to `STS` resource.
- Remove `OidcPrivateKeySecretArn` from `STS` resource.

## 0.0.266 Mar 17 2023

- Adjust `Oidc Configs` endpoints.

## 0.0.265 Mar 14 2023

- Rename `HypershiftEnabled` boolean to `HostedControlPlaneEnabled` in `Version` Type model.

## 0.0.264 Mar 10 2023

- Add `Hosted Oidc Configs` endpoints.

## 0.0.263 Mar 7 2023

- Add `HypershiftEnabled` boolean to `Version` Type model.

## 0.0.262 Mar 6 2023

- Add `Control Plane Upgrade Scheduler` endpoints.

## 0.0.261 Feb 22 2023

- Add `CommonAnnotations` and `CommonLabels` filed to Addon Type.

# 0.0.260 Feb 21 2023

- Add `ManagedPolicies` field to the `STS` type model.

# 0.0.259 Feb 21 2023

- Add master machine type and infra machine type to cluster nodes type

# 0.0.258 Feb 20 2023

- Export cluster name for mgmt, mgmt_parent, and svc clusters

# 0.0.257 Feb 17 2023

- Add `ByoOidc` type to Cluster type model
- Add addon upgrade policy to clusters_mgmt
- Add `Labels` and `Taints` to NodePool type

# 0.0.256 Feb 9 2023

- Add `LogType` field to Cluster Log type model
- Fix Addon status type and value constants

# 0.0.255 Feb 5 2023

- Add `Version` field to node pool

# 0.0.254 Feb 2 2023

- Add `PrivateLinkConfiguration` type with related endpoints

# 0.0.253 Jan 30 2023

- Update `Permission` resource attributes
  - Rename `ResourceType` to `Resource`

## 0.0.252 Jan 20 2023

- Update `STS` resource attributes
  - Remove `BoundServiceAccountSigningKey`
  - Remove `BoundServiceAccountKeyKmsId`
  - Rename `BoundServiceAccountKeySecretArn` to `OidcPrivateKeySecretArn`

## 0.0.251 Jan 20 2023

- Update `NodePool` with status attributes
- Added `current_compute` attribute in `ClusterStatus` for hosted clusters.
- Added missing variable to `addon environment variable` for addons mgmt

## 0.0.250 Jan 17 2023

- Add `Addon Inquiries API` to `addons_mgmt`.

## 0.0.249 Jan 13 2023

- Add `BoundServiceAccountKeySecretArn` attribute to the `Sts` model.

## 0.0.248 Jan 12 2023

- Add `AwsEtcdEncryption` type model and reference from `AWS`.
- Add `Enabled` attribute to `STS` model.

## 0.0.247 Jan 06 2023

- Corrected `Metrics` type on `DeletedSubscription`

## 0.0.246 Jan 04 2023

- Add Search to `Capabilities` resource

## 0.0.245 Dec 29 2022

- Add `BoundServiceAccountKeyKmsId` attribute to the `Sts` model.

## 0.0.244 Dec 25 2022

- Add `ARN` attribute to the `AWSSTSPolicy` model.

## 0.0.243 Dec 22 2022

- Add `BoundServiceAccountSigningKey` attribute to the `Sts` model.

## 0.0.242 Dec 21 2022

- Add `AddonNamespace` resource model.
- Add `CommonLabels` attribute to the `Addon` model.
- Add `CommonAnnotations` attribute to the `Addon` model.
- Add `MachineType` locator on `MachineTypes` model.

## 0.0.241 Dec 13 2022

- Add `AddonCluster` resource model.
- Add `AddonClusters` resource model.
- Add `AddonStatusConditionType` model.
- Add `AddonStatusConditionType` type model.
- Add `AddonStatusConditionValue` type model.
- Add `AddonStatus` resource model.
- Add `AddonStatuses` type model.
- Add `AddonStatuses` resource model.
- Add `Clusters` to root resource model.
- Add `DeletedSubscription` resource model.
- Add `DeletedSubscriptions` to root resource model.

## 0.0.240 Nov 30 2022

- Fix `AddonConfig` on `AddonConfigType` resource model.

## 0.0.239 Nov 30 2022

- Fix `NodePoolAutoScaling` resource model.
- Fix `AWSNodePool` field mapping.

## 0.0.238 Nov 28 2022

- Fix `AWSNodePool` on `NodePool` resource model.

## 0.0.237 Nov 27 2022

- Add `NodePools` resource model.
- Add `NodePools` locator on `Cluster` model.

## 0.0.236 Nov 23 2022

- Add extra fields to label model:
  - Type
  - ManagedBy
  - AccountID
  - SubscriptionID
  - OrganizationID

## 0.0.235 Nov 21 2022

- Add `Capabilities` locator in root resource model.

## 0.0.234 Nov 17 2022

- Add `Capabilities` resource model.

## 0.0.233 Nov 17 2022

- Add `SupportsHypershift` property to CloudRegion model.

## 0.0.232 Nov 07 2022

- Modify `availabilityZone` property in CloudProviderData model from `string` to `[]string`.

## 0.0.231 Nov 03 2022

- Add `AvailabilityZone` property to CloudProviderData model.
- Add `Public` property to Subnetwork model.

## 0.0.230 Nov 01 2022

- Add creation timestamp and modification timestamp to provision shard
- Add pull secret for addon version
- Add addon secret props for addon version config
- Add additional catalog sources for addon version
- Add addon parameter condition

## 0.0.229 Oct 24 2022

- Add Addon Management models
- Add GCP Encryption Keys to cluster model

## 0.0.228 Oct 14 2022

- Add marketplace specific enum for clusters mgmt
- Add Search method to ProvisionShards

## 0.0.227 Sep 28 2022

- Added Manifests to external_configuration.

## 0.0.226 Sep 21 2022

- Added expiry setting to managed service clusters.

## 0.0.225 Sep 20 2022

- [Hypershift] Expose /manifests

## 0.0.224 Sep 16 2022

- Add hypershift endpoint with its ManagementCluster.
- Align hypershift case usage.

## 0.0.223 Sep 12 2022

- Add `Version` property to CloudProviderData model.
- Add `InfraID` property to Cluster model.

## 0.0.222 Aug 29 2022

- Drop deprecated `DisplayName` property from ClusterRegistration model.

## 0.0.221 Aug 29 2022

- Add `ConsoleUrl` and `DisplayName` properties to ClusterRegistration model and correct documentation.

## 0.0.220 Aug 19 2022

- Add `ManagedBy` property to RoleBinding model.

## 0.0.219 Aug 17 2022

- Add Billing Model to Addon Installation

## 0.0.218 Aug 16 2022

- Change provision shard to include kube client configurations as well as server URL.

## 0.0.217 Aug 11 2022

- Change provision shard to include kube client configurations.
- Add GCP volume size to flavour API.

## 0.0.216 Aug 9 2022

- Add fleet manager related structures and API.

## 0.0.215 Aug 04 2022

- Add hypershift config to provision shard.

## 0.0.214 Aug 02 2022

- Add locator `label` to Generic Labels resource.

## 0.0.213 Aug 02 2022

- Add update function to provision shard API.

## 0.0.212 Aug 01 2022

- Add status to provision shard API.

## 0.0.211 Jul 27 2022

- Add API for adding and removing a provision shard.

## 0.0.210 Jul 25 2022

- Remove `DisplayName` field from Cluster model.

## 0.0.209 Jul 11 2022

- Add `Capabilities` property to AccountType model.

## 0.0.208 Jul 11 2022

- Add `Delete` method to RegistryCredentialResource model.

## 0.0.207 Jun 27 2022

- Add `Subnets` field to MachinePool type.

## 0.0.206 Jun 09 2022

- Add `ExcludeSubscriptionStatuses` field to ResourceReview type.

## 0.0.205 Jun 03 2022

- Add `BillingMarketplaceAccount` field to ReservedResource type.

## 0.0.204 Jun 02 2022

- Remove volume type from flavour
- Add Network Configuration for Managed Services

## 0.0.203 May 31 2022

- Add `MarketplaceAWS`, `MarketplaceAzure`, `MarketplaceRHM` billing models.

## 0.0.202 May 26 2022

- Add `CloudAccount` type.
- Add `CloudAccounts` field to QuotaCost type.
- Add `BillingMarketplaceAccount` field to Subscription type.

## 0.0.201 May 26 2022

- Adding groups claim to openID IDP

## 0.0.200 May 18 2022

- Add `hypershift.enabled` field to cluster type.

## 0.0.199 May 16 2022

- Fix cred request api model parameters

## 0.0.198 May 12 2022

- Add cred request to api model
- Add AWSRegionMachineTypes endpoint to api model

## 0.0.197 May 4 2022

- Change inflight check type Details field to Interface

## 0.0.196 Apr 20 2022

- Added Machine Pool Security Group Filters for Machine Pools and Cluster Nodes
- Drop RoleARN from AddOnInstallation

## 0.0.195 Apr 19 2022

- Added Import method to the HTPasswd IDP user collection.

## 0.0.194 Apr 14 2022

- Added availability zone fields to managed service cluster struct.

## 0.0.193 Apr 14 2022

- Add LimitedSupportReasonCount to cluster status struct.
- Add FullNames and LimitScope param fields for status-board status updates.
- Add Update method for `service_mgmt`

## 0.0.192 Apr 12 2022

- Add inflight checks and 'validating' state.
- Add STS missing fields.

## 0.0.191 Apr 11 2022

- Fix JSON representation of log severity.

## 0.0.190 Apr 8 2022

- Fix JSON names of identity provider types.
- Add enable minor version upgrades flag to upgrade policy.

## 0.0.189 Apr 5 2022

- Added QuotaRules to ocm-sdk-go
- added no_proxy field to the proxy project
- Added errors resource.
- Added errors support for status-board.

## 0.0.188 Apr 1 2022

- Add Status query param for incidents resource.

## 0.0.187 Mar 31 2022

- Add new `web_rca` service.

## 0.0.186 mar 30 2022

- Add ManagementCluster to ProvisionShard.

## 0.0.185 mar 30 2022

- Fixes for `CloudResources` resource.

## 0.0.184 mar 29 2022

- Added field for parameters to be specified for managed services.
- Added CloudResource type and resources.

## 0.0.183 Mar 17 2022

- Added field for parameters to be specified for managed services.

## 0.0.182 Mar 15 2022

- Adding `service_mgmt` service.

## 0.0.181 Mar 14 2022

- add aws sts policy
- Add ReleaseImage to Version

## 0.0.180 Mar 9 2022

- Add CloudProvider info to ProvisionShard

## 0.0.179 Mar 9 2022

- Fix cluster logs URL, should be `cluster_logs` instead of `cluster_logs_uuid`.

## 0.0.178 Mar 7 2022

- Add `managed_service` field to add-on type.
- Add `credentials_secret` field to add-on type.
- Add `region` field to provision shard.

## 0.0.177 Mar 3 2022

- Fix update method of environment endpoint, should be `Update` instead of
  `Patch`.

- Remove unimplemented `POST /api/service_logs/v1/cluster_logs/clusters/{uuid}/cluster_logs`
  method.

## 0.0.176 Mar 02 2022

- Add environment endpoint get and patch

## 0.0.175 Mar 01 2022

- Add new apis for addon config attribute
- Add list of requirements to addon parameter options
- Add Name fields to VPCs and Subnetworks
- Rename addon config env object

## 0.0.174 Feb 16 2022

- Add rhit_account_id to Account Class

## 0.0.173 Feb 11 2022

- addons: Support attributes necessary for STS.
- Add ProductIds param to Status Resource.
- Add Role bindings to Subscription.

## 0.0.172 Feb 4 2022

- Remove deprecated quota summary resource and type.
- Add QuotaVersion to ClusterAuth.
- Allow adding/removing operator roles.

## 0.0.171 Feb 03 2022

- Remove deprecated `SKUs` endpoint

## 0.0.170 Jan 28 2022

- Add `ServiceInfo` type status board service.

## 0.0.169 Jan 25 2022

- Version gate type: Add warning message field

## 0.0.168 Jan 10 2022

- Fix description of various API attributes
- OVN: Add network type selection
- adding field to hold validation error message

## 0.0.167 Dec 28 2021

- Change field name in version gate agreement link to version gate.

## 0.0.166 Dec 26 2021

- Change version gate agreement URL.

## 0.0.165 Dec 22 2021

- Add version raw id prefix to version gate.

## 0.0.164 Dec 22 2021

- Move version gates to be top level resource.

## 0.0.163 Dec 21 2021

- Add support for version gate deletion.

## 0.0.162 Dec 21 2021

- Add support for version gate creation.

## 0.0.161 Dec 20 2021

- Add support for version gate agreements.

## 0.0.160 Dec 20 2021

- Change version gates URL.

## 0.0.159 Dec 20 2021

- Add version gates support.

## 0.0.158 Dec 19 2021

- Adding subnetworks to vpc inquiry
- Add statuses path to service model, add some comments.
- [SDB-2509] Update OSL API schema to be compatible with ocm-sdk-go

## 0.0.157 Dec 10 2021

- Add support for multi-user `HTPasswd` IDP.

## 0.0.156 Dec 10 2021

- Add `updates` method to status board product resource.
- Fix status get method of status board.

## 0.0.155 Dec 8 2021

- Add `status_board` service.

## 0.0.154 Dec 3 2021

This version doesn't contain changes to the model, only to the development and
build workflows:

- Rename `master` branch to `main`.

  To adapt your local repository to the new branch name run the following
  commands:

  ```shell
  git branch -m master main
  git fetch origin
  git branch -u origin/main main
  git remote set-head origin -a
  ```

- Automatically add changes from `CHANGES.md` to release descriptions.

## 0.0.153 Nov 22 2021

- Enable FIPS mode

## 0.0.152 Nov 18 2021

- Update type `resource` to `clusterResources`
- Revert "Add Name field to LDAP identity provider"
- Remove addon install mode `singleNamespace`
- Add addon install mode `ownNamespace`
- Add channel to addon version class

## 0.0.151 Nov 7 2021

- Add Name field to LDAP identity provider

## 0.0.150 Oct 27 2021

- Fix addon installation version (addon_version vs version)
- Remove no_proxy attribute from SDK
- Add body to the external tracking event

## 0.0.149 Oct 11 2021

- Add AddOn Versions

## 0.0.148 Oct 05 2021

- Revert archived cluster endpoint

## 0.0.147 Sep 27 2021

- Add missing connection to clusters collection in the `service_logs` service.

## 0.0.146 Sep 21 2021

- Add Status to AddOnRequirement

## 0.0.145 Sep 13 2021

- Add Archived cluster endpoint
- Add cluster waiting state

## 0.0.144 Sep 12 2021

- Add cluster-wide proxy and gcp network attributes

## 0.0.143 Sep 9 2021

- Add Add() method to limited support reasons resource

## 0.0.142 Sep 9 2021

- Add Limited Support Reason and Templates API

## 0.0.141 Aug 18 2021

- Add EndOfLifeTimestamp for versions

## 0.0.140 Aug 11 2021

- Add check_optional_terms to TermsReview and SelfTermsReview
- Add reduceClusterList to ResourceReview and SelfResourceReview

## 0.0.139 Aug 3 2021

- Add kms key arn to aws ccs cluster encryption

## 0.0.138 Jul 22 2021

- Add addon inquiries endpoints get and list

## 0.0.137 Jul 19 2021

- Add cluster resource type
- Remove GET cluster deployment endpoint

## 0.0.136 Jul 19 2021

- Added Spot Market Options to MachinePool

## 0.0.135 Jul 6 2021

- Added username to service log LogEntry

## 0.0.134 Jun 22 2021

- Added InternalOnly flag to SubscriptionNotify

## 0.0.133 Jun 16 2021

- Added capabilities support to Organization

## 0.0.132 Jun 11 2021

- Add ccs_only flag for cloud regions

## 0.0.131 Jun 10 2021

- Add Feature Review and Self Feature Review Locators

## 0.0.130 Jun 10 2021

- Add Disable workload monitoring Field to Cluster Type
- Add Authorizations Feature Review and Self Feature Review

## 0.0.129 Jun 1 2021

- Add cloud provider inquiries to api-model
- sts: Add support role ARN

## 0.0.128 May 31 2021

- Add CCSOnly, GenericName fields to machine type.
- Add AcceleratedComputing value to MachineTypeCategory enum.

## 0.0.127 May 31 2021

- Remove the Dashboards resource.

## 0.0.126 May 27 2021

- Add ClusterConfigurationMode type under ClusterStatus
- sts: Change custom roles to instance roles

## 0.0.125 May 18 2021

- Added Custom IAM roles For STS

## 0.0.124 May 13 2021

- JobQueue: Updated attributes of Push

## 0.0.123 May 12 2021

- [SDB-2062] remove fields on ResourceQuota

## 0.0.122 May 12 2021

- JobQueue: Attribute Arguments added to Pop

## 0.0.121 May 4 2021

- JobQueue: Attributes to Pop changed

## 0.0.120 May 4 2021

- JobQueue: Attributes to Push changed

## 0.0.119 May 3 2021

- STS: Support attributes to allow STS clusters

## 0.0.118 May 2 2021

- Add 'hibernating', "powering-down' and 'resuming' cluster states.

## 0.0.117 Apr 27 2021

- JobQueue: ReceiptId is String and not Integer

## 0.0.116 Apr 27 2021

- Add a new service JobQueue

## 0.0.115 Apr 13 2021

- Add event_code and site_code to TermsReviewRequest type
- Add new SelfTermsReviewRequest type

## 0.0.114 Apr 6 2021

- related-resources: Add resource type and cloud provider
- event: Track ad-hoc authenticated events

## 0.0.113 Apr 6 2021

- Add RelatedResources struct to QuotaCost

## 0.0.112 Mar 30 2021

- Add Options to AddOnParameter type.
- aws: Support PrivateLink for fully-private clusters

## 0.0.111 Mar 17 2021

- Add subscription metrics.
- Add `deprovision` and `force` parameters to delete cluster method.
- Ensure all subscription fields are available.

## 0.0.110 Feb 22 2021

- organization: Add quota_cost endpoint resources

## 0.0.109 Feb 22 2021

- Remove deprecated upgrade channel group field

## 0.0.108 Feb 16 2021

- Add `billing_model` attribute to the `ReservedResource` type.
- Add `cluster_billing_model` attribute to the `Subscriptioin` type.

## 0.0.107 Feb 15 2021

- add addon sub operator type

## 0.0.106 Feb 8 2021

- Add billing_model field to cluster type
- subscriptions: Add label locator
- Update metamodel to v0.0.36

## 0.0.105 Feb 2 2021

- Add cluster hibernation support

## 0.0.104 Jan 27 2021

- Add addon requirement type.

## 0.0.103 Jan 26 2021

- Remove `cluster_admin_enabled` attribute from cluster type.
- Add missing subscription, cluster authorization and plan attributes.

## 0.0.102 Dec 17 2020

- add default value to add-on parameter type
- Add upgrade channel group for a cluster

## 0.0.101 Dec 2 2020

- Fix add-on installation delete endpoint

## 0.0.100 Nov 25 2020

- Remove node drain grace period from upgrade policy
- Add node drain grace period to the cluster
- Add etcd_encryption to sdk

## 0.0.99 Nov 16 2020

- Add deletion add-on installation endpoint
- Add Update method to addon installation resource

## 0.0.98 Nov 10 2020

- Change Taints to struct.

## 0.0.97 Nov 10 2020

- Remove BYOC flag from Cluster type.
- Add Taints field to MachinePool type.

## 0.0.96 Nov 2 2020

- Add Enabled to AddOnParameter type.

## 0.0.95 Oct 27 2020

- Add SubnetIDs to the AWS model.

## 0.0.94 Oct 26 2020

- [AMS] Add IncludeRedHatAssociates to SubscriptionNotify

## 0.0.93 Oct 26 2020

- version: Rename field from MOA to ROSA

## 0.0.92 Oct 21 2020

- Add RawID field to version type.

## 0.0.91 Oct 14 2020

- Remove redudant fields
- flavours: Remove infra and compute nodes
- Add AddOnParameter modal type Update AddOn to include list of AddOnParameters
- Add AddOnInstallationParameter modal type Update AddOnInstallation to include list of AddOnInstallationParameters

## 0.0.90 Oct 11 2020

- Add ComputeLabels attribute to ClusterNodes

## 0.0.89 Oct 8 2020

- Add machine pool locator to cluster resource

## 0.0.88 Oct 5 2020

- Add missing machine pools resource

## 0.0.87 Oct 5 2020

- Add missing machine pool resource

## 0.0.86 Oct 5 2020

- Added New Error Message implementation
- idp: Add HTPasswd provider
- Uptdating SDK with GCP credentials

## 0.0.85 Oct 5 2020

- Add upgrade policy state

## 0.0.84 Oct 4 2020

- Add machine pools link and type

## 0.0.83 Sep 24 2020

- add external resources to add on type model
- SDA-2952 - Add "hidden" option to AddOn

## 0.0.82 Sep 21 2020

- Added Install Error Details From Provisioner

## 0.0.81 Sep 14 2020

- Remove redundant ID from upgrade policy class
- Add key to label_type

## 0.0.80 Sep 14 2020

- Add upgrade policy type and resource
- Add terms review and self terms review
- Add dashboards summary

## 0.0.79 Sep 7 2020

- Add 'available_upgrades' list to version type
- Add CCS type and Attribute to Cluster type

## 0.0.78 Sep 4 2020

- Added New DNS_READY
- version: Add moa_enabled flag

## 0.0.77 Aug 23 2020

- Update to metamodel v0.0.32
- Add the ChannelGroup attribute to the Version model
- Add Available AWS regions command

## 0.0.76 Aug 13 2020

- Add link to missing provision shard

## 0.0.75 Aug 6 2020

- Added TokenAuthorization to root_resource
- Added SupportCase resource

## 0.0.74 Aug 5 2020

- [CS] Add hive_config to the provision shard
- [CS] Improving cluster logs endpoint
- [AMS] Added token authorization endpoint

## 0.0.73 Aug 3 2020

- Add capability_review endpoint
- Add support_cases endpoint

## 0.0.72 Jul 30 2020

- Fix comment
- Expose if a region supports multi AZ
- Add Update Identity Provider
- removing 'deprovision' suffix from logs endpoint
- add post method to subscription resource
- Add labels field to external configuration type
- Implement Batch Patch Ingresses API endpoint

## 0.0.71 Jul 21 2020

- Add API for getting cluster's provision shard
- Add API for getting provision shards

## 0.0.70 Jul 14 2020

- Add API for clusterdeployment labels
- add organization_id to cluster_registration
- label: Fix erroneous file extensions
- MachineType: Expose instance size enum

## 0.0.69 Jul 05 2020

- Added top level sku_rules endpoint to AMS

## 0.0.68 Jul 05 2020

- [AMS] Changed feature toggle API to /query with payload containing organization id

## 0.0.67 Jul 01 2020

- [AMS] Added SkuCount to ResourceQuota type

## 0.0.66 Jun 30 2020

- Change feature toggle query API to receive organization ID by POST

## 0.0.65 Jun 29 2020

- Added Syncsets API
- Added Uninstall Log
- Update to metamodel v0.0.30

## 0.0.64 Jun 21 2020

- Added Notify to root_resource in AMS

## 0.0.63 Jun 18 2020

- cluster: Remove support for expiration_timestamp
- Added top-level Notify endpoint to AMS

## 0.0.62 Jun 9 2020

- Add subscription notify endpoint

## 0.0.61 Jun 9 2020

- accounts_mgmt: Add 'fields' parameter to all list-requests
- accounts_mgmt: Support for Labels resources

## 0.0.60 Jun 3 2020

- Add parameters 'offset' and 'tail' to log resource

## 0.0.59 May 20 2020

- FeatureToggle: Add model and resource

## 0.0.58 May 15 2020

- AddOns: Add link attribute
- Update to metamodel v0.0.28

## 0.0.57 May 13 2020

- AddOnInstallations: Remove DELETE operation
- Added Label to Account

## 0.0.56 May 03 2020

- Added Label to Organization

## 0.0.55 Apr 23 2020

- Add enabled field to region
- Adding metrics.nodes to api model
- Adding cluster ingresses endpoint
- ClusterNodes: Add ComputeMachineType
- Network: Added HostPrefix

## 0.0.54 Apr 7 2020

- Update to metamodel 0.0.27

## 0.0.53 Apr 3 2020

- Add pull secret deletion
- Products: Add product attribute to cluster object
- Products: Support for top-level cluster types
- Add ClusterOperatorsConditions type
- Add ClusterAlertsFiring type and field in ClusterMetrics

## 0.0.52 Mar 26 2020

- Removal of `in` parameters for `Get` functions

## 0.0.51 Mar 25 2020

- Update AMS Models

## 0.0.50 Mar 24 2020

- Add sockets to cluster_metrics_type

## 0.0.49 Mar 24 2020

- Add `Ingress` resource.

## 0.0.48 Mar 22 2020

- Add `API` listening method.

## 0.0.47 Mar 19 2020

- Add `ClusterAdminEnabled` flag.
- Add `PullSecrets` endpoint.
- Fix `LDAPIdentityProvider` attribute name.

## 0.0.46 Mar 18 2020

- Add new fields to `AddOn` and `AddOnInstallation`.

## 0.0.45 Mar 11 2020

- Add `Organizations` attribute to GitHub `IdP`

## 0.0.44 Mar 9 2020

- Remove duplicated attribute.

## 0.0.43 Mar 9 2020

- Improve documentation of the `LogEntry` type.

## 0.0.42 Mar 5 2020

- Add `client_secret` attribute to _GitHub_ identity provider.

## 0.0.41 Feb 13 2020

- Add `target_namespace` and `install_mode` attributes to `AddOn` type.
- Add `state` attribute to `AWSInfrastructureAccessRole` type.

## 0.0.40 Feb 5 2020

- Add method to update flavour.

## 0.0.39 Feb 3 2020

- Add types and resources for cluster operator metrics.
- Add `deleting` and `removed` states to AWS infrastructure access role grant
  status.

## 0.0.38 Jan 23 2020

- Add `search` and `order` parameters to the method that lists registry
  credentials.
- Add `labels` parameter to the method that lists subscriptions.
- Add types and resources for management of AWS infrastructure access roles.

## 0.0.37 Jan 8 2020

- Add new `service_logs` service.
- Add types and resources for machine types.

## 0.0.36 Jan 3 2020

- Add types and resources for AWS infrastructure access roles.
- Add GCP flavour and change AWS flavour to contain also the instance type.

## 0.0.35 Jan 01 2020

- Fixes for `CurrentAccess` resource.

## 0.0.34 Jan 01 2020

- Add `CurrentAccess` resource.

## 0.0.33 Dec 31 2019

- Add `UpdatedAt` and `CreatedAt` fields to `Subscription` type.

## 0.0.32 Dec 24 2019

- Replace `AddOns` with `AddOnInstallations`.

## 0.0.31 Dec 19 2019

- Add `ban_code` attribute to `Account` type.

## 0.0.30 Dec 17 2019

- Add support for `ClusterUUID` field.

## 0.0.29 Dec 12 2019

- Allow subscription identifier on role binding.

## 0.0.28 Dec 10 2019

- Add `AddOnInstallation` type.

## 0.0.27 Dec 4 2019

- Add `resource_name` and `resource_cost` attributes to the add-on type.

## 0.0.26 Dec 2 2019

- Remove obsolete `aws` and `version` fields from the `Flavour` type.
- Add instance type fields to the `Flavour` type.
- Add `AWSVolume` and `AWSFlavour` types.
- Add attributes required for _BYOC_.
- Fix direction of `Body` parameters of updates.

## 0.0.25 Nov 28 2019

- Allow patching role binding.

## 0.0.24 Nov 23 2019

- Fix directions of paging parameters.
- Fix direction of `Body` parameter of `Update`.
- Add default values to paging parameters.
- Update to metamodel 0.0.17.

## 0.0.23 Nov 20 2019

- Add infra nodes to `FlavourNodes`.
- Refactor flavour nodes.

## 0.0.22 Nov 19 2019

- Add `socket_total_by_node_roles_os` metric query.

## 0.0.21 Nov 17 2019

- Added add-on resources and types.
- Added subscription reserved resources collection.

## 0.0.20 Nov 14 2019

- Query resource quota from root and delete by identifier.

## 0.0.19 Nov 8 2019

- Added identifiers to role binding type.

## 0.0.18 Nov 7 2019

- Added support to search role bindings and resource quota.

## 0.0.17 Oct 28 2019

- Added `Disconnected`, `DisplayName` and `ExternalClusterID` attributes to the
  cluster authorization request type.

## 0.0.16 Oct 27 2019

- Added `ResourceReview` resource to the authorizations service.

## 0.0.15 Oct 24 2019

- Added `search` parameter to the accounts `List` method.

## 0.0.14 Oct 24 2019

- Added `SKU` type.
- Improved organizations.
- Improved roles.

## 0.0.13 Oct 15 2019

- Added `AccessTokenAuth` type.
- Added `auths` attribute to `AccessToken` type.
- Update to metamodel 0.0.9.

## 0.0.12 Oct 10 2019

- Add `access_review` resource.

## 0.0.11 Oct 10 2019

- Add `export_control_review` resource.

## 0.0.10 Oct 7 2019

- Add `cpu_total_by_node_roles_os` metric query.

## 0.0.9 Oct 7 2019

- Add `type` attribute to the `ResourceQuota` type.
- Add `config_managed` attribute to the `RoleBinding` type.

## 0.0.8 Sep 17 2019

- Update methods don't return body.

## 0.0.7 Sep 16 2019

- Add `search` parameter to the `List` method of the subscriptions resource.

## 0.0.6 Sep 16 2019

- Remove the `creator` attribute of the `Cluster` type.

## 0.0.5 Sep 12 2019

- Add `order` parameter to the methods to list accounts and subscriptions.

## 0.0.4 Sep 12 2019

- Update to metamodel 0.0.6:
** Explicitly enable Go modules so that the build works correctly when the
   project is located inside the Go path.

## 0.0.3 Sep 11 2019

- Add `order` parameter to the collections that suport it.
- Add cloud providers collection.

## 0.0.2 Sep 10 2019

- Add `DisplayName` attribute to `Subscription` type.

## 0.0.1 Aug 20 2019

- Changed the type of the `ExpiresAt` attribute of the
  `ClusterRegistrationResponse` type from `long` to `string`.
