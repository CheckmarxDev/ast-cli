package params

// Flags
const (
	AgentFlag                    = "agent"
	AgentFlagUsage               = "Scan origin name"
	DefaultAgent                 = "ASTCLI"
	DebugFlag                    = "debug"
	DebugUsage                   = "Debug mode with detailed logs"
	RetryFlag                    = "retry"
	RetryDefault                 = 3
	RetryUsage                   = "Retry requests to AST on connection failure"
	RetryDelayFlag               = "retry-delay"
	RetryDelayDefault            = 20
	RetryDelayUsage              = "Time between retries in seconds, use with --" + RetryFlag
	SourcesFlag                  = "file-source"
	SourcesFlagSh                = "s"
	TenantFlag                   = "tenant"
	TenantFlagUsage              = "Checkmarx tenant"
	AsyncFlag                    = "async"
	WaitDelayFlag                = "wait-delay"
	ScanTimeoutFlag              = "scan-timeout"
	SourceDirFilterFlag          = "file-filter"
	SourceDirFilterFlagSh        = "f"
	IncludeFilterFlag            = "file-include"
	IncludeFilterFlagSh          = "i"
	ProjectIDFlag                = "project-id"
	BranchFlag                   = "branch"
	BranchFlagSh                 = "b"
	ScanIDFlag                   = "scan-id"
	BranchFlagUsage              = "Branch to scan"
	MainBranchFlag               = "branch"
	ScaResolverFlag              = "sca-resolver"
	ScaResolverParamsFlag        = "sca-resolver-params"
	AccessKeyIDFlag              = "client-id"
	AccessKeySecretFlag          = "client-secret"
	AccessKeyIDFlagUsage         = "The OAuth2 client ID"
	AccessKeySecretFlagUsage     = "The OAuth2 client secret"
	InsecureFlag                 = "insecure"
	InsecureFlagUsage            = "Ignore TLS certificate validations"
	ScanInfoFormatFlag           = "scan-info-format"
	FormatFlag                   = "format"
	FormatFlagUsageFormat        = "Format for the output. One of %s"
	FilterFlag                   = "filter"
	BaseURIFlag                  = "base-uri"
	ProxyFlag                    = "proxy"
	ProxyFlagUsage               = "Proxy server to send communication through"
	ProxyTypeFlag                = "proxy-auth-type"
	ProxyTypeFlagUsage           = "Proxy authentication type, (basic or ntlm)"
	TimeoutFlag                  = "timeout"
	TimeoutFlagUsage             = "Timeout for network activity, (default 5 seconds)"
	NtlmProxyDomainFlag          = "proxy-ntlm-domain"
	NtlmProxyDomainFlagUsage     = "Window domain when using NTLM proxy"
	BaseURIFlagUsage             = "The base system URI"
	BaseAuthURIFlag              = "base-auth-uri"
	BaseAuthURIFlagUsage         = "The base system IAM URI"
	AstAPIKeyFlag                = "apikey"
	AstAPIKeyUsage               = "The API Key to login to AST"
	ClientRolesFlag              = "roles"
	ClientRolesSh                = "r"
	ClientDescriptionFlag        = "description"
	ClientDescriptionSh          = "d"
	UsernameFlag                 = "username"
	UsernameSh                   = "u"
	PasswordFlag                 = "password"
	PasswordSh                   = "p"
	ProfileFlag                  = "profile"
	ProfileFlagUsage             = "The default configuration profile"
	Help                         = "help"
	TargetFlag                   = "output-name"
	TargetPathFlag               = "output-path"
	TargetFormatFlag             = "report-format"
	ProjectName                  = "project-name"
	ScanTypes                    = "scan-types"
	ScanTypeFlag                 = "scan-type"
	ScanResubmit                 = "resubmit"
	KicsRealtimeFile             = "file"
	KicsRealtimeEngine           = "engine"
	KicsRealtimeAdditionalParams = "additional-params"
	RemediationFiles             = "package-files"
	KicsRemediationFile          = "results-file"
	KicsProjectFile              = "kics-files"
	KicsSimilarityFilter         = "similarity-ids"
	RemediationPackage           = "package"
	RemediationPackageVersion    = "package-version"
	TagList                      = "tags"
	GroupList                    = "groups"
	ProjectGroupList             = "project-groups"
	ProjectTagList               = "project-tags"
	IncrementalSast              = "sast-incremental"
	PresetName                   = "sast-preset-name"
	Threshold                    = "threshold"
	KeyValuePairSize             = 2
	WaitDelayDefault             = 5
	SimilarityIDFlag             = "similarity-id"
	SeverityFlag                 = "severity"
	StateFlag                    = "state"
	CommentFlag                  = "comment"
	LanguageFlag                 = "language"
	VulnerabilityTypeFlag        = "vulnerability-type"
	CweIDFlag                    = "cwe-id"
	SCMTokenFlag                 = "token"
	AzureTokenUsage              = "Azure DevOps personal access token. Requires “Connected server” and “Code“ scope."
	GithubTokenUsage             = "GitHub OAuth token. Requires “Repo” scope and organization SSO authorization, if enforced by the organization"
	GitLabTokenUsage             = "GitLab OAuth token"
	BotCount                     = "Note: dependabot is not counted but other bots might be considered as contributors."
	URLFlag                      = "url"
	GitLabURLFlag                = "url-gitlab"
	URLFlagUsage                 = "API base URL"
	QueryIDFlag                  = "query-id"
	SSHKeyFlag                   = "ssh-key"
	RepoURLFlag                  = "repo-url"
	AstToken                     = "ast-token"
	SSHValue                     = "ssh-value"
	KicsContainerNameKey         = "kics-container-name"
	KicsPlatformsFlag            = "kics-platforms"
	KicsPlatformsFlagUsage       = "KICS Platform Flag"
	IacsPlatformsFlag            = "iacs-platforms"
	IacsPlatformsFlagUsage       = "IaC Security Platform Flag"
	// INDIVIDUAL FILTER FLAGS
	SastFilterFlag  = "sast-filter"
	SastFilterUsage = "SAST filter"
	KicsFilterFlag  = "kics-filter"
	IacsFilterFlag  = "iacs-filter"
	IacsFilterUsage = "IaC Security filter"
	KicsFilterUsage = "KICS filter"
	ScaFilterFlag   = "sca-filter"
	ScaFilterUsage  = "SCA filter"

	// PR decoration flags
	NamespaceFlag      = "namespace"
	NamespaceFlagUsage = "Github namespace is required to post the comments"
	RepoNameFlag       = "repo-name"
	RepoNameFlagUsage  = "Github repository details"
	PRNumberFlag       = "pr-number"
	PRNumberFlagUsage  = "Pull Request number for posting notifications and comments"
)

