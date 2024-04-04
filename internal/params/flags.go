package params

// Flags
const (
	AgentFlag                     = "agent"
	AgentFlagUsage                = "Scan origin name"
	ApplicationName               = "application-name"
	DefaultAgent                  = "ASTCLI"
	DebugFlag                     = "debug"
	DebugUsage                    = "Debug mode with detailed logs"
	RetryFlag                     = "retry"
	RetryDefault                  = 3
	RetryUsage                    = "Retry requests to Checkmarx One on connection failure"
	RetryDelayFlag                = "retry-delay"
	RetryDelayDefault             = 20
	RetryDelayPollingDefault      = 60
	RetryDelayUsage               = "Time between retries in seconds, use with --" + RetryFlag
	SourcesFlag                   = "file-source"
	SourcesFlagSh                 = "s"
	TenantFlag                    = "tenant"
	TenantFlagUsage               = "Checkmarx tenant"
	AsyncFlag                     = "async"
	WaitDelayFlag                 = "wait-delay"
	ScanTimeoutFlag               = "scan-timeout"
	PolicyTimeoutFlag             = "policy-timeout"
	IgnorePolicyFlag              = "ignore-policy"
	SourceDirFilterFlag           = "file-filter"
	SourceDirFilterFlagSh         = "f"
	IncludeFilterFlag             = "file-include"
	IncludeFilterFlagSh           = "i"
	ProjectIDFlag                 = "project-id"
	BranchFlag                    = "branch"
	BranchFlagSh                  = "b"
	ScanIDFlag                    = "scan-id"
	BranchFlagUsage               = "Branch to scan"
	MainBranchFlag                = "branch"
	ScaResolverFlag               = "sca-resolver"
	ScaResolverParamsFlag         = "sca-resolver-params"
	AccessKeyIDFlag               = "client-id"
	AccessKeySecretFlag           = "client-secret"
	AccessKeyIDFlagUsage          = "The OAuth2 client ID"
	AccessKeySecretFlagUsage      = "The OAuth2 client secret"
	InsecureFlag                  = "insecure"
	InsecureFlagUsage             = "Ignore TLS certificate validations"
	ScanInfoFormatFlag            = "scan-info-format"
	FormatFlag                    = "format"
	FormatFlagUsageFormat         = "Format for the output. One of %s"
	FilterFlag                    = "filter"
	BaseURIFlag                   = "base-uri"
	ProxyFlag                     = "proxy"
	ProxyFlagUsage                = "Proxy server to send communication through"
	IgnoreProxyFlag               = "ignore-proxy"
	IgnoreProxyFlagUsage          = "Ignore proxy configuration"
	ProxyTypeFlag                 = "proxy-auth-type"
	ProxyTypeFlagUsage            = "Proxy authentication type, (basic or ntlm)"
	TimeoutFlag                   = "timeout"
	TimeoutFlagUsage              = "Timeout for network activity, (default 5 seconds)"
	NtlmProxyDomainFlag           = "proxy-ntlm-domain"
	NtlmProxyDomainFlagUsage      = "Window domain when using NTLM proxy"
	BaseURIFlagUsage              = "The base system URI"
	BaseAuthURIFlag               = "base-auth-uri"
	BaseAuthURIFlagUsage          = "The base system IAM URI"
	AstAPIKeyFlag                 = "apikey"
	AstAPIKeyUsage                = "The API Key to login to Checkmarx One"
	ClientRolesFlag               = "roles"
	ClientRolesSh                 = "r"
	ClientDescriptionFlag         = "description"
	ClientDescriptionSh           = "d"
	UsernameFlag                  = "username"
	UsernameSh                    = "u"
	PasswordFlag                  = "password"
	PasswordSh                    = "p"
	ProfileFlag                   = "profile"
	ProfileFlagUsage              = "The default configuration profile"
	Help                          = "help"
	TargetFlag                    = "output-name"
	TargetPathFlag                = "output-path"
	TargetFormatFlag              = "report-format"
	ReportFormatPdfToEmailFlag    = "report-pdf-email"
	ReportFormatPdfOptionsFlag    = "report-pdf-options"
	ReportSbomFormatFlag          = "report-sbom-format"
	ReportSbomFormatLocalFlowFlag = "report-sbom-local-flow"
	ProjectName                   = "project-name"
	ScanTypes                     = "scan-types"
	ScanTypeFlag                  = "scan-type"
	ScanResubmit                  = "resubmit"
	KicsRealtimeFile              = "file"
	KicsRealtimeEngine            = "engine"
	KicsRealtimeAdditionalParams  = "additional-params"
	ScaRealtimeProjectDir         = "project-dir"
	ScaRealtimeProjectDirSh       = "p"
	RemediationFiles              = "package-files"
	KicsRemediationFile           = "results-file"
	KicsProjectFile               = "kics-files"
	KicsSimilarityFilter          = "similarity-ids"
	RemediationPackage            = "package"
	RemediationPackageVersion     = "package-version"
	TagList                       = "tags"
	GroupList                     = "groups"
	ProjectGroupList              = "project-groups"
	ProjectTagList                = "project-tags"
	IncrementalSast               = "sast-incremental"
	PresetName                    = "sast-preset-name"
	Threshold                     = "threshold"
	ThresholdFlagUsage            = "Local build threshold. Format <engine>-<severity>=<limit>. " +
		"Example: scan --threshold \"sast-critical=1;sast-high=10;sca-high=5;iac-security-low=10\""
	KeyValuePairSize         = 2
	WaitDelayDefault         = 5
	SimilarityIDFlag         = "similarity-id"
	SeverityFlag             = "severity"
	StateFlag                = "state"
	CommentFlag              = "comment"
	LanguageFlag             = "language"
	VulnerabilityTypeFlag    = "vulnerability-type"
	CweIDFlag                = "cwe-id"
	SCMTokenFlag             = "token"
	AzureTokenUsage          = "Azure DevOps personal access token. Requires “Connected server” and “Code“ scope."
	GithubTokenUsage         = "GitHub OAuth token. Requires “Repo” scope and organization SSO authorization, if enforced by the organization"
	GitLabTokenUsage         = "GitLab OAuth token"
	BotCount                 = "Note: dependabot is not counted but other bots might be considered as contributors."
	URLFlag                  = "url"
	GitLabURLFlag            = "url-gitlab"
	URLFlagUsage             = "API base URL"
	QueryIDFlag              = "query-id"
	SSHKeyFlag               = "ssh-key"
	RepoURLFlag              = "repo-url"
	AstToken                 = "ast-token"
	SSHValue                 = "ssh-value"
	KicsContainerNameKey     = "kics-container-name"
	KicsPlatformsFlag        = "kics-platforms"
	KicsPlatformsFlagUsage   = "KICS Platform Flag. Use ',' as the delimiter for arrays."
	IacsPlatformsFlag        = "iac-security-platforms"
	IacsPlatformsFlagUsage   = "IaC Security Platform Flag"
	ApikeyOverrideFlag       = "apikey-override"
	ExploitablePathFlag      = "sca-exploitable-path"
	LastSastScanTime         = "sca-last-sast-scan-time"
	ProjecPrivatePackageFlag = "project-private-package"
	SastRedundancyFlag       = "sast-redundancy"

	ScaPrivatePackageVersionFlag = "sca-private-package-version"

	// INDIVIDUAL FILTER FLAGS
	SastFilterFlag  = "sast-filter"
	SastFilterUsage = "SAST filter"
	KicsFilterFlag  = "kics-filter"
	IacsFilterFlag  = "iac-security-filter"
	IacsFilterUsage = "IaC Security filter"
	KicsFilterUsage = "KICS filter"
	ScaFilterFlag   = "sca-filter"
	ScaFilterUsage  = "SCA filter"

	// PR decoration flags
	NamespaceFlag            = "namespace"
	NamespaceFlagUsage       = "%s namespace is required to post the comments"
	RepoNameFlag             = "repo-name"
	RepoNameFlagUsage        = "%s repository details"
	PRNumberFlag             = "pr-number"
	PRNumberFlagUsage        = "Pull Request number for posting notifications and comments"
	PRIidFlag                = "mr-iid"
	PRIidFlagUsage           = "Gitlab IID (internal ID) of the merge request"
	PRGitlabProjectFlag      = "gitlab-project-id"
	PRGitlabProjectFlagUsage = "Gitlab project ID"

	// Chat (General)
	ChatAPIKey         = "chat-apikey"
	ChatConversationID = "conversation-id"
	ChatUserInput      = "user-input"
	ChatModel          = "model"

	// Chat (Kics)
	ChatKicsResultFile          = "result-file"
	ChatKicsResultLine          = "result-line"
	ChatKicsResultSeverity      = "result-severity"
	ChatKicsResultVulnerability = "result-vulnerability"

	// Mask
	MaskContent = "mask-content"

	// Chat (SAST)
	ChatSastScanResultsFile = "scan-results-file"
	ChatSastSourceDir       = "source-dir"
	ChatSastResultID        = "sast-result-id"
)

