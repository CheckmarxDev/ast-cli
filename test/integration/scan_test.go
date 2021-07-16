// +build integration

package integration

import (
	"context"
	"fmt"
	"github.com/checkmarxDev/ast-cli/internal/commands"
	"github.com/google/uuid"
	"log"
	"testing"
	"time"

	scansApi "github.com/checkmarxDev/scans/pkg/api/scans"
	scansRESTApi "github.com/checkmarxDev/scans/pkg/api/scans/rest/v1"
	"gotest.tools/assert"
)

// Sources to use in scans
const (
	sources    = "sources.zip"
	sourcesInc = "sources_inc.zip"
	fastRepo   = "https://github.com/CheckmarxDev/ast-cli"
	slowRepo   = "https://github.com/WebGoat/WebGoat"
)

// Type for scan workflow response, used to assert the validity of the command's response
type ScanWorkflowResponse struct {
	Source      string    `json:"source"`
	Timestamp   time.Time `json:"timestamp"`
	Information string    `json:"info"`
}

// Create scans from zip and url and perform assertions in executeScanTest
func TestScansE2E(t *testing.T) {

	for _, source := range []string{sources, fastRepo} {
		scanID, projectID := createScan(t, source, Tags)

		executeScanTest(t, projectID, scanID, Tags)
	}
}

// Perform a nowait scan and poll status until completed
func TestNoWaitScan(t *testing.T) {
	scanID, projectID := createScanNoWait(t, sources, map[string]string{})

	assert.Assert(t, pollScanUntilStatus(t, scanID, scansApi.ScanCompleted, FullScanWait, ScanPollSleep), "Polling should complete")

	executeScanTest(t, projectID, scanID, Tags)
}

// Perform an initial scan with complete sources and an incremental scan with a smaller wait time
func TestIncrementalScan(t *testing.T) {
	scanID, projectID := createScanIncremental(t, sources, map[string]string{})

	defer deleteScan(t, scanID)

	assert.Assert(t, pollScanUntilStatus(t, scanID, scansApi.ScanCompleted, FullScanWait, ScanPollSleep), "Polling should complete")

	executeScanTest(t, projectID, scanID, Tags)

	scanID, projectID = createScanIncremental(t, sourcesInc, map[string]string{})

	defer deleteScan(t, scanID)

	assert.Assert(t, pollScanUntilStatus(t, scanID, scansApi.ScanCompleted, IncScanWait, ScanPollSleep), "Polling should complete")

	executeScanTest(t, projectID, scanID, Tags)
}

// Get a scan workflow and assert its structure
func TestScanWorkflow(t *testing.T) {
	scanID, projectID := createScanNoWait(t, sources, map[string]string{})

	defer deleteProject(t, projectID)
	defer deleteScan(t, scanID)

	workflowCommand, buffer := createRedirectedTestCommand(t)

	err := execute(workflowCommand,
		"scan", "workflow",
		flag(commands.ScanIDFlag), scanID,
		flag(commands.FormatFlag), commands.FormatJSON,
	)
	assert.NilError(t, err, "Workflow should pass")

	var workflow []ScanWorkflowResponse
	_ = unmarshall(t, buffer, &workflow, "Reading workflow output should work")

	assert.Assert(t, len(workflow) > 0, "At least one item should exist in the workflow response")
}

// Start a scan guaranteed to take considerable time, cancel it and assert the status
func TestCancelScan(t *testing.T) {
	scanID, projectID := createScanNoWait(t, slowRepo, map[string]string{})

	defer deleteProject(t, projectID)
	defer deleteScan(t, scanID)

	workflowCommand := createASTIntegrationTestCommand(t)

	err := execute(workflowCommand,
		"scan", "cancel",
		flag(commands.ScanIDFlag), scanID,
	)
	assert.NilError(t, err, "Cancel should pass")

	scan := showScan(t, scanID)

	assert.Assert(t, scan.Status == scansApi.ScanCanceled, "Scan should be canceled")
}

