package params

var EnvVarsBinds = []struct {
	Key     string
	Env     string
	Default string
}{
	{BaseURIKey, BaseURIEnv, ""},
	{ProxyTypeKey, ProxyTypeEnv, "basic"},
	{ProxyDomainKey, ProxyDomainEnv, ""},
	{BaseAuthURIKey, BaseAuthURIEnv, ""},
	{AstAPIKey, AstAPIKeyEnv, ""},
	{IgnoreProxyKey, IgnoreProxyEnv, ""},
	{AgentNameKey, AgentNameEnv, "ASTCLI"},
	{CodeBashingPathKey, ScansPathEnv, "api/codebashing/lessons"},
	{CustomStatesAPIPathKey, CustomStatesAPIPathEnv, "api/custom-states"},
	{ScansPathKey, ScansPathEnv, "api/scans"},
	{ProjectsPathKey, ProjectsPathEnv, "api/projects"},
	{ApplicationsPathKey, ApplicationsPathEnv, "api/applications"},
	{GroupsPathKey, GroupsPathEnv, "auth/realms/organization/pip/groups"},
	{ResultsPathKey, ResultsPathEnv, "api/results"},
	{ScanSummaryPathKey, ScanSummaryPathEnv, "api/scan-summary"},
	{RisksOverviewPathKey, RisksOverviewPathEnv, "api/apisec/static/api/scan/%s/risks-overview"},
	{ScsScanOverviewPathKey, ScsScanOverviewPathEnv, "api/micro-engines/scans/%s/scan-overview"},
	{SastResultsPathKey, SastResultsPathEnv, "api/sast-results"},
	{SastResultsPredicatesPathKey, SastResultsPredicatesPathEnv, "api/sast-results-predicates"},
	{KicsResultsPathKey, KicsResultsPathEnv, "api/kics-results"},
	{KicsResultsPredicatesPathKey, KicsResultsPredicatesPathEnv, "api/kics-results-predicates"},
	{ScsResultsPredicatesPathKey, ScsResultsPredicatesPathEnv, "api/micro-engines/predicates"},
	{BflPathKey, BflPathEnv, "api/bfl"},
	{PRDecorationGithubPathKey, PRDecorationGithubPathEnv, "api/flow-publisher/pr/github"},
	{PRDecorationGitlabPathKey, PRDecorationGitlabPathEnv, "api/flow-publisher/pr/gitlab"},
	{PRDecorationBitbucketCloudPathKey, PRDecorationBitbucketCloudPathEnv, "api/flow-publisher/pr/bitbucket"},
	{PRDecorationBitbucketServerPathKey, PRDecorationBitbucketServerPathEnv, "api/flow-publisher/pr/bitbucket-server"},
	{PRDecorationAzurePathKey, PRDecorationAzurePathEnv, "api/flow-publisher/pr/azure"},
	{DescriptionsPathKey, DescriptionsPathEnv, "api/queries/descriptions"},
	{TenantConfigurationPathKey, TenantConfigurationPathEnv, "api/configuration/tenant"},
	{UploadsPathKey, UploadsPathEnv, "api/uploads"},
	{SastRmPathKey, SastRmPathEnv, "api/sast-rm"},
	{AstWebAppHealthCheckPathKey, AstWebAppHealthCheckPathEnv, "#/projects"},
	{AstKeycloakWebAppHealthCheckPathKey, AstKeycloakWebAppHealthCheckPathEnv, "auth"},
	{HealthcheckPathKey, HealthcheckPathEnv, "api/healthcheck"},
	{HealthcheckDBPathKey, HealthcheckDBPathEnv, "database"},
	{HealthcheckMessageQueuePathKey, HealthcheckMessageQueuePathEnv, "message-queue"},
	{HealthcheckObjectStorePathKey, HealthcheckObjectStorePathEnv, "object-store"},
	{HealthcheckInMemoryDBPathKey, HealthcheckInMemoryDBPathEnv, "in-memory-db"},
	{HealthcheckLoggingPathKey, HealthcheckLoggingPathEnv, "logging"},
	{HealthcheckScanFlowPathKey, HealthcheckScanFlowPathEnv, "scan-flow"},
	{HealthcheckSastEnginesPathKey, HealthcheckSastEnginesPathEnv, "sast-engines"},
	{QueriesPathKey, QueriesPathEnv, "api/queries"},
	{QueriesClonePathKey, QueriesCLonePathEnv, "clone"},
	{CreateOath2ClientPathKey, CreateOath2ClientPathEnv, "auth/realms/organization/pip/clients"},
	{SastMetadataPathKey, SastScanIncPathEnv, "api/sast-metadata"},
	{SastMetadataMetricsPathKey, SastScanIncMetricsPathEnv, "%s/metrics"},
	{LogsPathKey, LogsPathEnv, "api/logs"},
	{LogsEngineLogPathKey, LogsEngineLogPathEnv, "/%s/%s"},
	{AccessKeyIDConfigKey, AccessKeyIDEnv, ""},
	{AccessKeySecretConfigKey, AccessKeySecretEnv, ""},
	{TenantKey, TenantEnv, ""},
	{BranchKey, BranchEnv, ""},
	{AstRoleKey, AstRoleEnv, ScaAgent},
	{TokenExpirySecondsKey, TokenExpirySecondsEnv, "300"},
	{ClientTimeoutKey, ClientTimeoutEnv, "30"},
	{ResultsPdfReportPathKey, ResultsPdfReportPathEnv, "api/reports"},
	{ExportPathKey, ExportPathEnv, "api/sca/export"},
	{FeatureFlagsKey, FeatureFlagsEnv, "api/flags"},
	{PolicyEvaluationPathKey, PolicyEvaluationPathEnv, "api/policy_management_service_uri/evaluation"},
	{AccessManagementPathKey, AccessManagementPathEnv, "api/access-management"},
	{ByorPathKey, ByorPathEnv, "api/byor"},
	{AiProxyAzureAiRouteKey, AiProxyAzureAiRouteEnv, "api/ai-proxy/redirect/externalAzure"},
	{AiProxyCheckmarxAiRouteKey, AiProxyCheckmarxAiRouteEnv, "api/ai-proxy/redirect/azure"},
	{ASCAPortKey, ASCAPortEnv, ""},
	{ScsRepoTokenKey, ScsRepoTokenEnv, ""},
}
