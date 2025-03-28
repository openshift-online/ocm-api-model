//go:build tools
// +build tools

// go mod won't pull in code that isn't depended upon, but we have some code we don't depend on from code that must be included
// for our build to work.
package dependencymagnet

import (
	// this gives us clear dependency control of our generator, easy replaces for development, ease of vendored inspection, and fully local builds.
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/access_transparency/v1"
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/accounts_mgmt/v1"
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/addons_mgmt/v1"
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/aro_hcp/v1alpha1"
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/authorizations/v1"
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/clusters_mgmt/v1"
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/job_queue/v1"
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/osd_fleet_mgmt/v1"
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/service_logs/v1"
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/service_mgmt/v1"
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/status_board/v1"
	_ "github.com/openshift-online/ocm-api-model/clientapi/model/web_rca/v1"
)
