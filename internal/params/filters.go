package params

var BaseIncludeFilters = []string{
	"*.javasln",
	"*.project",
	"*.java",
	"*.jsp",
	"*.jspf",
	"*.tag",
	"*.tld",
	"*.hbs",
	"*.properties",
	"*.sln",
	"*.csproj",
	"*.cs",
	"*.cshtml",
	"*.xaml",
	"*.vb",
	"*.config",
	"*.asp",
	"*.bas",
	"*.vbp",
	"*.frm",
	"*.cls",
	"*.dsr",
	"*.ctl",
	"*.vb",
	"*.cpp",
	"*.c++",
	"*.cxx",
	"*.hpp",
	"*.hh",
	"*.h++",
	"*.hxx",
	"*.c",
	"*.cc",
	"*.h",
	"*.php",
	"*.php3",
	"*.php4",
	"*.php5",
	"*.php5*6",
	"*.phtm",
	"*.phtml",
	"*.tpl",
	"*.ctp",
	"*.twig",
	"*.apex",
	"*.apexp",
	"*.page",
	"*.component",
	"*.cls",
	"*.trigger",
	"*.tgr",
	"*.object",
	"*.report",
	"*.workflow",
	"*.rb",
	"*.rhtml",
	"*.rxml",
	"*.rjs",
	"*.erb",
	"*.js",
	"*.htm",
	"*.html",
	"*.json",
	"*.ts",
	"*.tsx",
	"*.vbs",
	"*.pl",
	"*.pm",
	"*.plx",
	"*.psgi",
	"*.java",
	"*.kt",
	"*.m",
	"*.h",
	"*.swift",
	"*.xib",
	"*.html",
	"*.htm",
	"*.pls",
	"*.sql",
	"*.pkh",
	"*.pks",
	"*.pkb",
	"*.pck",
	"*.py",
	"*.groovy",
	"*.gsh",
	"*.gvy",
	"*.gy",
	"*.scala",
	"*.sc",
	"*.conf",
	"*.go",
	"*.kt",
	"*.kts",
	"*.cbl",
	"*.cob",
	"*.eco",
	"*.pco",
	"*.sqb",
	"*.cpy",
	"*.aspx",
	"*.ascx",
	"*.config",
	"*.xml",
	"*.cgi",
	"*.inc",
	"*.jar",
	"*.js",
	"*.dll",
	"*.tf",
	"*.yaml",
	"*.yml",
	"*.gradle",
	"gradlew",
	"build.gradle",
	"build.sbt",
	"yarn.lock",
	"requirements.txt",
	"requirement.txt",
	"requirement*.txt",
	"composer.lock",
	"Dockerfile*",
	"dock*",
	"*.dart",
	"*.plist",
	"go.mod",
	"go.sum",
	"Podfile",
	"Podfile.lock",
	"*.cmp",
	"Directory.Packages.props",
}

var BaseExcludeFilters = []string{
	"!.vs",
	"!.vscode",
	"!.idea",
}

var KicsBaseFilters = []string{
	".tf",
	".yaml",
	".yml",
	".json",
	".auto.tfvars",
	".terraform.tfvars",
	"Dockerfile",
	".proto",
	".dockerfile",
}

var DisabledExclusions = map[string]bool{
	".git": true,
}
