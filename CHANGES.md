# Changes

This document describes the relevant changes between releases of the API model.

## 0.0.216 Aug 11 2022
- Change provision shard to include kube client configurations.
- Add GCP volume size to flavour API.
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

-  Add Name field to LDAP identity provider

## 0.0.150 Oct 27 2021

-  Fix addon installation version (addon_version vs version)
-  Remove no_proxy attribute from SDK
-  Add body to the external tracking event

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

-  Add cluster resource type
-  Remove GET cluster deployment endpoint

## 0.0.136 Jul 19 2021

-  Added Spot Market Options to MachinePool


## 0.0.135 Jul 6 2021

-  Added username to service log LogEntry

## 0.0.134 Jun 22 2021

-  Added InternalOnly flag to SubscriptionNotify

## 0.0.133 Jun 16 2021

-  Added capabilities support to Organization

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

- Add API for custerdeployment labels
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
