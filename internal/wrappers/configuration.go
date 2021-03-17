package wrappers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/checkmarxDev/ast-cli/internal/params"
	commonParams "github.com/checkmarxDev/ast-cli/internal/params"
	"github.com/spf13/viper"
)

const defaultProfileName = "default"

func PromptConfiguration() {
	reader := bufio.NewReader(os.Stdin)
	baseURI := viper.GetString(commonParams.BaseURIKey)
	accessKeySecret := viper.GetString(commonParams.AccessKeySecretConfigKey)
	accessKey := viper.GetString(commonParams.AccessKeyIDConfigKey)
	fmt.Print(fmt.Sprintf("AST Base URI [%s]: ", baseURI))
	baseURI, _ = reader.ReadString('\n')
	if len(baseURI) > 1 {
		baseURI = strings.Replace(baseURI, "\n", "", -1)
		baseURI = strings.Replace(baseURI, "\r\n", "", -1)
		setConfigPropertyQuiet(commonParams.BaseURIKey, baseURI)
	}
	fmt.Print(fmt.Sprintf("AST Access Key [%s]: ", obfuscateString(accessKey)))
	accessKey, _ = reader.ReadString('\n')
	if len(accessKey) > 1 {
		accessKey = strings.Replace(accessKey, "\n", "", -1)
		accessKey = strings.Replace(accessKey, "\r\n", "", -1)
		setConfigPropertyQuiet(commonParams.AccessKeyIDConfigKey, accessKey)
	}
	fmt.Print(fmt.Sprintf("AST Key Secret [%s]: ", obfuscateString(accessKeySecret)))
	accessKeySecret, _ = reader.ReadString('\n')
	if len(accessKeySecret) > 1 {
		accessKeySecret = strings.Replace(accessKeySecret, "\n", "", -1)
		accessKeySecret = strings.Replace(accessKeySecret, "\r\n", "", -1)
		setConfigPropertyQuiet(commonParams.AccessKeySecretConfigKey, accessKeySecret)
	}
}

func obfuscateString(str string) string {
	if len(str) > 4 {
		return "******" + str[len(str)-4:]
	} else if len(str) > 1 {
		return "******"
	} else {
		return ""
	}
}

func setConfigPropertyQuiet(propName, propValue string) {
	viper.Set(propName, propValue)
	// You should be able to  call WriteConfig() but it will fail if the
	// config file doesn't already exist, this is a known viper bug.
	// SafeWriteConfig() will not update files but it will create them, combined
	// this code will successfully update files.
	if viperErr := viper.SafeWriteConfig(); viperErr != nil {
		_ = viper.WriteConfig()
	}
}

func SetConfigProperty(propName, propValue string) {
	fmt.Println("Setting property [", propName, "] to value [", propValue, "]")
	setConfigPropertyQuiet(propName, propValue)
}

func LoadConfiguration() {
	profile := findProfile()
	viper.AddConfigPath("$HOME/.checkmarx")
	configFile := ".checkmarxcli"
	if profile != defaultProfileName {
		configFile += "_"
		configFile += profile
	}
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	_ = viper.ReadInConfig()
}

func findProfile() string {
	profileName := defaultProfileName
	for idx, b := range os.Args {
		if b == "--profile" {
			profileIdx := idx + 1
			if len(os.Args) > profileIdx {
				profileName = os.Args[profileIdx]
				fmt.Println("Using custom profile: ", profileName)
			}
		}
	}
	return profileName
}

func ShowConfiguration() {
	fmt.Println("Current Effective Configuration")
	baseURI := viper.GetString(params.BaseURIKey)
	fmt.Println("\tBaseURI: ", baseURI)
}
