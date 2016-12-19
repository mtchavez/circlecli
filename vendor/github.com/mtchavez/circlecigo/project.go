package circleci

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	// BuildLimitMax is the max builds to return from the API
	BuildLimitMax = 100
)

var (
	// ValidBuildFilters is the list of valid build statuses you can filter by
	ValidBuildFilters = map[string]string{
		"completed":  "completed",
		"successful": "successful",
		"failed":     "failed",
		"running":    "running",
	}
)

// Project - CircleCI API response for a project
type Project struct {
	IRCServer            string                    `json:"irc_server"`
	Scopes               []string                  `json:"scopes"`
	IRCKeyword           string                    `json:"irc_keyword"`
	Followed             bool                      `json:"followed"`
	VCSType              string                    `json:"vcs-type"`
	AWS                  *ProjectAWS               `json:"aws"`
	SlackWebhookURL      string                    `json:"slack_webhook_url"`
	FlowdockAPIToken     string                    `json:"flowdock_api_token"`
	Parallel             int                       `json:"parallel"`
	Username             string                    `json:"username"`
	CampfireRoom         string                    `json:"campfire_room"`
	Extra                string                    `json:"extra"`
	Branches             map[string]*ProjectBranch `json:"branches"`
	Jira                 string                    `json:"jira"`
	SlackSubdomain       string                    `json:"slack_subdomain"`
	Following            bool                      `json:"following"`
	Setup                string                    `json:"setup"`
	CampfireSubdomain    string                    `json:"campfire_subdomain"`
	SlackNotifyPrefs     string                    `json:"slack_notify_prefs"`
	IRCPassword          string                    `json:"irc_password"`
	VcsURL               string                    `json:"vcs_url"`
	DefaultBranch        string                    `json:"default_branch"`
	HipchatAPIToken      string                    `json:"hipchat_api_token"`
	IRCUsername          string                    `json:"irc_username"`
	Language             string                    `json:"language"`
	SlackChannelOverride string                    `json:"slack_channel_override"`
	HipchatNotify        string                    `json:"hipchat_notify"`
	SlackAPIToken        string                    `json:"slack_api_token"`
	HasUsableKey         bool                      `json:"has_usable_key"`
	IRCNotifyPrefs       string                    `json:"irc_notify_prefs"`
	CampfireToken        string                    `json:"campfire_token"`
	SlackChannel         string                    `json:"slack_channel"`
	FeatureFlags         map[string]bool           `json:"feature_flags"`
	CampfireNotifyPrefs  string                    `json:"campfire_notify_prefs"`
	HipchatRoom          string                    `json:"hipchat_room"`
	PostDependencies     string                    `json:"post_dependencies"`
	HerokuDeployUser     string                    `json:"heroku_deploy_user"`
	IRCChannel           string                    `json:"irc_channel"`
	Oss                  bool                      `json:"oss"`
	Reponame             string                    `json:"reponame"`
	HipchatNotifyPrefs   string                    `json:"hipchat_notify_prefs"`
	Compile              string                    `json:"compile"`
	Dependencies         string                    `json:"dependencies"`
	Test                 string                    `json:"test"`
	SSHKeys              []*ProjectSSHKey          `json:"ssh_keys"`
}

// ProjectAWS - CircleCI project aws settings
type ProjectAWS struct {
	Keypair *ProjectAWSKeypair `json:"keypair"`
}

// ProjectAWSKeypair - CircleCI project aws keypair
type ProjectAWSKeypair struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key_id"`
}

// ProjectBranch - CircleCI API project branch
type ProjectBranch struct {
	LastNonSuccess *ProjectBuildDetails   `json:"last_non_success"`
	LastSuccess    *ProjectBuildDetails   `json:"last_success"`
	PusherLogins   []string               `json:"pusher_logins"`
	RecentBuilds   []*ProjectBuildDetails `json:"recent_builds"`
	RunningBuilds  []*ProjectBuildDetails `json:"running_builds"`
}

// ProjectBuildDetails - CircleCI API details for a build from a project view
type ProjectBuildDetails struct {
	Outcome     string `json:"outcome"`
	Status      string `json:"status"`
	BuildNum    int    `json:"build_num"`
	VCSRevision string `json:"vcs_revision"`
	PushedAt    string `json:"pushed_at"`
	AddedAt     string `json:"added_at"`
}

// ProjectSSHKey - CircleCI API project ssh key
type ProjectSSHKey struct {
	Hostname    string `json:"hostname"`
	PublicKey   string `json:"public_key"`
	Fingerprint string `json:"fingerprint"`
}

// ProjectFollow - CircleCI API project follow response
type ProjectFollow struct {
	Followed   bool `json:"followed"`
	FirstBuild int  `json:"first_build"`
}

// ProjectClearCache - CircleCI API response when clearing a project cache
type ProjectClearCache struct {
	Status string `json:"status"`
}

