package circleci

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	defaultHost    = "circleci.com"
	defaultPort    = 443
	defaultScheme  = "https"
	defaultVersion = "v1"
)

var (
	defaultURL    = &url.URL{Host: defaultHost, Scheme: defaultScheme, Path: "/api/" + defaultVersion + "/"}
	defaultLogger = log.New(os.Stdout, "[circleci] ", log.LstdFlags)
	defaultClient = &Client{BaseURL: defaultURL, Logger: defaultLogger, httpClient: http.DefaultClient}
)

// APIResponse - Wraps an API request and gets returned from API interactions
// to access the raw response and status of the request.
type APIResponse struct {
	Response      *http.Response
	Body          []byte
	Error         error
	ErrorResponse *APIMessageResponse
}

// APIMessageResponse - Handles generic API responses from CircleCI
type APIMessageResponse struct {
	Message string `json:"message"`
}

// Success - Returns a boolean for determining if an APIResponse is successful
// Will return false for any non 2xx status codes.
func (resp *APIResponse) Success() bool {
	if resp.Error != nil {
		return false
	}
	statusCode := resp.Response.StatusCode
	if statusCode >= 200 && statusCode < 300 {
		return true
	}
	return false
}

// Client is for configuring settings to interact with the
// Client API to make requests
type Client struct {
	BaseURL    *url.URL
	Logger     *log.Logger
	Token      string
	httpClient *http.Client
}

type requestBody struct {
	io.Reader
}

func (rBody requestBody) Close() error {
	return nil
}

// NewClient takes an auth token and returns a new Client object
func NewClient(token string) *Client {
	return &Client{
		BaseURL:    defaultURL,
		Token:      token,
		Logger:     defaultLogger,
		httpClient: http.DefaultClient,
	}
}

func (client *Client) String() string {
	clientDetails := map[string]interface{}{
		"base_url": client.baseURL().String(),
		"token":    client.Token,
	}
	jsonClient, _ := json.Marshal(clientDetails)
	return string(jsonClient)
}

func (client *Client) request(method, urlString string, params url.Values, body interface{}, response interface{}) *APIResponse {
	apiResp := &APIResponse{}
	if params == nil {
		params = url.Values{}
	}
	if client.Token != "" {
		params.Set("circle-token", client.Token)
	}
	reqPath := &url.URL{Path: urlString, RawQuery: params.Encode()}
	reqURL := client.baseURL().ResolveReference(reqPath)
	httpReq, reqErr := newClientRequest(method, reqURL.String(), body)
	if reqErr != nil {
		apiResp.Error = reqErr
		return apiResp
	}
	httpResp, respErr := client.httpClient.Do(httpReq)
	if respErr != nil {
		apiResp.Error = respErr
		return apiResp
	}
	apiResp.Response = httpResp
	defer httpResp.Body.Close()
	bodyBytes, bodyErr := ioutil.ReadAll(httpResp.Body)
	if bodyErr != nil {
		apiResp.Error = bodyErr
		return apiResp
	}
	apiResp.Body = bodyBytes
	if httpResp.StatusCode == http.StatusNotFound {
		message := &APIMessageResponse{}
		json.Unmarshal(bodyBytes, message)
		apiResp.ErrorResponse = message
	}
	unmarshalErr := json.Unmarshal(bodyBytes, response)
	if unmarshalErr != nil {
		apiResp.Error = unmarshalErr
		return apiResp
	}
	return apiResp
}

func newClientRequest(method, urlString string, body interface{}) (*http.Request, error) {
	httpReq, reqErr := http.NewRequest(method, urlString, nil)
	if reqErr != nil {
		return nil, reqErr
	}
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("Content-Type", "application/json")
	if method != http.MethodGet && body != nil {
		bodyBytes, marshalErr := json.Marshal(body)
		if marshalErr != nil {
			return httpReq, marshalErr
		}
		httpReq.Body = &requestBody{bytes.NewBuffer(bodyBytes)}
	}
	return httpReq, nil
}

func (client *Client) baseURL() *url.URL {
	if client.BaseURL != nil {
		return client.BaseURL
	}
	return defaultURL
}
