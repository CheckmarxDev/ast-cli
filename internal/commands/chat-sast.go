package commands

import (
	"fmt"

	"github.com/checkmarx/ast-cli/internal/commands/chatsast"
	"github.com/checkmarx/ast-cli/internal/commands/util/printer"
	"github.com/checkmarx/ast-cli/internal/logger"
	"github.com/checkmarx/ast-cli/internal/params"
	"github.com/checkmarx/ast-cli/internal/wrappers"
	"github.com/checkmarxDev/gpt-wrapper/pkg/connector"
	"github.com/checkmarxDev/gpt-wrapper/pkg/message"
	"github.com/checkmarxDev/gpt-wrapper/pkg/role"
	"github.com/checkmarxDev/gpt-wrapper/pkg/wrapper"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const ScanResultsFileErrorFormat = "Error reading and parsing scan results %s"
const CreatePromptErrorFormat = "Error creating prompt for result ID %s"
const UserInputRequiredErrorFormat = "%s is required when %s is provided"

func ChatSastSubCommand(chatWrapper wrappers.ChatWrapper) *cobra.Command {
	chatSastCmd := &cobra.Command{
		Use:   "sast",
		Short: "OpenAI-based SAST results remediation",
		Long:  "Use OpenAI models to remediate SAST results and chat about them",
		RunE:  runChatSast(chatWrapper),
	}

	chatSastCmd.Flags().String(params.ChatAPIKey, "", "OpenAI API key")
	chatSastCmd.Flags().String(params.ChatConversationID, "", "ID of existing conversation")
	chatSastCmd.Flags().String(params.ChatUserInput, "", "User question")
	chatSastCmd.Flags().String(params.ChatModel, "", "OpenAI model version")
	chatSastCmd.Flags().String(params.ChatSastScanResultsFile, "", "Results file in JSON format containing SAST scan results")
	chatSastCmd.Flags().String(params.ChatSastSourceDir, "", "Source code root directory relevant for the results file")
	chatSastCmd.Flags().String(params.ChatSastLanguage, "", "Language of the result to remediate")
	chatSastCmd.Flags().String(params.ChatSastQuery, "", "Query of the result to remediate")
	chatSastCmd.Flags().String(params.ChatSastResultId, "", "ID of the result to remediate")

	_ = chatSastCmd.MarkFlagRequired(params.ChatAPIKey)
	_ = chatSastCmd.MarkFlagRequired(params.ChatSastScanResultsFile)
	_ = chatSastCmd.MarkFlagRequired(params.ChatSastSourceDir)
	_ = chatSastCmd.MarkFlagRequired(params.ChatSastResultId)

	return chatSastCmd
}

func runChatSast(sastChatWrapper wrappers.ChatSastWrapper) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		chatAPIKey, _ := cmd.Flags().GetString(params.ChatAPIKey)
		chatConversationID, _ := cmd.Flags().GetString(params.ChatConversationID)
		chatModel, _ := cmd.Flags().GetString(params.ChatModel)
		scanResultsFile, _ := cmd.Flags().GetString(params.ChatSastScanResultsFile)
		sourceDir, _ := cmd.Flags().GetString(params.ChatSastSourceDir)
		//sastLanguage, _ := cmd.Flags().GetString(params.SastLanguage) // TODO: add support for language
		//sastQuery, _ := cmd.Flags().GetString(params.SastQuery) // TODO: add support for query
		sastResultId, _ := cmd.Flags().GetString(params.ChatSastResultId)

		statefulWrapper := wrapper.NewStatefulWrapper(connector.NewFileSystemConnector(""), chatAPIKey, chatModel, dropLen, 0)

		newConversation := false
		var userInput string
		if chatConversationID == "" {
			newConversation = true
			chatConversationID = statefulWrapper.GenerateId().String()
		} else {
			userInput, _ = cmd.Flags().GetString(params.ChatUserInput)
			if userInput == "" {
				msg := fmt.Sprintf(UserInputRequiredErrorFormat, params.ChatUserInput, params.ChatConversationID)
				logger.PrintIfVerbose(msg)
				return outputError(cmd, uuid.Nil, errors.Errorf(msg))
			}
		}

		id, err := uuid.Parse(chatConversationID)
		if err != nil {
			logger.PrintIfVerbose(err.Error())
			return outputError(cmd, id, errors.Errorf(ConversationIDErrorFormat, chatConversationID))
		}

		var newMessages []message.Message
		if newConversation {
			prompt, err := buildPrompt(scanResultsFile, sastResultId, sourceDir)
			if err != nil {
				logger.PrintIfVerbose(err.Error())
				return outputError(cmd, id, err)
			}
			newMessages = append(newMessages, message.Message{
				Role:    role.User,
				Content: prompt,
			})
		} else {
			newMessages = append(newMessages, message.Message{
				Role:    role.User,
				Content: userInput,
			})
		}

		response, err := sastChatWrapper.Call(statefulWrapper, id, newMessages)
		if err != nil {
			return outputError(cmd, id, err)
		}

		responseContent := getMessageContents(response)

		return printer.Print(cmd.OutOrStdout(), &OutputModel{
			ConversationID: id.String(),
			Response:       responseContent,
		}, printer.FormatJSON)
	}
}

func buildPrompt(scanResultsFile, sastResultId, sourceDir string) (string, error) {
	scanResults, err := chatsast.ReadResultsSAST(scanResultsFile)
	if err != nil {
		return "", fmt.Errorf("error in build-prompt: %s: %w", fmt.Sprintf(ScanResultsFileErrorFormat, scanResultsFile), err)
	}

	if sastResultId == "" {
		return "", errors.Errorf(fmt.Sprintf("error in build-prompt: currently only --%s is supported", params.ChatSastResultId))
	}

	//languages := GetLanguages(scanResults, sastLanguage)
	//queriesByLanguage := GetQueries(scanResults, languages, sastQuery)
	sastResult, err := chatsast.GetResultById(scanResults, sastResultId)
	if err != nil {
		return "", fmt.Errorf("error in build-prompt: %w", err)
	}

	sources, err := chatsast.GetSourcesForResult(sastResult, sourceDir)
	if err != nil {
		return "", fmt.Errorf("error in build-prompt: %w", err)
	}

	prompt, err := chatsast.CreatePrompt(sastResult, sources)
	if err != nil {
		return "", fmt.Errorf("error in build-prompt: %s: %w", fmt.Sprintf(CreatePromptErrorFormat, sastResultId), err)
	}

	return prompt, nil
}

func getMessageContents(response []message.Message) []string {
	var responseContent []string
	for _, r := range response {
		responseContent = append(responseContent, r.Content)
	}
	return responseContent
}
