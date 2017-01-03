package circleci

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Build - CircleCI build response
type Build struct {
	VcsURL          string       `json:"vcs_url"`
	BuildURL        string       `json:"build_url"`
	BuildNum        int          `json:"build_num"`
	Branch          string       `json:"branch"`
	VcsRevison      string       `json:"vcs_revision"`
	CommitterName   string       `json:"committer_name"`
	CommitterEmail  string       `json:"committer_email"`
	Subject         string       `json:"subject"`
	Body            string       `json:"body"`
	Why             string       `json:"why"`
	DontBuild       string       `json:"dont_build"`
	QueuedAt        string       `json:"queued_at"`
	StartTime       string       `json:"start_time"`
	StopTime        string       `json:"stop_time"`
	BuildTimeMillis int          `json:"build_time_millis"`
	Username        string       `json:"username"`
	Reponame        string       `json:"reponame"`
	Lifecycle       string       `json:"lifecycle"`
	Outcome         string       `json:"outcome"`
	Status          string       `json:"status"`
	Steps           []*Step      `json:"steps"`
	RetryOf         int          `json:"retry_of"`
	PreviousBuild   *BuildStatus `json:"previous_build"`
}

// BuildStatus - CircleCI status for a build response
type BuildStatus struct {
	BuildNum int    `json:"build_num"`
	Status   string `json:"status"`
}

// Step is a step in the build
type Step struct {
	Name    string    `json:"name"`
	Actions []*Action `json:"actions"`
}

// Action is a single step action for a build
type Action struct {
	BashCommand        string     `json:"bash_command"`
	RunTimeMillis      int        `json:"run_time_millis"`
	Continue           string     `json:"continue"`
	Parallel           bool       `json:"parallel"`
	StartTime          *time.Time `json:"start_time"`
	Name               string     `json:"name"`
	Messages           []string   `json:"messages"`
	Step               int        `json:"step"`
	ExitCode           int        `json:"exit_code"`
	EndTime            *time.Time `json:"end_time"`
	Index              int        `json:"index"`
	Status             string     `json:"status"`
	Timedout           bool       `json:"timedout"`
	InfrastructureFail bool       `json:"infrastructure_fail"`
	Type               string     `json:"type"`
	Source             string     `json:"source"`
	Failed             bool       `json:"failed"`
}

// Artifact is for a build that has exposed artifacts from the build process
type Artifact struct {
	NodeIndex  int    `json:"node_index"`
	Path       string `json:"path"`
	PrettyPath string `json:"pretty_path"`
	URL        string `json:"url"`
}

// BuildTests - CircleCI API tests metadata response
type BuildTests struct {
	Tests []*BuildTest `json:"tests"`
}

// BuildTest - CircleCI API tests metadata object
type BuildTest struct {
	Message   string  `json:"message"`
	File      string  `json:"file"`
	Source    string  `json:"source"`
	RunTime   float64 `json:"run_time"`
	Result    string  `json:"result"`
	Name      string  `json:"name"`
	Classname string  `json:"classname"`
}

// BuildPostBody - The optional options to pass with a POST body when
// triggering a build
type BuildPostBody struct {
	Revision        string            `json:"revision"`
	Tag             string            `json:"tag"`
	Parallel        int               `json:"parallel"`
	BuildParameters map[string]string `json:"build_parameters"`
}

// NewBuild calls /project/:username/:reponame endpaoint to trigger a new
// build for the project
func (client *Client) NewBuild(username, project string, body *BuildPostBody) (*Build, *APIResponse) {
	build := &Build{}
	path := fmt.Sprintf("project/%s/%s", username, project)
	apiResp := client.request(http.MethodPost, path, nil, body, build)
	return build, apiResp
}

// BuildBranch calls /project/:username/:reponame/tree/:branch endpaoint to trigger a new
// build for the project
func (client *Client) BuildBranch(username, project, branch string, body *BuildPostBody) (*Build, *APIResponse) {
	build := &Build{}
	path := fmt.Sprintf("project/%s/%s/tree/%s", username, project, url.QueryEscape(branch))
	apiResp := client.request(http.MethodPost, path, nil, body, build)
	return build, apiResp
}

// GetBuild calls the /project/:username/:reponame/:build_num endpoint to return a build
func (client *Client) GetBuild(username, project string, buildNum int) (*Build, *APIResponse) {
	build := &Build{}
	path := fmt.Sprintf("project/%s/%s/%d", username, project, buildNum)
	apiResp := client.request(http.MethodGet, path, nil, nil, build)
	return build, apiResp
}

// RetryBuild calls the /project/:username/:reponame/:build_num/retry endpoint to retry a build
func (client *Client) RetryBuild(username, project string, buildNum int) (*Build, *APIResponse) {
	build := &Build{}
	path := fmt.Sprintf("project/%s/%s/%d/retry", username, project, buildNum)
	apiResp := client.request(http.MethodGet, path, nil, nil, build)
	return build, apiResp
}

// CancelBuild calls the /project/:username/:reponame/:build_num/cancel endpoint to cancel a build
func (client *Client) CancelBuild(username, project string, buildNum int) (*Build, *APIResponse) {
	build := &Build{}
	path := fmt.Sprintf("project/%s/%s/%d/cancel", username, project, buildNum)
	apiResp := client.request(http.MethodPost, path, nil, nil, build)
	return build, apiResp
}

// BuildArtifacts calls the /project/:username/:reponame/:build_num/artifacts
// endpoint to get all artifacts for a build
func (client *Client) BuildArtifacts(username, project string, buildNum int) ([]*Artifact, *APIResponse) {
	artifacts := []*Artifact{}
	path := fmt.Sprintf("project/%s/%s/%d/artifacts", username, project, buildNum)
	apiResp := client.request(http.MethodGet, path, nil, nil, &artifacts)
	return artifacts, apiResp
}

// BuildTests calls the /project/:username/:reponame/:build_num/tests
// endpoint to get all tests for a build
func (client *Client) BuildTests(username, project string, buildNum int) (*BuildTests, *APIResponse) {
	tests := &BuildTests{}
	path := fmt.Sprintf("project/%s/%s/%d/tests", username, project, buildNum)
	apiResp := client.request(http.MethodGet, path, nil, nil, tests)
	return tests, apiResp
}
