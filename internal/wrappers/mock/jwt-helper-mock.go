package mock

import (
	"github.com/checkmarx/ast-cli/internal/wrappers"
	"strings"
)

type JWTMockWrapper struct{}

// GetAllowedEngines mock for tests
func (*JWTMockWrapper) GetAllowedEngines(featureFlagsWrapper wrappers.FeatureFlagsWrapper) (allowedEngines map[string]bool, err error) {
	allowedEngines = make(map[string]bool)
	engines := []string{"sast", "iac-security", "sca", "api-security"}
	for _, value := range engines {
		allowedEngines[strings.ToLower(value)] = true
	}
	return allowedEngines, nil
}
