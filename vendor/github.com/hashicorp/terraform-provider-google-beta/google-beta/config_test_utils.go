// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"net/http/httptest"
	"strings"
)

const testFakeCredentialsPath = "./test-fixtures/fake_account.json"

// NewTestConfig create a config using the http test server.
func NewTestConfig(server *httptest.Server) *transport_tpg.Config {
	cfg := &transport_tpg.Config{}
	cfg.Client = server.Client()
	configureTestBasePaths(cfg, server.URL)
	return cfg
}

func configureTestBasePaths(c *transport_tpg.Config, url string) {
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	// Generated Products
	c.AccessApprovalBasePath = url
	c.AccessContextManagerBasePath = url
	c.ActiveDirectoryBasePath = url
	c.AlloydbBasePath = url
	c.ApiGatewayBasePath = url
	c.ApigeeBasePath = url
	c.AppEngineBasePath = url
	c.ArtifactRegistryBasePath = url
	c.BeyondcorpBasePath = url
	c.BigQueryBasePath = url
	c.BigqueryAnalyticsHubBasePath = url
	c.BigqueryConnectionBasePath = url
	c.BigqueryDatapolicyBasePath = url
	c.BigqueryDataTransferBasePath = url
	c.BigqueryReservationBasePath = url
	c.BigtableBasePath = url
	c.BillingBasePath = url
	c.BinaryAuthorizationBasePath = url
	c.CertificateManagerBasePath = url
	c.CloudAssetBasePath = url
	c.CloudBuildBasePath = url
	c.Cloudbuildv2BasePath = url
	c.CloudFunctionsBasePath = url
	c.Cloudfunctions2BasePath = url
	c.CloudIdentityBasePath = url
	c.CloudIdsBasePath = url
	c.CloudIotBasePath = url
	c.CloudRunBasePath = url
	c.CloudRunV2BasePath = url
	c.CloudSchedulerBasePath = url
	c.CloudTasksBasePath = url
	c.ComputeBasePath = url
	c.ContainerAnalysisBasePath = url
	c.ContainerAttachedBasePath = url
	c.DatabaseMigrationServiceBasePath = url
	c.DataCatalogBasePath = url
	c.DataformBasePath = url
	c.DataFusionBasePath = url
	c.DataLossPreventionBasePath = url
	c.DataplexBasePath = url
	c.DataprocBasePath = url
	c.DataprocMetastoreBasePath = url
	c.DatastoreBasePath = url
	c.DatastreamBasePath = url
	c.DeploymentManagerBasePath = url
	c.DialogflowBasePath = url
	c.DialogflowCXBasePath = url
	c.DNSBasePath = url
	c.DocumentAIBasePath = url
	c.EssentialContactsBasePath = url
	c.FilestoreBasePath = url
	c.FirebaseBasePath = url
	c.FirebaseDatabaseBasePath = url
	c.FirebaseExtensionsBasePath = url
	c.FirebaseHostingBasePath = url
	c.FirebaseStorageBasePath = url
	c.FirestoreBasePath = url
	c.GameServicesBasePath = url
	c.GKEBackupBasePath = url
	c.GKEHubBasePath = url
	c.GKEHub2BasePath = url
	c.GkeonpremBasePath = url
	c.HealthcareBasePath = url
	c.IAM2BasePath = url
	c.IAMBetaBasePath = url
	c.IAMWorkforcePoolBasePath = url
	c.IapBasePath = url
	c.IdentityPlatformBasePath = url
	c.KMSBasePath = url
	c.LoggingBasePath = url
	c.LookerBasePath = url
	c.MemcacheBasePath = url
	c.MLEngineBasePath = url
	c.MonitoringBasePath = url
	c.NetworkManagementBasePath = url
	c.NetworkSecurityBasePath = url
	c.NetworkServicesBasePath = url
	c.NotebooksBasePath = url
	c.OrgPolicyBasePath = url
	c.OSConfigBasePath = url
	c.OSLoginBasePath = url
	c.PrivatecaBasePath = url
	c.PublicCABasePath = url
	c.PubsubBasePath = url
	c.PubsubLiteBasePath = url
	c.RedisBasePath = url
	c.ResourceManagerBasePath = url
	c.RuntimeConfigBasePath = url
	c.SecretManagerBasePath = url
	c.SecurityCenterBasePath = url
	c.SecurityScannerBasePath = url
	c.ServiceDirectoryBasePath = url
	c.ServiceManagementBasePath = url
	c.ServiceUsageBasePath = url
	c.SourceRepoBasePath = url
	c.SpannerBasePath = url
	c.SQLBasePath = url
	c.StorageBasePath = url
	c.StorageTransferBasePath = url
	c.TagsBasePath = url
	c.TPUBasePath = url
	c.VertexAIBasePath = url
	c.VmwareengineBasePath = url
	c.VPCAccessBasePath = url
	c.WorkflowsBasePath = url
	c.WorkstationsBasePath = url

	// Handwritten Products / Versioned / Atypical Entries
	c.CloudBillingBasePath = url
	c.ComposerBasePath = url
	c.ContainerBasePath = url
	c.DataprocBasePath = url
	c.DataflowBasePath = url
	c.IamCredentialsBasePath = url
	c.ResourceManagerV3BasePath = url
	c.IAMBasePath = url
	c.ServiceNetworkingBasePath = url
	c.BigQueryBasePath = url
	c.StorageTransferBasePath = url
	c.BigtableAdminBasePath = url
}