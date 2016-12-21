package circleci

import (
	"fmt"
	"net/http"
)

// EnvVar - CircleCI API ENV variable response
type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// EnvVars calls the /project/:username/:project/envvar endpoint to get
// all ENV variables set for the project builds
func (client *Client) EnvVars(username, project string) ([]*EnvVar, *APIResponse) {
	envVars := []*EnvVar{}
	path := fmt.Sprintf("project/%s/%s/envvar", username, project)
	apiResp := client.request(http.MethodGet, path, nil, nil, &envVars)
	return envVars, apiResp
}

// AddEnvVar calls the /project/:username/:project/envvar endpoint
// to set an ENV variable for the project builds
func (client *Client) AddEnvVar(username, project, name, value string) (*EnvVar, *APIResponse) {
	addedVar := &EnvVar{}
	addVar := &EnvVar{Name: name, Value: value}
	path := fmt.Sprintf("project/%s/%s/envvar", username, project)
	apiResp := client.request(http.MethodPost, path, nil, addVar, addedVar)
	return addedVar, apiResp
}

// GetEnvVar calls the /project/:username/:project/envvar/:name endpoint
// to get an ENV variable for the project builds
func (client *Client) GetEnvVar(username, project, name string) (*EnvVar, *APIResponse) {
	envVar := &EnvVar{}
	path := fmt.Sprintf("project/%s/%s/envvar/%s", username, project, name)
	apiResp := client.request(http.MethodGet, path, nil, nil, envVar)
	return envVar, apiResp
}

// DeleteEnvVar calls the /project/:username/:project/envvar/:name endpoint
// to delete an ENV variable for the project builds. Returns an APIMessageResponse
// instead of an EnvVar
func (client *Client) DeleteEnvVar(username, project, name string) (*APIMessageResponse, *APIResponse) {
	resp := &APIMessageResponse{}
	path := fmt.Sprintf("project/%s/%s/envvar/%s", username, project, name)
	apiResp := client.request(http.MethodDelete, path, nil, nil, resp)
	return resp, apiResp
}
