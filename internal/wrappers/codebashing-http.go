package wrappers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	commonParams "github.com/checkmarx/ast-cli/internal/params"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	failedToParseCodeBashing    = "Failed to parse list results"
	failedGettingCodeBashingURL = "Authentication failed, not able to retrieve codebashing base link"
	tenThousand                 = "10000"
	limit                       = "limit"
	incorrectFlags              = "No codebashing link available"
)

type CodeBashingHTTPWrapper struct {
	path string
}

func NewCodeBashingHTTPWrapper(path string) *CodeBashingHTTPWrapper {
	return &CodeBashingHTTPWrapper{
		path: path,
	}
}

func (r *CodeBashingHTTPWrapper) GetCodeBashingLinks(params map[string]string, codeBashingURL *string) (
	*[]CodeBashingCollection,
	*WebError,
	error,
) {
	clientTimeout := viper.GetUint(commonParams.ClientTimeoutKey)
	params[limit] = tenThousand
	resp, err := SendHTTPRequestWithQueryParams(http.MethodGet, r.path, params, nil, clientTimeout)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	switch resp.StatusCode {
	case http.StatusBadRequest, http.StatusInternalServerError:
		errorModel := WebError{}
		err = decoder.Decode(&errorModel)
		if err != nil {
			return nil, nil, errors.Wrapf(err, failedToParseCodeBashing)
		}
		return nil, &errorModel, nil
	case http.StatusOK:
		var decoded []CodeBashingCollection
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, nil, errors.Wrapf(err, failedToParseCodeBashing)
		}
		err = json.Unmarshal(body, &decoded)
		if err != nil {
			return nil, nil, errors.Wrapf(err, failedToParseCodeBashing)
		}
		if decoded[0].Path == "" {
			return nil, nil, errors.Errorf(incorrectFlags)
		}
		decoded[0].Path = *codeBashingURL + decoded[0].Path
		return &decoded, nil, nil
	default:
		return nil, nil, errors.Errorf("response status code %d", resp.StatusCode)
	}
}

func (r *CodeBashingHTTPWrapper) GetCodeBashingURL(field string) (*string, error) {
	tokenExpirySeconds := viper.GetInt(commonParams.TokenExpirySecondsKey)
	accessToken := getClientCredentialsFromCache(tokenExpirySeconds)
	if accessToken == nil {
		authURI, err := getAuthURI()
		if err != nil {
			return nil, err
		}
		accessKeyID := viper.GetString(commonParams.AccessKeyIDConfigKey)
		accessKeySecret := viper.GetString(commonParams.AccessKeySecretConfigKey)
		astAPIKey := viper.GetString(commonParams.AstAPIKey)
		if accessKeyID == "" && astAPIKey == "" {
			return nil, errors.Errorf(fmt.Sprintf(FailedToAuth, "access key ID"))
		}
		accessToken, err = getClientCredentials(accessKeyID, accessKeySecret, astAPIKey, authURI)
		if err != nil {
			return nil, errors.Errorf(failedGettingCodeBashingURL)
		}
	}
	token, _, err := new(jwt.Parser).ParseUnverified(*accessToken, jwt.MapClaims{})
	if err != nil {
		return nil, errors.Errorf(failedGettingCodeBashingURL)
	}
	var url = ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		url = claims[field].(string)
	} else {
		return nil, errors.Errorf(failedGettingCodeBashingURL)
	}
	return &url, nil
}

func (r *CodeBashingHTTPWrapper) BuildCodeBashingParams(apiParams []CodeBashingParamsCollection) (map[string]string, error) {
	// Marshall entire object to string
	params := make(map[string]string)
	viewJSON, err := json.Marshal(apiParams)
	if err != nil {
		return nil, err
	}
	params["results"] = string(viewJSON)
	return params, nil
}
