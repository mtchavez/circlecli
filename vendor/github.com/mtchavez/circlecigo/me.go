package circleci

import "net/http"

// Me - CircleCI me endpoint response to get account information
type Me struct {
	Admin             bool     `json:"admin"`
	AllEmails         []string `json:"all_emails"`
	AnalyticsID       string   `json:"analytics_id"`
	AvatarURL         string   `json:"avatar_url"`
	Containers        int      `json:"containers"`
	DevAdmin          bool     `json:"dev_admin"`
	GitubID           int      `json:"github_id"`
	GithubOAuthScopes []string `json:"github_oauth_scopes"`
	Login             string   `json:"login"`
	Name              string   `json:"name"`
	SelectedEmail     string   `json:"selected_email"`
}

// Me calls the /me CircleCI endpoint to get account information. Requires auth token.
func (client *Client) Me() (*Me, *APIResponse) {
	me := &Me{}
	apiResp := client.request(http.MethodGet, "me", nil, nil, me)
	return me, apiResp
}
