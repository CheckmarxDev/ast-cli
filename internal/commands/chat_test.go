package commands

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/google/uuid"
	"gotest.tools/assert"
)

func TestChatHelp(t *testing.T) {
	execCmdNilAssertion(t, "help", "chat")
}

func TestChatInvalidId(t *testing.T) {
	err, buffer := executeRedirectedTestCommand("chat",
		"--conversation-id", "invalidId",
		"--chat-apikey", "apiKey",
		"--user-input", "userInput",
		"--result-file", "file",
		"--result-line", "0",
		"--result-severity", "LOW",
		"--result-vulnerability", "Vulnerability")
	assert.NilError(t, err)
	output, err := io.ReadAll(buffer)
	assert.NilError(t, err)
	s := string(output)
	assert.Assert(t, s == fmt.Sprintf(ConversationIdErrorFormat, "invalidId"))
}

func TestChatInvalidFile(t *testing.T) {
	err, buffer := executeRedirectedTestCommand("chat",
		"--conversation-id", uuid.New().String(),
		"--chat-apikey", "apiKey",
		"--user-input", "userInput",
		"--result-file", "invalidfile",
		"--result-line", "0",
		"--result-severity", "LOW",
		"--result-vulnerability", "Vulnerability")
	assert.NilError(t, err)
	output, err := io.ReadAll(buffer)
	assert.NilError(t, err)
	s := strings.ToLower(string(output))
	assert.Assert(t, s == fmt.Sprintf(FileErrorFormat, "invalidfile"))
}

func TestChatInvalidApiKey(t *testing.T) {
	err, buffer := executeRedirectedTestCommand("chat",
		"--conversation-id", uuid.New().String(),
		"--chat-apikey", "apiKey",
		"--user-input", "userInput",
		"--result-file", "./data/Dockerfile",
		"--result-line", "0",
		"--result-severity", "LOW",
		"--result-vulnerability", "Vulnerability")
	assert.NilError(t, err)
	output, err := io.ReadAll(buffer)
	assert.NilError(t, err)
	s := strings.ToLower(string(output))
	assert.Assert(t, strings.Contains(s, "api_key"), s)
}
