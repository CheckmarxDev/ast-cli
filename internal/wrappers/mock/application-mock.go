package mock

import (
	applicationErrors "github.com/checkmarx/ast-cli/internal/errors"
	"github.com/checkmarx/ast-cli/internal/wrappers"
	"github.com/pkg/errors"
	"time"
)

type ApplicationsMockWrapper struct{}

func (a ApplicationsMockWrapper) Get(params map[string]string) (*wrappers.ApplicationsResponseModel, *wrappers.ErrorModel, error) {
	if params["name"] == NoPermissionApp {
		return nil, nil, errors.Errorf(applicationErrors.ApplicationNoPermission)
	}
	if params["name"] == ApplicationDoesntExist {
		return nil, nil, errors.Errorf(applicationErrors.ApplicationDoesntExist)
	}
	mockApplication := wrappers.Application{
		ID:          "mockID",
		Name:        "MOCK",
		Description: "This is a mock application",
		Criticality: 2,
		ProjectIds:  []string{"ProjectID1", "ProjectID2"},
		CreatedAt:   time.Now(),
	}

	response := &wrappers.ApplicationsResponseModel{
		TotalCount:   1,
		Applications: []wrappers.Application{mockApplication},
	}

	return response, nil, nil
}
