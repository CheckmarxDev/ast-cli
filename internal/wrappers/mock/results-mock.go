package mock

import (
	"fmt"

	"github.com/checkmarx/ast-cli/internal/wrappers"
)

type ResultsMockWrapper struct{}

var containersResults = &wrappers.ScanResult{
	Type:     "containers",
	Severity: "medium",
	ScanResultData: wrappers.ScanResultData{
		PackageName:       "image-mock",
		PackageVersion:    "1.1",
		ImageName:         "image-mock",
		ImageTag:          "1.1",
		ImageFilePath:     "DockerFile",
		ImageOrigin:       "Docker",
		PackageIdentifier: "mock",
		QueryID:           12.4,
		QueryName:         "mock-query-name",
	},
	Description: "mock-description",
	VulnerabilityDetails: wrappers.VulnerabilityDetails{
		CvssScore: 4.5,
		CveName:   "CVE-2021-1234",
		CweID:     "CWE-1234",
	},
}
var scsResults = &wrappers.ScanResultsCollection{
	TotalCount: 2,
	Results: []*wrappers.ScanResult{
		{
			Type:                 "sscs-Secret Detection",
			ID:                   "bhXbZjjoQZdGAwUhj6MLo9sh4fA=",
			SimilarityID:         "6deb156f325544aaefecee846b49a948571cecd4445d2b2b391a490641be5845",
			Status:               "NEW",
			State:                "TO_VERIFY",
			Severity:             "HIGH",
			Created:              "2024-07-30T12:49:56Z",
			FirstFoundAt:         "2023-07-06T10:28:49Z",
			FoundAt:              "2024-07-30T12:49:56Z",
			FirstScanID:          "3d922bcd-00fe-4774-b182-d51e739dff81",
			Description:          "Generic API Key has detected secret for file application.properties.",
			VulnerabilityDetails: wrappers.VulnerabilityDetails{},
		},
		{
			Type:         "sscs-Scorecard",
			ID:           "n2a8iCzrIgbCe+dGKYk+cAApO0U=",
			SimilarityID: "65323789a325544aaefecee846b49a948571cecd4445d2b2b391a490641be5845",
			Status:       "NEW",
			State:        "TO_VERIFY",
			Severity:     "HIGH",
			Created:      "2024-07-30T12:49:56Z",
			FirstFoundAt: "2023-07-06T10:28:49Z",
			FoundAt:      "2024-07-30T12:49:56Z",
			FirstScanID:  "3d922bcd-00fe-4774-b182-d51e739dff81",
			Description:  "score is 0: branch protection not enabled on development/release branches:\\nWarn: branch protection not enabled for branch 'main'",

			VulnerabilityDetails: wrappers.VulnerabilityDetails{},
		},
	},
}

