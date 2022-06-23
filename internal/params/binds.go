package params

var EnvVarsBinds = []struct {
	Key     string
	Env     string
	Default string
}{
	{BaseURIKey, BaseURIEnv, ""},
	{ProxyKey, ProxyEnv, ""},
	{ProxyTypeKey, ProxyTypeEnv, "basic"},
	{ProxyDomainKey, ProxyDomainEnv, ""},
	{BaseAuthURIKey, BaseAuthURIEnv, ""},
	{AstAPIKey, AstAPIKeyEnv, ""},
	{AgentNameKey, AgentNameEnv, "ASTCLI"},
	{CodeBashingPathKey, ScansPathEnv, "api/codebashing/lessons"},
	{ScansPathKey, ScansPathEnv, "api/scans"},
	{ProjectsPathKey, ProjectsPathEnv, "api/projects"},
	{GroupsPathKey, GroupsPathEnv, "auth/realms/organization/pip/groups"},
	{ResultsPathKey, ResultsPathEnv, "api/results"},
	{ScaPackagePathKey, ScaPackagePathEnv, "api/sca/risk-management/risk-reports/"},
	{SastResultsPathKey, SastResultsPathEnv, "api/sast-results"},
	{SastResultsPredicatesPathKey, SastResultsPredicatesPathEnv, "api/sast-results-predicates"},
	{KicsResultsPathKey, KicsResultsPathEnv, "api/kics-results"},
	{KicsResultsPredicatesPathKey, KicsResultsPredicatesPathEnv, "api/kics-results-predicates"},
	{BflPathKey, BflPathEnv, "api/bfl"},
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
	{AstAuthenticationPathConfigKey, AstAuthenticationPathEnv, "auth/realms/organization/protocol/openid-connect/token"},
	{TenantKey, TenantEnv, "organization"},
	{BranchKey, BranchEnv, ""},
	{AstRoleKey, AstRoleEnv, ScaAgent},
	{TokenExpirySecondsKey, TokenExpirySecondsEnv, "300"},
	{ClientTimeoutKey, ClientTimeoutEnv, "5"},
}
