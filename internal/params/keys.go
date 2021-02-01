package params

import "strings"

var (
	BaseURIKey                          = strings.ToLower(BaseURIEnv)
	AstUsernameKey                      = strings.ToLower(AstUsernameEnv)
	AstPasswordKey                      = strings.ToLower(AstPasswordEnv)
	AstTokenKey                         = strings.ToLower(AstTokenEnv)
	ScansPathKey                        = strings.ToLower(ScansPathEnv)
	ProjectsPathKey                     = strings.ToLower(ProjectsPathEnv)
	ResultsPathKey                      = strings.ToLower(ResultsPathEnv)
	BflPathKey                          = strings.ToLower(BflPathEnv)
	UploadsPathKey                      = strings.ToLower(UploadsPathEnv)
	SastRmPathKey                       = strings.ToLower(SastRmPathEnv)
	AccessKeyIDConfigKey                = strings.ToLower(AccessKeyIDEnv)
	AccessKeySecretConfigKey            = strings.ToLower(AccessKeySecretEnv)
	AstAuthenticationPathConfigKey      = strings.ToLower(AstAuthenticationPathEnv)
	CredentialsFilePathKey              = strings.ToLower(CredentialsFilePathEnv)
	TokenExpirySecondsKey               = strings.ToLower(TokenExpirySecondsEnv)
	AstRoleKey                          = strings.ToLower(AstRoleEnv)
	AstWebAppHealthCheckPathKey         = strings.ToLower(AstWebAppHealthCheckPathEnv)
	AstKeycloakWebAppHealthCheckPathKey = strings.ToLower(AstKeycloakWebAppHealthCheckPathEnv)
	HealthcheckPathKey                  = strings.ToLower(HealthcheckPathEnv)
	HealthcheckDBPathKey                = strings.ToLower(HealthcheckDBPathEnv)
	HealthcheckMessageQueuePathKey      = strings.ToLower(HealthcheckMessageQueuePathEnv)
	HealthcheckObjectStorePathKey       = strings.ToLower(HealthcheckObjectStorePathEnv)
	HealthcheckInMemoryDBPathKey        = strings.ToLower(HealthcheckInMemoryDBPathEnv)
	HealthcheckLoggingPathKey           = strings.ToLower(HealthcheckDBPathEnv)
	HealthcheckScanFlowPathKey          = strings.ToLower(HealthcheckScanFlowPathEnv)
	HealthcheckSastEnginesPathKey       = strings.ToLower(HealthcheckSastEnginesPathEnv)
	QueriesPathKey                      = strings.ToLower(QueriesPathEnv)
	QueriesClonePathKey                 = strings.ToLower(QueriesCLonePathEnv)
	CreateOath2ClientPathKey            = strings.ToLower(CreateOath2ClientPathEnv)
	SastScanIncPathKey                  = strings.ToLower(SastScanIncPathEnv)
	SastScanIncEngineLogPathKey         = strings.ToLower(SastScanIncEngineLogPathEnv)
	SastScanIncMetricsPathKey           = strings.ToLower(SastScanIncMetricsPathEnv)
	LogsPathKey                         = strings.ToLower(LogsPathEnv)
)
