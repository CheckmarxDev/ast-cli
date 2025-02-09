//go:build !integration

package commands

import (
	"fmt"
	"github.com/checkmarx/ast-cli/internal/wrappers"
	"github.com/checkmarx/ast-cli/internal/wrappers/mock"
	"testing"

	"gotest.tools/assert"
)

func TestTriageHelp(t *testing.T) {
	execCmdNilAssertion(t, "help", "triage")
}

func TestRunShowTriageCommand(t *testing.T) {
	execCmdNilAssertion(t, "triage", "show", "--project-id", "MOCK", "--similarity-id", "MOCK", "--scan-type", "sast")
}

func TestRunUpdateTriageCommand(t *testing.T) {
	execCmdNilAssertion(
		t,
		"triage",
		"update",
		"--project-id",
		"MOCK",
		"--similarity-id",
		"MOCK",
		"--state",
		"confirmed",
		"--comment",
		"Testing commands.",
		"--severity",
		"low",
		"--scan-type",
		"sast")
}

func TestRunShowTriageCommandWithNoInput(t *testing.T) {
	err := execCmdNotNilAssertion(t, "triage", "show")
	assert.Assert(t, err.Error() == "required flag(s) \"project-id\", \"scan-type\", \"similarity-id\" not set")
}

func TestRunUpdateTriageCommandWithNoInput(t *testing.T) {
	err := execCmdNotNilAssertion(t, "triage", "update")
	fmt.Println(err)
	assert.Assert(
		t,
		err.Error() == "required flag(s) \"project-id\", \"scan-type\", \"severity\", \"similarity-id\", \"state\" not set")
}

func TestTriageGetStatesFlag(t *testing.T) {
	mockWrapper := &mock.CustomStatesMockWrapper{}
	cmd := triageGetStatesSubCommand(mockWrapper)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	assert.NilError(t, err)
	states, err := mockWrapper.GetAllCustomStates(false)
	assert.NilError(t, err)
	assert.Equal(t, len(states), 2)
	cmd.SetArgs([]string{"--all"})
	err = cmd.Execute()
	assert.NilError(t, err)
	states, err = mockWrapper.GetAllCustomStates(true)
	assert.NilError(t, err)
	assert.Equal(t, len(states), 3)
}

func TestGetCustomStateID(t *testing.T) {
	tests := []struct {
		name                string
		state               string
		mockWrapper         wrappers.CustomStatesWrapper
		expectedStateID     string
		expectedErrorString string
	}{
		{
			name:            "State found",
			state:           "demo3",
			mockWrapper:     &mock.CustomStatesMockWrapper{},
			expectedStateID: "3",
		},
		{
			name:                "State not found",
			state:               "nonexistent",
			mockWrapper:         &mock.CustomStatesMockWrapper{},
			expectedStateID:     "",
			expectedErrorString: "No matching state found for nonexistent",
		},
		{
			name:                "Error fetching states",
			state:               "nonexistent",
			mockWrapper:         &mock.CustomStatesMockWrapperWithError{},
			expectedStateID:     "",
			expectedErrorString: "Failed to fetch custom states",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			stateID, err := getCustomStateID(tt.mockWrapper, tt.state)
			if tt.expectedErrorString != "" {
				assert.ErrorContains(t, err, tt.expectedErrorString)
			} else {
				assert.NilError(t, err)
			}
			assert.Equal(t, stateID, tt.expectedStateID)
		})
	}
}

func TestIsCustomState(t *testing.T) {
	tests := []struct {
		state    string
		isCustom bool
	}{
		{"TO_VERIFY", false},
		{"to_verify", false},
		{"NOT_EXPLOITABLE", false},
		{"PROPOSED_NOT_EXPLOITABLE", false},
		{"CONFIRMED", false},
		{"URGENT", false},
		{"CUSTOM_STATE_1", true},
		{"CUSTOM_STATE_2", true},
	}

	for _, tt := range tests {
		t.Run(tt.state, func(t *testing.T) {
			result := isCustomState(tt.state)
			assert.Equal(t, result, tt.isCustom)
		})
	}
}
func TestRunTriageUpdateWithNotFoundCustomState(t *testing.T) {
	mockResultsPredicatesWrapper := &mock.ResultsPredicatesMockWrapper{}
	mockFeatureFlagsWrapper := &mock.FeatureFlagsMockWrapper{}
	mockCustomStatesWrapper := &mock.CustomStatesMockWrapper{}
	mock.Flag = wrappers.FeatureFlagResponseModel{
		Name:   "test-flag",
		Status: true,
	}
	cmd := triageUpdateSubCommand(mockResultsPredicatesWrapper, mockFeatureFlagsWrapper, mockCustomStatesWrapper)
	cmd.SetArgs([]string{
		"--similarity-id", "MOCK",
		"--project-id", "MOCK",
		"--severity", "low",
		"--state", "CUSTOM_STATE_1",
		"--scan-type", "sast",
	})

	err := cmd.Execute()
	assert.ErrorContains(t, err, "Failed to get custom state ID for state: CUSTOM_STATE_1: No matching state found for CUSTOM_STATE_1")
}

