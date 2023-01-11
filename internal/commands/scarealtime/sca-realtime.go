package scarealtime

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"fmt"
	"os/exec"

	commonParams "github.com/checkmarx/ast-cli/internal/params"

	"github.com/MakeNowJust/heredoc"
	"github.com/checkmarx/ast-cli/internal/logger"
	"github.com/checkmarx/ast-cli/internal/wrappers"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var scaResolverResultsFileNameDir = ScaResolverWorkingDir + "/cx-sca-realtime-results.json"

const scaResolverProjectName = "cx-cli-sca-realtime-project"
const bitSize = 32

func NewScaRealtimeCommand(scaRealTimeWrapper wrappers.ScaRealTimeWrapper) *cobra.Command {
	scaRealtimeScanCmd := &cobra.Command{
		Use:   "sca-realtime",
		Short: "Create and run sca scan",
		Long:  "The sca-realtime command enables the ability to create, run and retrieve results from a sca scan using sca resolver.",
		// TODO: update example
		Example: heredoc.Doc(
			`
			$ cx scan kics-realtime --file <file> --additional-params <additional-params> --engine <engine>
		`,
		),
		// TODO: update documentation link
		Annotations: map[string]string{
			"command:doc": heredoc.Doc(
				`	
			https://checkmarx.com/resource/documents/en/34965-68643-scan.html#UUID-350af120-85fa-9f20-7051-6d605524b4fc
			`,
			),
		},
		RunE: RunScaRealtime(scaRealTimeWrapper),
	}

	scaRealtimeScanCmd.PersistentFlags().StringP(
		commonParams.ScaRealtimeProjectDir,
		commonParams.ScaRealtimeProjectDirSh,
		"",
		"Path to the project on which SCA Resolver will run",
	)

	err := scaRealtimeScanCmd.MarkPersistentFlagRequired(commonParams.ScaRealtimeProjectDir)
	if err != nil {
		log.Fatal(err)
	}

	return scaRealtimeScanCmd
}

// RunScaRealtime Main method responsible to run sca realtime feature
func RunScaRealtime(scaRealTimeWrapper wrappers.ScaRealTimeWrapper) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, _ []string) error {
		// Validate provided directory
		projectDirPath, err := validateProvidedProjectDirectory(cmd)
		if err != nil {
			return err
		}

		fmt.Println("Running SCA Realtime...")

		// Handle SCA Resolver. Checks if it already exists and if it is in the latest version
		err = downloadSCAResolverAndHashFileIfNeeded(&Params)
		if err != nil {
			return err
		}

		// Run SCA Resolver in the provided directory
		err = executeSCAResolver(projectDirPath)
		if err != nil {
			return err
		}

		// Gets SCA vulnerabilities from SCA APIs
		err = getSCAVulnerabilities(scaRealTimeWrapper)
		if err != nil {
			return err
		}

		return nil
	}
}

// executeSCAResolver Executes sca resolver for a specific path
func executeSCAResolver(projectPath string) error {
	args := []string{
		"offline",
		"-s",
		projectPath,
		"-n",
		scaResolverProjectName,
		"-r",
		scaResolverResultsFileNameDir,
	}

	logger.PrintIfVerbose(fmt.Sprintf("Running SCA resolver with args: %v \n", args))

	_, err := exec.Command(Params.ExecutableFilePath, args...).Output()
	if err != nil {
		return err
	}

	logger.PrintIfVerbose("SCA Resolver finished successfully!")

	return nil
}