// Parameter values
const (
	IDQueryParam               = "id"
	IDsQueryParam              = "ids"
	IDRegexQueryParam          = "id-regex"
	LimitQueryParam            = "limit"
	OffsetQueryParam           = "offset"
	ScanIDQueryParam           = "scan-id"
	ScanIDsQueryParam          = "scan-ids"
	QueryIDQueryParam          = "query-id"
	TagsKeyQueryParam          = "tags-keys"
	TagsValueQueryParam        = "tags-values"
	StatusesQueryParam         = "statuses"
	StatusQueryParam           = "status"
	BranchNameQueryParam       = "branch-name"
	ProjectIDQueryParam        = "project-id"
	FromDateQueryParam         = "from-date"
	ToDateQueryParam           = "to-date"
	SeverityQueryParam         = "severity"
	StateQueryParam            = "state"
	GroupQueryParam            = "group"
	QueryQueryParam            = "query"
	NodeIDsQueryParam          = "node-ids"
	IncludeNodesQueryParam     = "include-nodes"
	SortQueryParam             = "sort"
	Profile                    = "default"
	BaseURI                    = ""
	BaseIAMURI                 = ""
	Tenant                     = ""
	Branch                     = ""
	RetrySBOMFlag              = "retry-sbom"
	RetrySBOMDefault           = 1000
	RetrySBOMUsage             = "Retry requests to Checkmarx One on sbom creation"
	ScanPolicyDefaultTimeout   = 4
	ResultPolicyDefaultTimeout = 1
)

// Results
const (
	SastType             = "sast"
	KicsType             = "kics"
	APISecurityType      = "api-security"
	ContainersType       = "containers"
	APIDocumentationFlag = "apisec-swagger-filter"
	IacType              = "iac-security"
	IacLabel             = "IaC Security"
	APISecurityLabel     = "API Security"
	ScaType              = "sca"
	APISecType           = "apisec"
	Success              = "success"
)

// ScaAgent AST Role
const ScaAgent = "SCA_AGENT"

var (
	Version = "dev"
)
