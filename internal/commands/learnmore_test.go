//go:build !integration

package commands

import (
	"github.com/checkmarx/ast-cli/internal/commands/util"
	"testing"

	"gotest.tools/assert"
)

func TestLearnMoreHelp(t *testing.T) {
	execCmdNilAssertion(t, "utils", "learn-more", "--help")
}

func TestLearnMoreMockQueryIdSummaryConsole(t *testing.T) {
	execCmdNilAssertion(t, "utils", "learn-more", "--query-id", "MOCK")
}

func TestLearnMoreMockQueryIdJsonFormat(t *testing.T) {
	execCmdNilAssertion(t, "utils", "learn-more", "--query-id", "MOCK", "--format", "json")
}

func TestLearnMoreMockQueryIdListFormat(t *testing.T) {
	execCmdNilAssertion(t, "utils", "learn-more", "--query-id", "MOCK", "--format", "list")
}

func TestLearnMoreMockQueryIdTableFormat(t *testing.T) {
	execCmdNilAssertion(t, "utils", "learn-more", "--query-id", "MOCK", "--format", "table")
}

func TestLearnMoreMockMissingQueryId(t *testing.T) {
	err := execCmdNotNilAssertion(t, "utils", "learn-more")
	assert.Assert(t, err.Error() == "required flag(s) \"query-id\" not set")
}

func TestLearnMoreMockQueryIdInvalidFormat(t *testing.T) {
	err := execCmdNotNilAssertion(t, "utils", "learn-more", "--query-id", "MOCK", "--format", "MOCK")
	assert.Assert(t, err.Error() == "Invalid format MOCK")
}

func TestCommand(t *testing.T) {
	cmd := util.NewLearnMoreCommand(nil)
	assert.Assert(t, cmd != nil, "Learnmore command must exist")
}
