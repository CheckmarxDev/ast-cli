//go:build windows

package scarealtime

import (
	"github.com/checkmarx/ast-cli/internal/services/osinstaller"
)

var Params = osinstaller.InstallableRealTime{
	ExecutableFilePath: "ScaResolver.exe",
	DownloadURL:        "https://sca-downloads.s3.amazonaws.com/cli/latest/ScaResolver-win64.zip",
	HashDownloadURL:    "https://sca-downloads.s3.amazonaws.com/cli/latest/ScaResolver-win64.zip.sha256sum",
	FileName:           "ScaResolver.zip",
	HashFileName:       "ScaResolver.zip.sha256sum",
	WorkingDirName:     "SCARealtime",
}
