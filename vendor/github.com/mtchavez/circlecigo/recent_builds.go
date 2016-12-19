package circleci

import (
	"net/http"
	"net/url"
)

// RecentBuilds calls the /recent-builds endpoint to get recent builds across
// all followed projects. Can supply limit and offset params.
func (client *Client) RecentBuilds(params url.Values) ([]*Build, *APIResponse) {
	builds := []*Build{}
	apiResp := client.request(http.MethodGet, "recent-builds", params, nil, &builds)
	return builds, apiResp
}
