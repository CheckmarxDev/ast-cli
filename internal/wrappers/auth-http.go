package wrappers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/checkmarxDev/ast-cli/internal/params"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type AuthHTTPWrapper struct {
	path string
}

func NewAuthHTTPWrapper() AuthWrapper {
	return &AuthHTTPWrapper{}
}

const failedToParseCreateClientResult = "failed to parse create client result"

func (a *AuthHTTPWrapper) SetPath(newPath string) {
	a.path = newPath
}

func (a *AuthHTTPWrapper) CreateOauth2Client(client *Oath2Client, username, password,
	adminClientID, adminClientSecret string) (*ErrorMsg, error) {
	jsonBytes, err := json.Marshal(client)
	if err != nil {
		return nil, err
	}
	// Update the auth path, delayed to here because bind not ready in main.go
	createClientPath := viper.GetString(params.CreateOath2ClientPathKey)
	tenant := viper.GetString(params.TenantKey)
	createClientPath = strings.Replace(createClientPath, "organization", tenant, 1)
	a.SetPath(createClientPath)
	// send the request
	res, err := SendHTTPRequestPasswordAuth(http.MethodPost, a.path, bytes.NewBuffer(jsonBytes), DefaultTimeoutSeconds,
		username, password, adminClientID, adminClientSecret)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusBadRequest:
		decoder := json.NewDecoder(res.Body)
		errorMsg := ErrorMsg{}
		err = decoder.Decode(&errorMsg)
		if err != nil {
			return nil, errors.Wrap(err, failedToParseCreateClientResult)
		}
		return &errorMsg, nil
	case http.StatusOK:
		return nil, nil
	default:
		b, err := ioutil.ReadAll(res.Body)
		return nil, errors.Errorf("response status code %d body %s", res.StatusCode, func() string {
			if err != nil {
				return ""
			}
			return string(b)
		}())
	}
}
