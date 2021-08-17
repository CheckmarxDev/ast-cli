[![CircleCI](https://circleci.com/gh/CheckmarxDev/ast-cli/tree/master.svg?style=svg&circle-token=32eeef7505db60c11294e63db64e70882bde83b0)](https://circleci.com/gh/CheckmarxDev/ast-cli/tree/master)
## Building from source code
### Windows 
``` powershell
setx GOOS=windows 
setx GOARCH=am
go build -o ./bin/cx.exe ./cmd
```

### Linux 

``` bash
export GOARCH=amd64
export GOOS=linux
go build -o ./bin/cx ./cmd
```

### Macintosh

``` bash
export GOOS=darwin 
export GOARCH=amd64
go build -o ./bin/cx-mac ./cmd
```

** **

## Contribution

We appreciate feedback and contribution to the CLI! Before you get started, please see the following:

- [Checkmarx contribution guidelines](docs/contributing.md)
- [Checkmarx Code of Conduct](docs/code_of_conduct.md)

** **

## Basic CLI Operation

This document provides many examples of using the AST CLI but it's impossible to cover every possible action. You can  you can always fall back to the (--help or -h) option, ex:

### Windows

``` powershell
cx.exe --help
```

### Linux/Mac
``` bash
./cx -h
```

You will notice help shows a list of available commands and a summary of global parameters. The (--help) command also lets you dig into commands for more context specific help, ex:

``` bash
# Show help for the scan command
./cx scan -h
Manage scans

Usage:
  cx scan [command]

Available Commands:
  cancel      Cancel one or more scans from running
  create      Create and run a new scan
  delete      Deletes one or more scans
  list        List all scans in the system
  show        Show information about a scan
  tags        Get a list of all available tags to filter by
  workflow    Show information about a scan workflow
  
# At this point you can dig into the "create" command
./cx scan create -h
Create and run a new scan

Usage:
  cx scan create [flags]

Flags:
  -b, --branch string             Branch to scan
  -f, --file-filter string        Source file filtering pattern
  -i, --file-include string       Comma separated list of extra inclusions, ex: *zip,file.txt
  -s, --file-source string        Sources like: directory, zip file or git URL.
      --format string             Format for the output. One of [table json list] (default "list")
  -h, --help                      help for create
      --nowait                    Wait for scan completion (default true)
      --project-name string       Name of the project
      --sast-incremental string   Incremental SAST scan should be performed. (default "false")
      --sast-preset-name string   The name of the Checkmarx preset to use.
      --scan-types string         Scan types, ex: (sast,kics,sca)
      --tags string               List of tags, ex: (tagA,tagB:val,etc)
      --wait-delay int            Polling wait time in seconds (default 5)     
      ....
```

You may have noticed the parameters accepted by the CLI vary based on the commands issued but the following parameters are available throughout the CLI command hierarchy:

- (--base-uri), the URL of the AST server.
- (--base-auth-uri), optionally provides alternative KeyCloak endpoint to (--base-uri).
- (--client-id), the client ID used for authentication (see Authentication documentation).
- (--client-secret), the secret that corrosponds to the client-id  (see Authentication documentation).
- (--apikey), the token to authenticate with (see Authentication documentation).
- (--proxy), optional proxy server to use (see Proxy Support documentation).
- (--insecure), indicates CLI should ignore TLS certificate validations.
- (--tenant), your AST tenant name.
- (--profile), specifies the CLI profile to store options in (see Environment and Configuration documentation).

Many CLI variables can be provided using environment variables, configuration variables or CLI parameters. The follow precidence is used when the same value is found in settings:

1. CLI parameters, these always overide configuration and environment variables.
2. Configuration variables always overide environment variables.
3. Environment variables are the first order precidence. 

## CLI Configurations

The CLI allows you to permanently store some CLI options in configuration files. The configuration files are kept in the users home directory under a subdirectory named  ($HOME/.checkmarx). 

``` bash
./cx configure set --prop-name cx_base_uri --prop-value "http://<your-server>[:<port>]"
./cx configure set --prop-name cx_base_auth_uri --prop-value "http://<your-server>[:<port>]"
./cx configure set --prop-name cx_client_id --prop-value <your-key>
./cx configure set --prop-name cx_client_secret --prop-value <your-secret>
./cx configure set --prop-name http_proxy --prop-value <your-proxy>
./cx configure set --prop-name cx_apikey --prop-value <your-apikey>
./cx configure set --prop-name cx_tenant --prop-value <your-tenant-name>
```

The (--profile) option provides a powerful tool to quickly switch through differnent sets of configurations. You can add a profile name to any CLI command and it utilize the corsponding profile settings. The following example setups up altnerative profile named "test" and calls a CLI command utilizing it.

``` bash
# Create the alternative profile
./cx configure set --prop-name cx_base_uri --prop-value "http://your-test-server" --profile test
# Use the cx_base_uri from the "test" profile.
./cx scan list --profile test
# This uses the default profile (if it exists)
./cx scan list
```

The configure command supports an interactive mode that prompt you for the following common options: base-uri, client ID and secret, ex:

``` bash
./cx configure
AST Base URI [http://<your-domain]: <your-updated-domain>
AST Base Auth URI [<your-updated-domain>]: <your-updated-auth-domain>
AST Tenant [organization]: <your-tenant-name>
Do you want to use API Key authentication? (Y/N): n
AST Client ID  [******f23d]: <your-updated-client-idy>
AST Client Secret [******8913]: <your-updated-client-secret>
```

If the CLI has previously stored values they will show up like you see in the previous example, you can just press enter if you want to keep the existing value.

These values can be stored in CLI configurations:

- cx_apikey: the apikey to authenticate with (see Authentication documentation).
- cx_tenant: your AST tenant name.
- cx_base_uri: the URL of the AST server.
- cx_base_auth_uri: the URL of the AST server.
- http_proxy: optional proxy server to use (see Proxy Support documentation). 
- cx_client_id: the client ID used for authentication (see Authentication documentation).
- cx_client_secret: the secret that corrosponds to the client-id  (see Authentication documentation).

## Authentication

The CLI supports apikey and key/secret based authentication.

API Key based authentication is the easiest method to use in CI/CD environments. Tokens are generated through KeyCloak and can be created with a predictable lifetime. Once you have a token you can use it from the CLI like this:

``` bash
./cx --apikey <your-apikey> scan list 
```

You can optionally configure the token into your stored CLI configuration values like this:

``` bash
./cx configure set cx_apikey <your-apikey>
# The following command will automatically use the stored apikey
./cx scan list
```

You can also store the apikey in the environment like this:

``` bash
export CX_APIKEY=<your-apikey>
./cx scan list
```

Key/secret authentication requires you to first use an AST username and password to create the key and secret for authentication. The following example shows how to create a key/secret and then use it:

``` bash
./cx auth register -u <username> -p <password>
CX_CLIENT_ID=<generated-client-id>
CX_CLIENT_SECRET=<generated-client-secret>
```

Once you generated your key and secret they can be used like this:

``` bash
./cx --client-id <your-client-id> --secret <your-client-secret> scan list 
```

You can optionally configure the key/secret into your stored CLI configuration values like this:

``` bash
./cx configure set cx_client_id <your-client-id>
./cx configure set cx_client_secret <your-client-secret>
# The following command will automatically use the stored key/secret
./cx scan list
```

You can also store the key/secret in the environment like this:

``` bash
export CX_CLIENT_ID=<your-client-id>
export CX_CLIENT_SECRET=<your-client-secret>
./cx scan list
```

## Tenant Configuration

If you're are using a hosted AST instance you will need to include your tenant name when using the CLI. Failure to configure your tenant information correctly will result in communication failures with AST. 

The tenant name can be provided in three different ways:

1. As a CLI argument (--tenant <your-tenant-name>).
2. As an environment variable (CX_TENANT=<your-tenant-name>).
3. Through the (cx configure) action. 

## Command Completion

In most console environments the CLI supports command completion. In order for command completion to work you need configure your console correctly. You can skip the following setup if you have installed the CLI from a prepacked source like Home Brew, Apt, Chocolately, etc. 

Bash command completion setup:

``` bash
$ source < (cx utils completion --shell bash)
# To load completions for each session, execute once:
# Linux:
$ cx utils completion --shell bash > /etc/bash_completion.d/cx
# macOS:
$ cx utils completion --shell bash > /usr/local/etc/bash_completion.d/cx
```

Zsh command completion setup:

``` bash
# If shell completion is not already enabled in your environment,
# you will need to enable it.  You can execute the following once:	
$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ cx utils completion --shell zsh > "${fpath[1]}/_cx"	
# You will need to start a new shell for this setup to take effect.
```

PowerShell command completion setup:

``` powershell
PS> cx utils completion --shell powershell | Out-String | Invoke-Expression
	
# To load completions for every new session, run:
PS> cx utils completion --shell powershell > cx.ps1
# and source this file from your PowerShell profile.
```

Once command completion has been setup you simply press [tab] [tab] to get a list of supported commands and arguments. The result you see will depend on the context when you press [tab] [tab]. The following example shows the results when using the (project list) command:

``` bash
./cx proje [tab]
# completes to
./cx project [tab]
# shows the following options
create  -- Creates a new project
delete  -- Delete a project
list    -- List all projects in the system
show    -- Show information about a project
tags    -- Get a list of all available tags
```



## Triggering Scans

You need to specify a project using the (--project-name) parameter when you create a scan. If the project doesn't exist then it will be created automatically; however, if the project exists it will be reused. The following examples will all use the same project name for simplicity.

You can optionally specify the name of the preset to use when scanning projects using the (--sast-preset-name) parameter. If you don't specify the preset name then "Checkmarx Default" will be used. 

You can indicate if an incremental or full scan should be performed with the (--sast-incremental) parameter. If you don't provide the incremental flag then a full scan will be triggered.

The (--scan-types) parameter is used to indicate which types of scan should be performed by AST. You can provide a comma separated list of scan types if you want multiple scans to be performed, but you specify at least one scan type. The following scan types are available:

- SAST
- SCA
- KICS

The most important thing you need to decide is where the scan is going to come from. The type of parameter you pass to the (--file-source) option will be examined, and the type of scan being used will automatically be determined, you have the following options: 

1. A zip file with your source code. The CLI will assume any source that ends with (.zip) is contains zipped source files.
3. A host git repo. The CLI will assume that any source that starts with (http://) or (https://) points to a GIT repository.
3. A directory with your source code. If this option is specified the (--file-filter) option can be specified to control what is sent to AST.

**NOTE**: for simplicity the following examples assume you have stored your authentication and base-uri information in either environment variables or CLI configuration parameters. These values are required but will not appear in the commands.

**NOTE**: to show different real world situations optional parameters will sometimes but not always be used.

After you create a scan the CLI will wait until it is completed or an error has been encountered. The default polling interval is 5 seconds but you can overide that with the (--wait-delay) option. You can also turn off the wait mode with (--nowait true).

Scanning zipped code archives can be achieved like this:

``` bash
./cx scan create -s <your-file>.zip --project-name "testproj" --sast-preset-name "Checkmarx Default" --sast-incremental "false" --scan-types "sast"
```

If you decide to scan a local directory you can provide filters that determine which resources are sent to the AST server.
For more details, refer to the [filtering](#filtering) section.
The following example shows how to scan a folder. The filter in this case will include any go files that start with an 's':

``` bash
./cx scan create -s <path-to-your-folder> -f "s*.go" --sast-project-name "testproj" --sast-incremental "false" --project-type "sast" --scan-types "sast,sca,kics"
```

Git repositories can be scanned like this:

``` bash
./cx scan create -s <your-repo-url> --project-name "testproj" --scan-types "sast" 
```

When you're scanning repos AST will fetch the code directly from the repository.

You can specify the name of branch to scan to can with the (--branch) option. If the (--file-source) type indicates a GIT repository the (--branch) value will also indicate the branch to pull code from. The (--branch) option can be used like this:

``` bash
./cx scan create -s <your-repo-url> --project-name "testproj" --scans-type "sast" --branch yourBranch
```

You can disable polling mode like this:

``` bash
./cx scan create -s <your-repo-url> --project-name "testproj" --scan-types "sast" --nowait true
```

You can view the information about a scan like this:

``` bash
./cx scan show --scan-id <scan-id>
```

You can cancel running scans like this:

``` bash
# Cancel a single scan
./cx scan cancel --scan-id <scan-id>
# Cancel multiple scans at once
./cx scan cancel --scan-id <scan-id>,<scan-id>,etc
```

You can delete existing scan like this:

``` bash
# Delete a single scan
./cx scan delete --scan-id <scan-id>
# Cancel multiple scans at once
./cx scan delete --scan-id <scan-id>,<scan-id>,etc
```

### Filtering

The filters are based on an inclusion and exclusion model.
There is a default inclusion list so only files scannable by AST are included.
Filters you provide with the filtering flag will be applied on top of the default inclusion list.
Building on the previous example:

You can include more than one set of files and directories by separating the inclusion patterns with a comma, example:

``` bash
./cx scan create -s <path-to-your-folder> -f "s*,f*" --project-name "testproj" --sast-preset-name "Checkmarx Default" --sast-incremental "false" --scan-types "sast,kics"
```

In this previous example any files that start with 's' will be included, as well as any files that start with 'f'. You can add an exclusion into the list by prepending the pattern with a '!'. The following query demonstrates exclusion by filtering files that end with 'dat':

``` bash
./cx scan create -s <path-to-your-folder> -f "s*,f*,!*dat" --project-name "testproj" --sast-preset-name "Checkmarx Default" --sast-incremental "false" --scan-types "sast" 
```

You can also add extra patterns to the inclusion list so more of your files are included before applying the filters.
If you want to include zip files and a configuration file called project.cfg, which are typically not included:

``` bash
./cx scan create -s <path-to-your-folder> -f "s*,f*,!*dat" -i "*.zip,project.cfg" --project-name "testproj" --sast-preset-name "Checkmarx Default" --sast-incremental "false" --scan-types "sast" 
```

## Managing Projects

You can create, delete, list or show details about AST projects using the CLI. You specifically create projects before triggering scans though, the (scan create) will automatically create projects that don't exist for you. The commands just provide a help way to work with projects.

You can create projects like this:

``` bash
./cx project create --branch "test" --project-name "createTest" --repo-url "https://github.com/tsunez/checkmarxTest.git"

Project ID       Name       Created at          Updated at          Tags Groups 
----------       ----       ----------          ----------          ---- ------ 
56939423....    createTest  03-22-21 08:27:34   03-22-21 08:27:34   []   [] 
```

The only required parameter is (--project-name). If the (--repo-url) or (--branch) don't make sense for your purposes then just skip them.

You can list existing projects like this:

``` bash
./cx project list

Project ID       Name       Created at          Updated at          Tags Groups 
----------       ----       ----------          ----------          ---- ------ 
56939423....    createTest  03-22-21 08:27:34   03-22-21 08:27:34   []   [] 
```

You can show the details about a specific project like this:

``` bash
./cx project show --project-id <your-project-id>

Project ID       Name       Created at          Updated at          Tags Groups 
----------       ----       ----------          ----------          ---- ------ 
56939423....    createTest  03-22-21 08:27:34   03-22-21 08:27:34   []   [] 
```

Finally you can delete a project like this:

``` bash
./cx project delete --project-id <your-project-id>
```



## Retreiving Results

...todo, we're still waiting to work out result retreival

## Proxy Support

The CLI full supports proxy servers and optional proxy server authentication. When the proxy server variable is found all CLI operations will be routed through the target server. Proxy server URLs should like this: "http[s]://your-server.com:[port]"

Proxy support is enabled by creating an environment variable named HTTP_PROXY. You can also specify the proxy by storing the proxy URL in a CLI configuration variable (proxy_http), or directly through with the CLI with the parameter  (--proxy).

The following example demonstraights the use of a proxy server:

``` bash
./cx scan list --proxy "http://<your-proxy>:8081"
```

The AST CLI supports proxies that provide that use NTLM based authentication. The following example shows the use of an NTLM proxy:

``` bash
./cx scan list --proxy "http://<ntlm-user>:<ntlm-password>@<your-proxy>:8081" --proxy-auth-type ntlm --proxy-ntlm-domain "<your-domain>"
```

If you choose to use a proxy that uses NTLM authentication you will need to provide a username, password and domain name for your Active Directory.

## Environment Variables

| Environment Variable | Description                                                  |
| -------------------- | ------------------------------------------------------------ |
| **CX_CLIENT_ID**     | Key portion of key/secret authentication pair.               |
| **CX_CLIENT_SECRET** | Secret portion of key/secret authentication pair.            |
| **CX_APIKEY**        | API Key for token based authentication.                      |
| **CX_TENANT**        | Your AST tenant name.                                        |
| **CX_BASE_URI**      | The URI of the AST server.                                   |
| **CX_BASE_AUTH_URI** | The URI of KeyCloak instance. This optional and only required when you're not using AST's built in KeyCloak instance. |
| **HTTP_PROXY**       | When provided this variable will trigger the CLI to use the proxy server pointed to (see proxy support documentation). |
| **CX_BRANCH**        | The branch to scan.                                          |

** **

## License

This repo is covered under [Apache 2.0](LICENSE).