func TestRunTriageUpdateWithCustomState(t *testing.T) {
	mockResultsPredicatesWrapper := &mock.ResultsPredicatesMockWrapper{}
	mockFeatureFlagsWrapper := &mock.FeatureFlagsMockWrapper{}
	mockCustomStatesWrapper := &mock.CustomStatesMockWrapper{}
	mock.Flag = wrappers.FeatureFlagResponseModel{
		Name:   "test-flag",
		Status: true,
	}
	cmd := triageUpdateSubCommand(mockResultsPredicatesWrapper, mockFeatureFlagsWrapper, mockCustomStatesWrapper)
	cmd.SetArgs([]string{
		"--similarity-id", "MOCK",
		"--project-id", "MOCK",
		"--severity", "low",
		"--state", "demo2",
		"--scan-type", "sast",
	})

	err := cmd.Execute()
	assert.NilError(t, err)
}

func TestRunTriageUpdateWithSystemState(t *testing.T) {
	mockResultsPredicatesWrapper := &mock.ResultsPredicatesMockWrapper{}
	mockFeatureFlagsWrapper := &mock.FeatureFlagsMockWrapper{}
	mockCustomStatesWrapper := &mock.CustomStatesMockWrapper{}

	cmd := triageUpdateSubCommand(mockResultsPredicatesWrapper, mockFeatureFlagsWrapper, mockCustomStatesWrapper)
	cmd.SetArgs([]string{
		"--similarity-id", "MOCK",
		"--project-id", "MOCK",
		"--severity", "low",
		"--state", "TO_VERIFY",
		"--scan-type", "sast",
	})

	err := cmd.Execute()
	assert.NilError(t, err)
}

func TestDetermineSystemOrCustomState(t *testing.T) {
	tests := []struct {
		name                string
		state               string
		customStateID       string
		mockWrapper         wrappers.CustomStatesWrapper
		mockFeatureFlags    wrappers.FeatureFlagsWrapper
		flag                wrappers.FeatureFlagResponseModel
		expectedState       string
		expectedCustomState string
		expectedError       string
	}{
		{
			name:                "Custom state with valid state name",
			state:               "demo2",
			customStateID:       "",
			mockWrapper:         &mock.CustomStatesMockWrapper{},
			mockFeatureFlags:    &mock.FeatureFlagsMockWrapper{},
			flag:                wrappers.FeatureFlagResponseModel{Name: "test-flag", Status: true},
			expectedState:       "",
			expectedCustomState: "2",
			expectedError:       "",
		},
		{
			name:                "Custom state with valid state ID",
			state:               "",
			customStateID:       "2",
			mockWrapper:         &mock.CustomStatesMockWrapper{},
			flag:                wrappers.FeatureFlagResponseModel{Name: "test-flag", Status: true},
			expectedState:       "",
			expectedCustomState: "2",
			expectedError:       "",
		},
		{
			name:                "System state",
			state:               "TO_VERIFY",
			customStateID:       "",
			mockWrapper:         &mock.CustomStatesMockWrapper{},
			flag:                wrappers.FeatureFlagResponseModel{Name: "test-flag", Status: true},
			expectedState:       "TO_VERIFY",
			expectedCustomState: "",
			expectedError:       "",
		},
		{
			name:                "State ID required when state is not provided",
			state:               "",
			customStateID:       "",
			mockWrapper:         &mock.CustomStatesMockWrapper{},
			flag:                wrappers.FeatureFlagResponseModel{Name: "test-flag", Status: true},
			expectedState:       "",
			expectedCustomState: "",
			expectedError:       "state-id is required when state is not provided",
		},
		{
			name:                "Failed to get custom state ID",
			state:               "INVALID_STATE",
			customStateID:       "",
			mockWrapper:         &mock.CustomStatesMockWrapperWithError{},
			flag:                wrappers.FeatureFlagResponseModel{Name: "test-flag", Status: true},
			expectedState:       "",
			expectedCustomState: "",
			expectedError:       "Failed to get custom state ID for state: INVALID_STATE",
		},
		{
			name:                "Both state and state ID provided - valid custom state",
			state:               "demo2",
			customStateID:       "2",
			mockWrapper:         &mock.CustomStatesMockWrapper{},
			flag:                wrappers.FeatureFlagResponseModel{Name: "test-flag", Status: true},
			expectedState:       "",
			expectedCustomState: "2",
			expectedError:       "",
		},
		{
			name:                "Both state and state ID provided - valid system state",
			state:               "TO_VERIFY",
			customStateID:       "2",
			mockWrapper:         &mock.CustomStatesMockWrapper{},
			flag:                wrappers.FeatureFlagResponseModel{Name: "test-flag", Status: true},
			expectedState:       "TO_VERIFY",
			expectedCustomState: "",
			expectedError:       "",
		},
		{
			name:                "Both state and state ID provided - invalid state name",
			state:               "INVALID_STATE",
			customStateID:       "2",
			mockWrapper:         &mock.CustomStatesMockWrapperWithError{},
			flag:                wrappers.FeatureFlagResponseModel{Name: "test-flag", Status: true},
			expectedState:       "",
			expectedCustomState: "2",
			expectedError:       "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mock.Flag = tt.flag
			state, customStateID, err := determineSystemOrCustomState(tt.mockWrapper, tt.mockFeatureFlags, tt.state, tt.customStateID)
			if tt.expectedError != "" {
				assert.ErrorContains(t, err, tt.expectedError)
			} else {
				assert.NilError(t, err)
			}
			assert.Equal(t, state, tt.expectedState)
			assert.Equal(t, customStateID, tt.expectedCustomState)
		})
	}
}