// Parameter values
const (
	IDQueryParam           = "id"
	IDsQueryParam          = "ids"
	IDRegexQueryParam      = "id-regex"
	LimitQueryParam        = "limit"
	OffsetQueryParam       = "offset"
	ScanIDQueryParam       = "scan-id"
	ScanIDsQueryParam      = "scan-ids"
	QueryIDQueryParam      = "query-id"
	TagsKeyQueryParam      = "tags-keys"
	TagsValueQueryParam    = "tags-values"
	StatusesQueryParam     = "statuses"
	StatusQueryParam       = "status"
	BranchNameQueryParam   = "branch-name"
	ProjectIDQueryParam    = "project-id"
	FromDateQueryParam     = "from-date"
	ToDateQueryParam       = "to-date"
	SeverityQueryParam     = "severity"
	StateQueryParam        = "state"
	GroupQueryParam        = "group"
	QueryQueryParam        = "query"
	NodeIDsQueryParam      = "node-ids"
	IncludeNodesQueryParam = "include-nodes"
	SortQueryParam         = "sort"
	Profile                = "default"
	BaseURI                = ""
	BaseIAMURI             = ""
	Tenant                 = ""
	Branch                 = ""
)

// Results
const (
	SastType = "sast"
	KicsType = "kics"
	IacType  = "iacs"
	IacLabel = "IaC Security"
	ScaType  = "sca"
)

// ScaAgent AST Role
const ScaAgent = "SCA_AGENT"

var (
	Version = "dev"
)
