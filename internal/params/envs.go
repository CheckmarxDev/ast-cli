package params

const (
	TenantEnv                           = "CX_TENANT"
	BranchEnv                           = "CX_BRANCH"
	BaseURIEnv                          = "CX_BASE_URI"
	ClientTimeoutEnv                    = "CX_TIMEOUT"
	ProxyEnv                            = "HTTP_PROXY"
	CxProxyEnv                          = "CX_HTTP_PROXY"
	ProxyTypeEnv                        = "CX_PROXY_AUTH_TYPE"
	ProxyDomainEnv                      = "CX_PROXY_NTLM_DOMAIN"
	BaseAuthURIEnv                      = "CX_BASE_AUTH_URI"
	AstAPIKeyEnv                        = "CX_APIKEY"
	AccessKeyIDEnv                      = "CX_CLIENT_ID"
	AccessKeySecretEnv                  = "CX_CLIENT_SECRET"
	ScansPathEnv                        = "CX_SCANS_PATH"
	CodeBashingPathEnv                  = "CX_CODEBASHING_PATH"
	GroupsPathEnv                       = "CX_GROUPS_PATH"
	AgentNameEnv                        = "CX_AGENT_NAME"
	ProjectsPathEnv                     = "CX_PROJECTS_PATH"
	ApplicationsPathEnv                 = "CX_APPLICATIONS_PATH"
	ResultsPathEnv                      = "CX_RESULTS_PATH"
	ScanSummaryPathEnv                  = "CX_SCAN_SUMMARY_PATH"
	ScaPackagePathEnv                   = "CX_SCA_PACKAGE_PATH"
	RisksOverviewPathEnv                = "CX_RISKS_OVERVIEW_PATH"
	ScsScanOverviewPathEnv              = "CX_SCS_SCAN_OVERVIEW_PATH"
	SastResultsPathEnv                  = "CX_SAST_RESULTS_PATH"
	SastResultsPredicatesPathEnv        = "CX_SAST_RESULTS_PREDICATES_PATH"
	KicsResultsPathEnv                  = "CX_KICS_RESULTS_PATH"
	KicsResultsPredicatesPathEnv        = "CX_KICS_RESULTS_PREDICATES_PATH"
	BflPathEnv                          = "CX_BFL_PATH"
	PRDecorationGithubPathEnv           = "CX_PR_DECORATION_GITHUB_PATH"
	PRDecorationGitlabPathEnv           = "CX_PR_DECORATION_GITLAB_PATH"
	SastRmPathEnv                       = "CX_SAST_RM_PATH"
	UploadsPathEnv                      = "CX_UPLOADS_PATH"
	TokenExpirySecondsEnv               = "CX_TOKEN_EXPIRY_SECONDS"
	AstRoleEnv                          = "CX_AST_ROLE"
	AstWebAppHealthCheckPathEnv         = "CX_AST_WEB_APP_HEALTH_CHECK_PATH"
	AstKeycloakWebAppHealthCheckPathEnv = "CX_AST_KEYCLOAK_WEB_APP_HEALTH_CHECK_PATH"
	HealthcheckPathEnv                  = "CX_HEALTHCHECK_PATH"
	HealthcheckDBPathEnv                = "CX_HEALTHCHECK_DB_PATH"
	HealthcheckMessageQueuePathEnv      = "CX_HEALTHCHECK_MESSAGE_QUEUE_PATH"
	HealthcheckObjectStorePathEnv       = "CX_HEALTHCHECK_OBJECT_STORE_PATH"
	HealthcheckInMemoryDBPathEnv        = "CX_HEALTHCHECK_IN_MEMORY_DB_PATH"
	HealthcheckLoggingPathEnv           = "CX_HEALTHCHECK_LOGGING_PATH"
	HealthcheckScanFlowPathEnv          = "CX_HEALTHCHECK_SCAN_FLOW_PATH"
	HealthcheckSastEnginesPathEnv       = "CX_HEALTHCHECK_SAST_ENGINES_PATH"
	QueriesPathEnv                      = "CX_QUERIES_PATH"
	QueriesCLonePathEnv                 = "CX_QUERIES_CLONE_PATH"
	CreateOath2ClientPathEnv            = "CX_CREATE_OATH2_CLIENT_PATH"
	SastScanIncPathEnv                  = "CX_SAST_SCAN_INC_PATH"
	SastScanIncMetricsPathEnv           = "CX_SAST_SCAN_INC_METRICS_PATH"
	LogsPathEnv                         = "CX_LOGS_PATH"
	LogsEngineLogPathEnv                = "CX_LOGS_ENGINE_LOG_PATH"
	DescriptionsPathEnv                 = "CX_DESCRIPTIONS_PATH"
	TenantConfigurationPathEnv          = "CX_TENANT_CONFIGURATION_PATH"
	ResultsPdfReportPathEnv             = "CX_RESULTS_PDF_REPORT_PATH"
	ResultsSbomReportPathEnv            = "CX_RESULTS_SBOM_PATH"
	ResultsSbomReportProxyPathEnv       = "CX_RESULTS_SBOM_PROXY_PATH"
	FeatureFlagsEnv                     = "CX_FEATURE_FLAGS_PATH"
	UploadURLEnv                        = "CX_UPLOAD_URL"
	PolicyEvaluationPathEnv             = "CX_POLICY_EVALUATION_PATH"
	AccessManagementPathEnv             = "CX_ACCESS_MANAGEMENT_PATH"
	ByorPathEnv                         = "CX_BYOR_PATH"
	IgnoreProxyEnv                      = "CX_IGNORE_PROXY"
	VorpalPortEnv                       = "CX_VORPAL_PORT"
	ExportPathEnv                       = "CX_EXPORT_PATH"
)
