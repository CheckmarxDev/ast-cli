package params

// Flags
const (
	AllStatesFlag                = "all"
	AgentFlag                    = "agent"
	AgentFlagUsage               = "Scan origin name"
	ApplicationName              = "application-name"
	DefaultAgent                 = "ASTCLI"
	DebugFlag                    = "debug"
	DebugUsage                   = "Debug mode with detailed logs"
	RetryFlag                    = "retry"
	RetryDefault                 = 3
	RetryUsage                   = "Retry requests to Checkmarx One on connection failure"
	RetryDelayFlag               = "retry-delay"
	RetryDelayDefault            = 20
	RetryDelayPollingDefault     = 60
	RetryDelayUsage              = "Time between retries in seconds, use with --" + RetryFlag
	SourcesFlag                  = "file-source"
	SourcesFlagSh                = "s"
	TenantFlag                   = "tenant"
	TenantFlagUsage              = "Checkmarx tenant"
	AsyncFlag                    = "async"
	WaitDelayFlag                = "wait-delay"
	ScanTimeoutFlag              = "scan-timeout"
	PolicyTimeoutFlag            = "policy-timeout"
	IgnorePolicyFlag             = "ignore-policy"
	SourceDirFilterFlag          = "file-filter"
	SourceDirFilterFlagSh        = "f"
	ImportFilePath               = "import-file-path"
	IncludeFilterFlag            = "file-include"
	IncludeFilterFlagSh          = "i"
	ProjectIDFlag                = "project-id"
	BranchFlag                   = "branch"
	BranchFlagSh                 = "b"
	ScanIDFlag                   = "scan-id"
	CodeRepositoryFlag           = "code-repository-url"
	CodeRepositoryFlagUsage      = "Code repository URL (required for self-hosted SCMs)"
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
	ASCALatestVersion            = "asca-latest-version"
	BaseURIFlag                  = "base-uri"
	ProxyFlag                    = "proxy"
	ProxyFlagUsage               = "Proxy server to send communication through"
	IgnoreProxyFlag              = "ignore-proxy"
	IgnoreProxyFlagUsage         = "Ignore proxy configuration"
	ProxyTypeFlag                = "proxy-auth-type"
	ProxyTypeFlagUsage           = "Proxy authentication type, (basic or ntlm)"
	TimeoutFlag                  = "timeout"
	TimeoutFlagUsage             = "Timeout for network activity, (default 5 seconds)"
	NtlmProxyDomainFlag          = "proxy-ntlm-domain"
	SastFastScanFlag             = "sast-fast-scan"
	NtlmProxyDomainFlagUsage     = "Window domain when using NTLM proxy"
	BaseURIFlagUsage             = "The base system URI"
	BaseAuthURIFlag              = "base-auth-uri"
	BaseAuthURIFlagUsage         = "The base system IAM URI"
	AstAPIKeyFlag                = "apikey"
	AstAPIKeyUsage               = "The API Key to login to Checkmarx One"
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
	ReportFormatPdfToEmailFlag   = "report-pdf-email"
	ReportFormatPdfOptionsFlag   = "report-pdf-options"
	ReportSbomFormatFlag         = "report-sbom-format"
	ProjectName                  = "project-name"
	ScanTypes                    = "scan-types"
	ScanTypeFlag                 = "scan-type"
	ScanResubmit                 = "resubmit"
	KicsRealtimeFile             = "file"
	KicsRealtimeEngine           = "engine"
	KicsRealtimeAdditionalParams = "additional-params"
	ScaRealtimeProjectDir        = "project-dir"
	ScaRealtimeProjectDirSh      = "p"
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
	ThresholdFlagUsage           = "Local build threshold. Format <engine>-<severity>=<limit>. " +
		"Example: scan --threshold \"sast-high=10;sca-high=5;iac-security-low=10\""
	KeyValuePairSize             = 2
	WaitDelayDefault             = 5
	SimilarityIDFlag             = "similarity-id"
	SeverityFlag                 = "severity"
	StateFlag                    = "state"
	CustomStateIDFlag            = "state-id"
	CommentFlag                  = "comment"
	LanguageFlag                 = "language"
	VulnerabilityTypeFlag        = "vulnerability-type"
	CweIDFlag                    = "cwe-id"
	SCMTokenFlag                 = "token"
	AzureTokenUsage              = "Azure DevOps personal access token. Requires “Connected server” and “Code“ scope."
	GithubTokenUsage             = "GitHub Personal Access Token (PAT). Requires “Repo” scope and organization SSO authorization, if enforced by the organization"
	GitLabTokenUsage             = "GitLab OAuth token"
	BitbucketTokenUsage          = "Bitbucket OAuth token"
	BotCount                     = "Note: dependabot is not counted but other bots might be considered as contributors."
	DisabledReposCount           = "Note: Disabled repositories are not counted."
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
	KicsPlatformsFlagUsage       = "KICS Platform Flag. Use ',' as the delimiter for arrays."
	IacsPlatformsFlag            = "iac-security-platforms"
	IacsPlatformsFlagUsage       = "IaC Security Platform Flag"
	ApikeyOverrideFlag           = "apikey-override"
	ExploitablePathFlag          = "sca-exploitable-path"
	LastSastScanTime             = "sca-last-sast-scan-time"
	ProjecPrivatePackageFlag     = "project-private-package"
	SastRedundancyFlag           = "sast-redundancy"
	ContainerImagesFlag          = "container-images"
	ContainersTypeFlag           = "container-security"
	VSCodeAgent                  = "VS Code"
	EclipseAgent                 = "Eclipse"
	VisualStudioAgent            = "Visual Studio"
	JetbrainsAgent               = "Jetbrains"
	ScaPrivatePackageVersionFlag = "sca-private-package-version"
	ScaHideDevAndTestDepFlag     = "sca-hide-dev-test-dependencies"

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
	NamespaceFlag                    = "namespace"
	NamespaceFlagUsage               = "%s namespace is required to post the comments"
	RepoNameFlag                     = "repo-name"
	RepoNameFlagUsage                = "%s repository details"
	PRNumberFlag                     = "pr-number"
	PRNumberFlagUsage                = "Pull Request number for posting notifications and comments"
	PRIidFlag                        = "mr-iid"
	PRIidFlagUsage                   = "Gitlab IID (internal ID) of the merge request"
	PRGitlabProjectFlag              = "gitlab-project-id"
	PRGitlabProjectFlagUsage         = "Gitlab project ID"
	AzureProjectFlag                 = "project"
	AzureProjectFlagUsage            = "Azure project name or project ID"
	CodeRespositoryUsernameFlag      = "code-repository-username"
	CodeRespositoryUsernameFlagUsage = "Azure username for code repository"
	ProjectKeyFlag                   = "project-key"
	ProjectKeyFlagUsage              = "Key of the project containing the repository"
	PRBBIDFlag                       = "pr-id"
	PRBBIDFlagUsage                  = "Bitbucket PR ID"

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

	// SCS Engines Enabled
	SCSEnginesFlag = "scs-engines"

	// SCS (Github)
	SCSRepoTokenFlag = "scs-repo-token"
	SCSRepoURLFlag   = "scs-repo-url"

	// Containers Config Flags
	ContainersFileFolderFilterFlag      = "containers-file-folder-filter"
	ContainersImageTagFilterFlag        = "containers-image-tag-filter"
	ContainersPackageFilterFlag         = "containers-package-filter"
	ContainersExcludeNonFinalStagesFlag = "containers-exclude-non-final-stages"
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
	TagsEmptyQueryParam        = "empty-tags"
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
	ScanPolicyDefaultTimeout   = 4
	ResultPolicyDefaultTimeout = 1
)

