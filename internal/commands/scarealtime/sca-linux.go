//go:build linux

package scarealtime

import (
	"github.com/checkmarx/ast-cli/internal/services/osinstaller"
)

var Params = osinstaller.InstallableRealTime{
	ExecutableFilePath: "ScaResolver",
	DownloadURL:        "https://sca-downloads.s3.amazonaws.com/cli/latest/ScaResolver-linux64.tar.gz",
	HashDownloadURL:    "https://sca-downloads.s3.amazonaws.com/cli/latest/ScaResolver-linux64.tar.gz.sha256sum",
	FileName:           "ScaResolver.tar.gz",
	HashFileName:       "ScaResolver.tar.gz.sha256sum",
	WorkingDirName:     "SCARealtime",
}
