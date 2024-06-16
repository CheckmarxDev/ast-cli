package commands

import (
	"path/filepath"

	"github.com/checkmarx/ast-cli/internal/commands/util/printer"
	errorConstants "github.com/checkmarx/ast-cli/internal/constants/errors"
	commonParams "github.com/checkmarx/ast-cli/internal/params"
	"github.com/checkmarx/ast-cli/internal/services"
	"github.com/checkmarx/ast-cli/internal/wrappers"
	"github.com/checkmarx/ast-cli/internal/wrappers/grpcs"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func runScanVorpalCommand(jwtWrapper wrappers.JWTWrapper, featureFlagsWrapper wrappers.FeatureFlagsWrapper) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		vorpalLatestVersion, _ := cmd.Flags().GetBool(commonParams.VorpalLatestVersion)
		fileSourceFlag, _ := cmd.Flags().GetString(commonParams.SourcesFlag)
		agent, _ := cmd.Flags().GetString(commonParams.AgentFlag)
		var port = viper.GetInt(commonParams.VorpalPortKey)
		vorpalWrapper := grpcs.NewVorpalGrpcWrapper(port)
		vorpalParams := services.VorpalScanParams{
			FilePath:            fileSourceFlag,
			VorpalUpdateVersion: vorpalLatestVersion,
			IsDefaultAgent:      agent == commonParams.DefaultAgent,
			JwtWrapper:          jwtWrapper,
			FeatureFlagsWrapper: featureFlagsWrapper,
			VorpalWrapper:       vorpalWrapper,
		}
		scanResult, err := ExecuteVorpalScan(vorpalParams)
		if err != nil {
			return err
		}

		err = printer.Print(cmd.OutOrStdout(), scanResult, printer.FormatJSON)
		if err != nil {
			return err
		}

		return nil
	}
}

func ExecuteVorpalScan(vorpalParams services.VorpalScanParams) (*grpcs.ScanResult, error) {
	if filepath.Ext(vorpalParams.FilePath) == "" && vorpalParams.FilePath != "" {
		return nil, errors.New(errorConstants.FileExtensionIsRequired)
	}
	return services.CreateVorpalScanRequest(vorpalParams)
}