func (r ResultsMockWrapper) GetAllResultsByScanID(params map[string]string) (
	*wrappers.ScanResultsCollection,
	*wrappers.WebError,
	error,
) {
	if params["scan-id"] == "MOCK_NO_VULNERABILITIES" {
		return &wrappers.ScanResultsCollection{
			TotalCount: 0,
			Results:    nil,
		}, nil, nil
	}
	if params["scan-id"] == "CONTAINERS_ONLY" {
		return &wrappers.ScanResultsCollection{
			TotalCount: 1,
			Results: []*wrappers.ScanResult{
				containersResults,
			},
		}, nil, nil
	}
	if params["scan-id"] == "SCS" {
		return scsResults, nil, nil
	}
	if params["scan-id"] == "NIL_RESULTS" {
		return &wrappers.ScanResultsCollection{
			TotalCount: 0,
			Results:    nil,
		}, nil, nil
	}
	const mock = "mock"
	var dependencyPath = wrappers.DependencyPath{ID: mock, Name: mock, Version: mock, IsResolved: true, IsDevelopment: false, Locations: nil}
	var dependencyArray = [][]wrappers.DependencyPath{{dependencyPath}}
	return &wrappers.ScanResultsCollection{
		TotalCount: 8,
		Results: []*wrappers.ScanResult{
			{
				Type:     "sast",
				ID:       "1",
				Severity: "high",
				ScanResultData: wrappers.ScanResultData{
					LanguageName: "JavaScript",
					QueryName:    "mock-query-name-1",
					Nodes: []*wrappers.ScanResultNode{
						{
							FileName: "dummy-file-name-1",
							Line:     10,
							Column:   10,
							Length:   20,
						},
						{
							FileName: "dummy-file-name-1",
							Line:     11,
							Column:   3,
							Length:   10,
						},
					},
				},
			},
			{
				Type:     "sast",
				ID:       "2",
				Severity: "high",
				ScanResultData: wrappers.ScanResultData{
					LanguageName: "Java",
					QueryName:    "mock-query-name-2",
					Nodes: []*wrappers.ScanResultNode{
						{
							FileName: "dummy-file-name-2",
							Line:     10,
							Column:   10,
							Length:   20,
						},
						{
							FileName: "dummy-file-name-2",
							Line:     11,
							Column:   3,
							Length:   10,
						},
					},
				},
			},
			{
				Type:     "sast",
				Severity: "high",
				ID:       "3",
				ScanResultData: wrappers.ScanResultData{
					LanguageName: "Java",
					QueryName:    "mock-query-name-2",
					Nodes: []*wrappers.ScanResultNode{
						{
							FileName: "dummy-file-name-2",
							Line:     10,
							Column:   10,
							Length:   20,
						},
						{
							FileName: "dummy-file-name-2",
							Line:     11,
							Column:   3,
							Length:   10,
						},
						{
							FileName: "dummy-file-name-2",
							Line:     12,
							Column:   3,
							Length:   10,
						},
					},
				},
			},
			{
				Type:     "sast",
				ID:       "4",
				Severity: "high",
				ScanResultData: wrappers.ScanResultData{
					LanguageName: "Java",
					QueryName:    "mock-query-name-3",
					Nodes: []*wrappers.ScanResultNode{
						{
							FileName: "dummy-file-name-3",
							Line:     10,
							Column:   10,
							Length:   20,
						},
						{
							FileName: "dummy-file-name-3",
							Line:     11,
							Column:   3,
							Length:   10,
						},
					},
				},
			},
			{
				Type:     "sast",
				ID:       "5",
				Severity: "high",
				ScanResultData: wrappers.ScanResultData{
					LanguageName: "Java",
					QueryName:    "mock-query-name-3",
					Nodes: []*wrappers.ScanResultNode{
						{
							FileName: "dummy-file-name-4",
							Line:     10,
							Column:   10,
							Length:   20,
						},
						{
							FileName: "dummy-file-name-4",
							Line:     11,
							Column:   3,
							Length:   10,
						},
					},
				},
			},
			{
				Type:     "sca",
				Severity: "medium",
				ScanResultData: wrappers.ScanResultData{
					ScaPackageCollection: &wrappers.ScaPackageCollection{
						ID:                  "mock",
						FixLink:             "mock",
						Locations:           nil,
						DependencyPathArray: dependencyArray,
						Outdated:            false,
					},
					PackageIdentifier: "mock",
					QueryID:           12.4,
					QueryName:         "mock-query-name",
					Nodes: []*wrappers.ScanResultNode{
						{
							FileName: "dummy-file-name",
							Line:     10,
							Column:   10,
							Length:   20,
						},
						{
							FileName: "dummy-file-name",
							Line:     0,
							Column:   3,
							Length:   10,
						},
					},
				},
			},
			containersResults,
			{
				Type:     "kics",
				Severity: "low",
			},
		},
	}, nil, nil
}

func (r ResultsMockWrapper) GetResultsURL(projectID string) (string, error) {
	return fmt.Sprintf("projects/%s/overview", projectID), nil
}
