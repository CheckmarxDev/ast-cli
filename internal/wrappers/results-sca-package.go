package wrappers

type ScaPackageCollection struct {
	Id                  string             `json:"id,omitempty"`
	Locations           []*string          `json:"locations,omitempty"`
	DependencyPathArray [][]DependencyPath `json:"dependencyPaths,omitempty"`
	Outdated            bool               `json:"outdated,omitempty"`
}

type DependencyPath struct {
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Version       string `json:"version,omitempty"`
	IsResolved    bool   `json:"isResolved,omitempty"`
	IsDevelopment bool   `json:"isDevelopment,omitempty"`
}
