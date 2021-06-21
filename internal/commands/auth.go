package commands

import (
	"fmt"

	"github.com/checkmarxDev/ast-cli/internal/params"
	"github.com/checkmarxDev/ast-cli/internal/wrappers"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	failedAuthValidate   = "Failed authentication!"
	failedCreatingClient = "failed creating client"
	pleaseProvideFlag    = "%s: Please provide %s flag"
	adminClientID        = "ast-app"
	adminClientSecret    = "1d71c35c-818e-4ee8-8fb1-d6cbf8fe2e2a"
)

type ClientCreated struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
}

func NewAuthCommand(authWrapper wrappers.AuthWrapper) *cobra.Command {
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "Manage authentication",
	}
	createClientCmd := &cobra.Command{
		Use:   "register",
		Short: "Register new OAuth2 client for ast",
		Long: "Register new OAuth2 client and outputs its generated credentials in the format <key>=<value>.\n" +
			"\n" +
			"  cx auth register -u <user> -p <pass> \n",
		RunE: runRegister(authWrapper),
	}
	createClientCmd.PersistentFlags().StringP(clientDescriptionFlag, clientDescriptionSh, "", "A client description")
	createClientCmd.PersistentFlags().StringP(usernameFlag, usernameSh, "", "Username for Ast user that privileges to "+
		"create clients")
	createClientCmd.PersistentFlags().StringP(passwordFlag, passwordSh, "", "Password for Ast user that privileges to "+
		"create clients")
	createClientCmd.PersistentFlags().StringSliceP(clientRolesFlag, clientRolesSh, []string{"ast-admin"},
		"A list of roles of the client")
	validLoginCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validates a client/secret",
		Long:  "Validates if a client/secret pair can communicate with AST.",
		RunE:  validLogin(),
	}
	authCmd.AddCommand(createClientCmd, validLoginCmd)
	return authCmd
}

func validLogin() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		scansWrapper := wrappers.NewHTTPScansWrapper(viper.GetString(params.ScansPathKey))
		paramsList := make(map[string]string)
		_, _, err := scansWrapper.Get(paramsList)
		if err != nil {
			return errors.Errorf("%s", failedAuthValidate)
		}

		fmt.Println("Successfully authenticated to AST server!")

		return nil
	}
}

func runRegister(authWrapper wrappers.AuthWrapper) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		username, _ := cmd.Flags().GetString(usernameFlag)
		if username == "" {
			return errors.Errorf(pleaseProvideFlag, failedCreatingClient, usernameFlag)
		}

		password, _ := cmd.Flags().GetString(passwordFlag)
		if password == "" {
			return errors.Errorf(pleaseProvideFlag, failedCreatingClient, passwordFlag)
		}

		roles, _ := cmd.Flags().GetStringSlice(clientRolesFlag)
		description, _ := cmd.Flags().GetString(clientDescriptionFlag)
		generatedClientID := "ast-plugins-" + uuid.New().String()
		generatedClientSecret := uuid.New().String()
		client := &wrappers.Oath2Client{
			Name:        generatedClientID,
			Roles:       roles,
			Description: description,
			Secret:      generatedClientSecret,
		}

		errorMsg, err := authWrapper.CreateOauth2Client(client, username, password, adminClientID, adminClientSecret)
		if err != nil {
			fmt.Println("Could not create OAuth2 credentials!")
			return nil
		}

		if errorMsg != nil {
			return errors.Errorf("%s: CODE: %d, %s", failedCreatingClient, errorMsg.Code, errorMsg.Message)
		}

		AuthGeneratedClientID = generatedClientID
		AuthGeneratedClientSecret = generatedClientSecret
		fmt.Printf("%s=%s\n", params.AccessKeyIDEnv, generatedClientID)
		fmt.Printf("%s=%s\n", params.AccessKeySecretEnv, generatedClientSecret)
		return nil
	}
}

var AuthGeneratedClientID = ""
var AuthGeneratedClientSecret = ""