// getSCAVulnerabilities Call SCA API to get vulnerabilities from sca resolver results
func getSCAVulnerabilities(scaRealTimeWrapper wrappers.ScaRealTimeWrapper) error {
	scaResolverResults, err := readSCAResolverResultsFromFile()
	if err != nil {
		return err
	}

	var modelResults []wrappers.ScaVulnerabilitiesResponseModel

	for _, dependencyResolutionResult := range scaResolverResults.DependencyResolutionResults {
		// We're using a map to avoid adding repeated packages in request body
		dependencyMap := make(map[string]wrappers.ScaDependencyBodyRequest)

		for i := range dependencyResolutionResult.Dependencies {
			var dependency = dependencyResolutionResult.Dependencies[i]
			dependencyMap[dependency.ID.NodeID] = wrappers.ScaDependencyBodyRequest{
				PackageName:    dependency.ID.Name,
				Version:        dependency.ID.Version,
				PackageManager: dependency.ResolvingModuleType,
			}
			if len(dependency.Children) > 0 {
				for _, dependencyChildren := range dependency.Children {
					dependencyMap[dependencyChildren.NodeID] = wrappers.ScaDependencyBodyRequest{
						PackageName:    dependencyChildren.Name,
						Version:        dependencyChildren.Version,
						PackageManager: dependency.ResolvingModuleType,
					}
				}
			}
		}

		// Get all ScaDependencyBodyRequest from the map to call SCA API
		var bodyRequest []wrappers.ScaDependencyBodyRequest
		for _, value := range dependencyMap {
			bodyRequest = append(bodyRequest, value)
		}

		// We need to call the SCA API for each DependencyResolution so that we can save the file name
		vulnerabilitiesResponseModel, errorModel, errVulnerabilities := scaRealTimeWrapper.GetScaVulnerabilitiesPackages(bodyRequest)
		if errorModel != nil {
			return errors.Errorf("%s: CODE: %d, %s", "An error occurred while getting sca vulnerabilities", errorModel.Code, errorModel.Message)
		}
		if errVulnerabilities != nil {
			return errVulnerabilities
		}

		// Add file name for each vulnerability to display in IDEs
		for _, vulnerability := range vulnerabilitiesResponseModel {
			vulnerability.FileName = dependencyResolutionResult.PackageManagerFile
			modelResults = append(modelResults, vulnerability)
		}
	}

	// Convert SCA Results to Scan Results to make it easier to display it in IDEs
	err = convertToScanResults(modelResults)
	if err != nil {
		return err
	}

	return nil
}

// convertToScanResults Convert SCA Results to Scan Results to make it easier to display it in IDEs
func convertToScanResults(data []wrappers.ScaVulnerabilitiesResponseModel) error {
	var results []*wrappers.ScanResult

	for _, packageData := range data {
		for _, vulnerability := range packageData.Vulnerabilities {
			score, _ := strconv.ParseFloat(vulnerability.Cvss3.BaseScore, bitSize)

			results = append(results, &wrappers.ScanResult{
				Type:        vulnerability.Type,
				ScaType:     "vulnerability",
				Label:       commonParams.ScaType,
				Description: vulnerability.Description,
				Severity:    strings.ToUpper(vulnerability.Severity),
				VulnerabilityDetails: wrappers.VulnerabilityDetails{
					CweID:     vulnerability.Cve,
					CvssScore: score,
					CveName:   vulnerability.Cve,
					CVSS: wrappers.VulnerabilityCVSS{
						Version:            vulnerability.VulnerabilityVersion,
						AttackVector:       vulnerability.Cvss3.AttackVector,
						Availability:       vulnerability.Cvss3.Availability,
						Confidentiality:    vulnerability.Cvss3.Confidentiality,
						AttackComplexity:   vulnerability.Cvss3.AttackComplexity,
						IntegrityImpact:    vulnerability.Cvss3.Integrity,
						Scope:              vulnerability.Cvss3.Scope,
						PrivilegesRequired: vulnerability.Cvss3.PrivilegesRequired,
						UserInteraction:    vulnerability.Cvss3.UserInteraction,
					},
				},
				ScanResultData: wrappers.ScanResultData{
					PackageData: vulnerability.References,
					ScaPackageCollection: &wrappers.ScaPackageCollection{
						FixLink: "https://devhub.checkmarx.com/cve-details/" + vulnerability.Cve,
					},
					Nodes: []*wrappers.ScanResultNode{{
						FileName: packageData.FileName,
					}},
					PackageIdentifier: packageData.PackageName,
				},
			})
		}
	}

	resultsCollection := wrappers.ScanResultsCollection{
		Results:    results,
		TotalCount: uint(len(results)),
	}

	resultsJSON, errs := json.Marshal(resultsCollection)
	if errs != nil {
		return errors.Errorf("%s", errs)
	}
	fmt.Println(string(resultsJSON))

	return nil
}

// validateProvidedProjectDirectory Checks if the provided directory exists in file system
func validateProvidedProjectDirectory(cmd *cobra.Command) (string, error) {
	logger.PrintIfVerbose("Checking if provided project path exists...")
	projectDirPath, _ := cmd.Flags().GetString(commonParams.ScaRealtimeProjectDir)
	pathExists, err := fileExists(projectDirPath)
	if err != nil {
		return "", err
	}

	if !pathExists {
		return "", errors.Errorf("Provided path does not exist: %s", projectDirPath)
	}

	return projectDirPath, nil
}

// readSCAResolverResultsFromFile Get SCA Resolver results from file to build SCA API request body
func readSCAResolverResultsFromFile() (ScaResultsFile, error) {
	file, err := ioutil.ReadFile(scaResolverResultsFileNameDir)
	if err != nil {
		return ScaResultsFile{}, err
	}

	data := ScaResultsFile{}
	_ = json.Unmarshal(file, &data)

	return data, nil
}
