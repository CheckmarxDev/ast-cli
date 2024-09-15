//go:build darwin && amd64

package ASCAconfig

import (
	"github.com/checkmarx/ast-cli/internal/services/osinstaller"
)

var Params = osinstaller.InstallationConfiguration{
	ExecutableFile:  "ASCA_darwin_x64",
	DownloadURL:     "https://download.checkmarx.com/ASCA-binary/ASCA_darwin_x64.tar.gz",
	HashDownloadURL: "https://download.checkmarx.com/ASCA-binary/hash.txt",
	FileName:        "ASCA.tar.gz",
	HashFileName:    "hash.txt",
	WorkingDirName:  "CxASCA",
}