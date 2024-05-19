package wrappers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	commonParams "github.com/checkmarx/ast-cli/internal/params"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

type SastIncrementalHTTPWrapper struct {
	path        string
	contentType string
}

type ResultWithSequence struct {
	Sequence int
	Model    *SastMetadataModel
}

const BatchSize = 200

func NewSastIncrementalHTTPWrapper(path string) SastMetadataWrapper {
	return &SastIncrementalHTTPWrapper{
		path:        path,
		contentType: "application/json",
	}
}

func (s *SastIncrementalHTTPWrapper) getSastMetadata(params map[string]string) (*SastMetadataModel, error) {
	clientTimeout := viper.GetUint(commonParams.ClientTimeoutKey)
	resp, err := SendPrivateHTTPRequestWithQueryParams(http.MethodGet, s.path, params, http.NoBody, clientTimeout)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body)

	defer func() {
		if err == nil {
			_ = resp.Body.Close()
		}
	}()

	switch resp.StatusCode {
	case http.StatusBadRequest, http.StatusInternalServerError:
		errorModel := ErrorModel{}
		err = decoder.Decode(&errorModel)
		if err != nil {
			return nil, fmt.Errorf("%v %s", err, failedToParseGetAll)
		}
		return nil, err
	case http.StatusOK:
		model := SastMetadataModel{}
		err = decoder.Decode(&model)
		if err != nil {
			return nil, fmt.Errorf("%v %s", err, failedToParseGetAll)
		}
		return &model, nil
	case http.StatusNotFound:
		return nil, fmt.Errorf("scan not found")
	default:
		return nil, fmt.Errorf("response status code %d", resp.StatusCode)
	}
}

func (s *SastIncrementalHTTPWrapper) GetSastMetadataByIDs(scanIDs []string) (*SastMetadataModel, error) {
	totalBatches := (len(scanIDs) + BatchSize - 1) / BatchSize
	maxConcurrentGoroutines := 10
	semaphore := make(chan struct{}, maxConcurrentGoroutines)

	var wg sync.WaitGroup
	results := make(chan ResultWithSequence, totalBatches)
	errors := make(chan error, totalBatches)

	for i := 0; i < totalBatches; i++ {
		start := i * BatchSize
		end := start + BatchSize
		if end > len(scanIDs) {
			end = len(scanIDs)
		}

		batchParams := map[string]string{
			commonParams.ScanIDsQueryParam: strings.Join(scanIDs[start:end], ","),
		}

		wg.Add(1)
		semaphore <- struct{}{}
		go func(seq int) {
			defer wg.Done()
			defer func() { <-semaphore }()

			result, err := s.getSastMetadata(batchParams)
			if err != nil {
				errors <- err
				return
			}
			results <- ResultWithSequence{Sequence: seq, Model: result}
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	var allErrors []error
	for err := range errors {
		allErrors = append(allErrors, err)
	}
	if len(allErrors) > 0 {
		return nil, allErrors[0]
	}

	var sortedResults []ResultWithSequence
	for result := range results {
		sortedResults = append(sortedResults, result)
	}
	// sort results by sequence - we need to keep the order of the scans as they were requested
	sortedResults = sortResults(sortedResults)

	finalResult := &SastMetadataModel{}
	for _, result := range sortedResults {
		finalResult.TotalCount += result.Model.TotalCount
		finalResult.Scans = append(finalResult.Scans, result.Model.Scans...)
		finalResult.Missing = append(finalResult.Missing, result.Model.Missing...)
	}

	return finalResult, nil
}

func sortResults(results []ResultWithSequence) []ResultWithSequence {
	slices.SortFunc(results, func(a, b ResultWithSequence) int {
		if a.Sequence < b.Sequence {
			return -1
		}
		if a.Sequence > b.Sequence {
			return 1
		}
		return 0
	})
	return results
}