// Projects calls /projects to get all the projects you follow
func (client *Client) Projects() ([]*Project, *APIResponse) {
	projects := []*Project{}
	apiResp := client.request(http.MethodGet, "projects", nil, nil, &projects)
	return projects, apiResp
}

// ProjectFollow calls the /project/:username/:project/follow endpoint to follow a project
func (client *Client) ProjectFollow(username, project string) (*ProjectFollow, *APIResponse) {
	follow := &ProjectFollow{}
	path := fmt.Sprintf("project/%s/%s/follow", username, project)
	apiResp := client.request(http.MethodGet, path, nil, nil, follow)
	return follow, apiResp
}

// ProjectUnfollow calls the /project/:username/:project/unfollow endpoint to unfollow a project
func (client *Client) ProjectUnfollow(username, project string) (*ProjectFollow, *APIResponse) {
	follow := &ProjectFollow{}
	path := fmt.Sprintf("project/%s/%s/unfollow", username, project)
	apiResp := client.request(http.MethodGet, path, nil, nil, follow)
	return follow, apiResp
}

// ProjectClearCache calls the /project/:username/:project/unfollow endpoint to unfollow a project
func (client *Client) ProjectClearCache(username, project string) (*ProjectClearCache, *APIResponse) {
	clearCache := &ProjectClearCache{}
	path := fmt.Sprintf("project/%s/%s/build-cache", username, project)
	apiResp := client.request(http.MethodDelete, path, nil, nil, clearCache)
	return clearCache, apiResp
}

// ProjectRecentBuilds /project/:username/:project endpoint to get all the
// recent builds for a project.
// Takes url.Values to set limit of returned builds (no more than BuildLimitMax)
// and filter to filter builds by status
// and an offset of number of builds to page through
func (client *Client) ProjectRecentBuilds(username, project string, params url.Values) ([]*Build, *APIResponse) {
	builds := []*Build{}
	if params == nil {
		params = url.Values{}
	}
	client.verifyBuildsParams(&params)
	path := fmt.Sprintf("project/%s/%s", username, project)
	apiResp := client.request(http.MethodGet, path, params, nil, &builds)
	return builds, apiResp

}

// ProjectRecentBuildsBranch /project/:username/:project/tree/:branch endpoint to get all the
// recent builds for a branch of a project.
// Takes url.Values to set limit of returned builds (no more than BuildLimitMax)
// and filter to filter builds by status
// and an offset of number of builds to page through
func (client *Client) ProjectRecentBuildsBranch(username, project, branch string, params url.Values) ([]*Build, *APIResponse) {
	builds := []*Build{}
	if params == nil {
		params = url.Values{}
	}
	client.verifyBuildsParams(&params)
	path := fmt.Sprintf("project/%s/%s/tree/%s", username, project, url.QueryEscape(branch))
	apiResp := client.request(http.MethodGet, path, params, nil, &builds)
	return builds, apiResp

}

// ProjectEnable calls /project/:username/:project/enable to enable a project
// which will add an SSH key to VCS and will require privileges to do so
func (client *Client) ProjectEnable(username, project string) (*Project, *APIResponse) {
	enabledProject := &Project{}
	path := fmt.Sprintf("project/%s/%s/enable", username, project)
	apiResp := client.request(http.MethodPost, path, nil, nil, enabledProject)
	return enabledProject, apiResp
}

// ProjectSettings calls /project/:username/:project/settings to get the
// settings for a project
func (client *Client) ProjectSettings(username, project string) (*Project, *APIResponse) {
	foundProject := &Project{}
	path := fmt.Sprintf("project/%s/%s/settings", username, project)
	apiResp := client.request(http.MethodGet, path, nil, nil, foundProject)
	return foundProject, apiResp
}

// verifyBuildsParams ensures limit param is not greater than the max
// and that the filter is a valid option
func (client *Client) verifyBuildsParams(params *url.Values) {
	limitParam := params.Get("limit")
	var parsedLimit int
	var limitErr error
	if limitParam != "" {
		parsedLimit, limitErr = strconv.Atoi(limitParam)
	}
	if limitErr != nil || (parsedLimit > 100 || parsedLimit < 0) {
		client.Logger.Printf("Invalid limit, defaulting to %d", BuildLimitMax)
		params.Set("limit", strconv.Itoa(BuildLimitMax))
	}
	filterParam := params.Get("filter")
	if filterParam != "" && !validBuildFilter(filterParam) {
		client.Logger.Printf("Invalid filter %s, defaulting to empty", filterParam)
		params.Del("filter")
	}
}

// validBuildFilter verifies that a filter is in the ValidBuildFilters
func validBuildFilter(filter string) bool {
	_, valid := ValidBuildFilters[strings.ToLower(filter)]
	return valid
}