// Results
const (
	SastType                       = "sast"
	KicsType                       = "kics"
	APISecurityType                = "api-security"
	AIProtectionType               = "AI Protection"
	ContainersType                 = "containers"
	APIDocumentationFlag           = "apisec-swagger-filter"
	IacType                        = "iac-security"
	IacLabel                       = "IaC Security"
	APISecurityLabel               = "API Security"
	ScaType                        = "sca"
	APISecType                     = "apisec"
	ScsType                        = "scs"
	SscsType                       = "sscs"
	MicroEnginesType               = "microengines" // the scs scan type for scans API
	Success                        = "success"
	SCSScorecardType               = "sscs-scorecard"
	SCSSecretDetectionType         = "sscs-secret-detection"
	EnterpriseSecretsLabel         = "Enterprise Secrets"
	EnterpriseSecretsType          = "enterprise-secrets"
	SCSScorecardOverviewType       = "Scorecard"
	SCSSecretDetectionOverviewType = "2ms"
)

// ScaAgent AST Role
const ScaAgent = "SCA_AGENT"

var (
	Version = "dev"
)

// Custom states
const IncludeDeletedQueryParam = "include-deleted"
const True = "true"

// System States
const ToVerify = "TO_VERIFY"
const NotExploitable = "NOT_EXPLOITABLE"
const ProposedNotExploitable = "PROPOSED_NOT_EXPLOITABLE"
const CONFIRMED = "CONFIRMED"
const URGENT = "URGENT"