// Generic scan test execution
// - Get scan with 'scan list' and assert status and IDs
// - Get scan with 'scan show' and assert the ID
// - Assert all tags exist and are assigned to the scan
// - Delete the scan and assert it is deleted
func executeScanTest(t *testing.T, projectID string, scanID string, tags map[string]string) {
	defer deleteProject(t, projectID)

	response := listScanByID(t, scanID)

	assert.Equal(t, len(response), 1, "Total scans should be 1")
	assert.Equal(t, response[0].ID, scanID, "Scan ID should match the created scan's ID")
	assert.Equal(t, response[0].ProjectID, projectID, "Project ID should match the created scan's project ID")
	assert.Assert(t, response[0].Status == scansApi.ScanCompleted, "Scan should be completed")

	scan := showScan(t, scanID)
	assert.Equal(t, scan.ID, scanID, "Scan ID should match the created scan's ID")

	allTags := getAllTags(t, "scan")
	for key := range tags {
		_, ok := allTags[key]
		assert.Assert(t, ok, "Get all tags response should contain all created tags. Missing %s", key)

		val, ok := scan.Tags[key]
		assert.Assert(t, ok, "Scan should contain all created tags. Missing %s", key)
		assert.Equal(t, val, Tags[key], "Tag value should be equal")
	}

	deleteScan(t, scanID)

	response = listScanByID(t, scanID)

	assert.Equal(t, len(response), 0, "Total scans should be 0 as the scan was deleted")
}

func createScan(t *testing.T, source string, tags map[string]string) (string, string) {
	return executeCreateScan(t, getCreateArgs(source, tags))
}

func createScanNoWait(t *testing.T, source string, tags map[string]string) (string, string) {
	return executeCreateScan(t, append(getCreateArgs(source, tags), "--nowait"))
}

func createScanIncremental(t *testing.T, source string, tags map[string]string) (string, string) {
	return executeCreateScan(t, append(getCreateArgs(source, tags), "--sast-incremental"))
}

func getCreateArgs(source string, tags map[string]string) []string {
	projectName := fmt.Sprintf("integration_test_scan_%s", uuid.New().String())
	args := []string{
		"scan", "create",
		flag(commands.ProjectName), projectName,
		flag(commands.SourcesFlag), source,
		flag(commands.ScanTypes), "sast,kics,sca",
		flag(commands.PresetName), "Checkmarx Default",
		flag(commands.FormatFlag), commands.FormatJSON,
		flag(commands.TagList), formatTags(tags),
	}
	return args
}

func executeCreateScan(t *testing.T, args []string) (string, string) {
	createCommand, buffer := createRedirectedTestCommand(t)

	err := execute(createCommand, args...)
	assert.NilError(t, err, "Creating a scan should pass")

	createdScan := scansRESTApi.ScanResponseModel{}
	_ = unmarshall(t, buffer, &createdScan, "Reading scan response JSON should pass")

	assert.Assert(t, createdScan.Status == scansApi.ScanQueued)

	log.Printf("Scan ID %s created in test", createdScan.ID)

	return createdScan.ID, createdScan.ProjectID
}

func deleteScan(t *testing.T, scanID string) {
	deleteScanCommand := createASTIntegrationTestCommand(t)
	err := execute(deleteScanCommand,
		"scan", "delete",
		flag(commands.ScanIDFlag), scanID,
	)
	assert.NilError(t, err, "Deleting a scan should pass")
}

func listScanByID(t *testing.T, scanID string) []scansRESTApi.ScanResponseModel {
	scanFilter := fmt.Sprintf("scan-ids=%s", scanID)

	getCommand, outputBuffer := createRedirectedTestCommand(t)
	err := execute(getCommand,
		"scan", "list",
		flag(commands.FormatFlag), commands.FormatJSON,
		flag(commands.FilterFlag), scanFilter,
	)
	assert.NilError(t, err, "Getting the scan should pass")

	// Read response from buffer
	var scanList []scansRESTApi.ScanResponseModel
	_ = unmarshall(t, outputBuffer, &scanList, "Reading scan response JSON should pass")

	return scanList
}

func showScan(t *testing.T, scanID string) scansRESTApi.ScanResponseModel {

	getCommand, outputBuffer := createRedirectedTestCommand(t)

	err := execute(getCommand,
		"scan", "show",
		flag(commands.FormatFlag), commands.FormatJSON,
		flag(commands.ScanIDFlag), scanID,
	)
	assert.NilError(t, err, "Getting the scan should pass")

	// Read response from buffer
	scan := scansRESTApi.ScanResponseModel{}
	_ = unmarshall(t, outputBuffer, &scan, "Reading scan response JSON should pass")

	return scan
}

func pollScanUntilStatus(t *testing.T, scanID string, requiredStatus scansApi.ScanStatus, timeout, sleep int) bool {
	log.Printf("Set timeout of %d seconds for the scan to complete...\n", timeout)
	// Wait for the scan to finish. See it's completed successfully
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return false
		default:
			log.Printf("Polling scan %s\n", scanID)
			scan := listScanByID(t, scanID)
			if s := string(scan[0].Status); s == string(requiredStatus) {
				return true
			} else if s == scansApi.ScanFailed || s == scansApi.ScanCanceled ||
				s == scansApi.ScanCompleted {
				return false
			} else {
				time.Sleep(time.Duration(sleep) * time.Second)
			}
		}
	}
}
