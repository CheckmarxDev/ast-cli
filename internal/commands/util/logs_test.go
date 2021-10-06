package util

import (
	"github.com/checkmarxDev/ast-cli/internal/wrappers/mock"
	"gotest.tools/assert"
	"testing"
)

func TestNewLogsCommand(t *testing.T) {
	logsWrapper := &mock.LogsMockWrapper{}
	cmd := NewLogsCommand(logsWrapper)
	assert.Assert(t, cmd != nil, "Logs command must exist")

	//TODO: execute command without flags and try to catch exit code 0
}
