// Package feeds provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package feeds

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"github.com/amzapi/selling-partner-api-sdk/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// CreateFeedDocument request  with any body
	CreateFeedDocumentWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	CreateFeedDocument(ctx context.Context, body CreateFeedDocumentJSONRequestBody) (*http.Response, error)

	// GetFeedDocument request
	GetFeedDocument(ctx context.Context, feedDocumentId string) (*http.Response, error)

	// GetFeeds request
	GetFeeds(ctx context.Context, params *GetFeedsParams) (*http.Response, error)

	// CreateFeed request  with any body
	CreateFeedWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	CreateFeed(ctx context.Context, body CreateFeedJSONRequestBody) (*http.Response, error)

	// CancelFeed request
	CancelFeed(ctx context.Context, feedId string) (*http.Response, error)

	// GetFeed request
	GetFeed(ctx context.Context, feedId string) (*http.Response, error)
}

func (c *Client) CreateFeedDocumentWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewCreateFeedDocumentRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CreateFeedDocument(ctx context.Context, body CreateFeedDocumentJSONRequestBody) (*http.Response, error) {
	req, err := NewCreateFeedDocumentRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetFeedDocument(ctx context.Context, feedDocumentId string) (*http.Response, error) {
	req, err := NewGetFeedDocumentRequest(c.Endpoint, feedDocumentId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetFeeds(ctx context.Context, params *GetFeedsParams) (*http.Response, error) {
	req, err := NewGetFeedsRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CreateFeedWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewCreateFeedRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CreateFeed(ctx context.Context, body CreateFeedJSONRequestBody) (*http.Response, error) {
	req, err := NewCreateFeedRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) CancelFeed(ctx context.Context, feedId string) (*http.Response, error) {
	req, err := NewCancelFeedRequest(c.Endpoint, feedId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetFeed(ctx context.Context, feedId string) (*http.Response, error) {
	req, err := NewGetFeedRequest(c.Endpoint, feedId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewCreateFeedDocumentRequest calls the generic CreateFeedDocument builder with application/json body
func NewCreateFeedDocumentRequest(endpoint string, body CreateFeedDocumentJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateFeedDocumentRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewCreateFeedDocumentRequestWithBody generates requests for CreateFeedDocument with any type of body
func NewCreateFeedDocumentRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/feeds/2021-06-30/documents")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetFeedDocumentRequest generates requests for GetFeedDocument
func NewGetFeedDocumentRequest(endpoint string, feedDocumentId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "feedDocumentId", feedDocumentId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/feeds/2021-06-30/documents/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetFeedsRequest generates requests for GetFeeds
func NewGetFeedsRequest(endpoint string, params *GetFeedsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/feeds/2021-06-30/feeds")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.FeedTypes != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "feedTypes", *params.FeedTypes); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.MarketplaceIds != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "marketplaceIds", *params.MarketplaceIds); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageSize != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "pageSize", *params.PageSize); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.ProcessingStatuses != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "processingStatuses", *params.ProcessingStatuses); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.CreatedSince != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "createdSince", *params.CreatedSince); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.CreatedUntil != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "createdUntil", *params.CreatedUntil); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "nextToken", *params.NextToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateFeedRequest calls the generic CreateFeed builder with application/json body
func NewCreateFeedRequest(endpoint string, body CreateFeedJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateFeedRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewCreateFeedRequestWithBody generates requests for CreateFeed with any type of body
func NewCreateFeedRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/feeds/2021-06-30/feeds")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewCancelFeedRequest generates requests for CancelFeed
func NewCancelFeedRequest(endpoint string, feedId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "feedId", feedId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/feeds/2021-06-30/feeds/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetFeedRequest generates requests for GetFeed
func NewGetFeedRequest(endpoint string, feedId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "feedId", feedId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/feeds/2021-06-30/feeds/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// CreateFeedDocument request  with any body
	CreateFeedDocumentWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateFeedDocumentResp, error)

	CreateFeedDocumentWithResponse(ctx context.Context, body CreateFeedDocumentJSONRequestBody) (*CreateFeedDocumentResp, error)

	// GetFeedDocument request
	GetFeedDocumentWithResponse(ctx context.Context, feedDocumentId string) (*GetFeedDocumentResp, error)

	// GetFeeds request
	GetFeedsWithResponse(ctx context.Context, params *GetFeedsParams) (*GetFeedsResp, error)

	// CreateFeed request  with any body
	CreateFeedWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateFeedResp, error)

	CreateFeedWithResponse(ctx context.Context, body CreateFeedJSONRequestBody) (*CreateFeedResp, error)

	// CancelFeed request
	CancelFeedWithResponse(ctx context.Context, feedId string) (*CancelFeedResp, error)

	// GetFeed request
	GetFeedWithResponse(ctx context.Context, feedId string) (*GetFeedResp, error)
}

type CreateFeedDocumentResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CreateFeedDocumentResponse
}

// Status returns HTTPResponse.Status
func (r CreateFeedDocumentResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateFeedDocumentResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetFeedDocumentResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetFeedDocumentResponse
}

// Status returns HTTPResponse.Status
func (r GetFeedDocumentResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetFeedDocumentResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetFeedsResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetFeedsResponse
}

// Status returns HTTPResponse.Status
func (r GetFeedsResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetFeedsResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateFeedResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CreateFeedResponse
}

// Status returns HTTPResponse.Status
func (r CreateFeedResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateFeedResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CancelFeedResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CancelFeedResponse
}

// Status returns HTTPResponse.Status
func (r CancelFeedResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CancelFeedResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetFeedResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetFeedResponse
}

// Status returns HTTPResponse.Status
func (r GetFeedResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetFeedResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreateFeedDocumentWithBodyWithResponse request with arbitrary body returning *CreateFeedDocumentResponse
func (c *ClientWithResponses) CreateFeedDocumentWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateFeedDocumentResp, error) {
	rsp, err := c.CreateFeedDocumentWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateFeedDocumentResp(rsp)
}

func (c *ClientWithResponses) CreateFeedDocumentWithResponse(ctx context.Context, body CreateFeedDocumentJSONRequestBody) (*CreateFeedDocumentResp, error) {
	rsp, err := c.CreateFeedDocument(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateFeedDocumentResp(rsp)
}

// GetFeedDocumentWithResponse request returning *GetFeedDocumentResponse
func (c *ClientWithResponses) GetFeedDocumentWithResponse(ctx context.Context, feedDocumentId string) (*GetFeedDocumentResp, error) {
	rsp, err := c.GetFeedDocument(ctx, feedDocumentId)
	if err != nil {
		return nil, err
	}
	return ParseGetFeedDocumentResp(rsp)
}

// GetFeedsWithResponse request returning *GetFeedsResponse
func (c *ClientWithResponses) GetFeedsWithResponse(ctx context.Context, params *GetFeedsParams) (*GetFeedsResp, error) {
	rsp, err := c.GetFeeds(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetFeedsResp(rsp)
}

// CreateFeedWithBodyWithResponse request with arbitrary body returning *CreateFeedResponse
func (c *ClientWithResponses) CreateFeedWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateFeedResp, error) {
	rsp, err := c.CreateFeedWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateFeedResp(rsp)
}

func (c *ClientWithResponses) CreateFeedWithResponse(ctx context.Context, body CreateFeedJSONRequestBody) (*CreateFeedResp, error) {
	rsp, err := c.CreateFeed(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateFeedResp(rsp)
}

// CancelFeedWithResponse request returning *CancelFeedResponse
func (c *ClientWithResponses) CancelFeedWithResponse(ctx context.Context, feedId string) (*CancelFeedResp, error) {
	rsp, err := c.CancelFeed(ctx, feedId)
	if err != nil {
		return nil, err
	}
	return ParseCancelFeedResp(rsp)
}

// GetFeedWithResponse request returning *GetFeedResponse
func (c *ClientWithResponses) GetFeedWithResponse(ctx context.Context, feedId string) (*GetFeedResp, error) {
	rsp, err := c.GetFeed(ctx, feedId)
	if err != nil {
		return nil, err
	}
	return ParseGetFeedResp(rsp)
}

// ParseCreateFeedDocumentResp parses an HTTP response from a CreateFeedDocumentWithResponse call
func ParseCreateFeedDocumentResp(rsp *http.Response) (*CreateFeedDocumentResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CreateFeedDocumentResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CreateFeedDocumentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetFeedDocumentResp parses an HTTP response from a GetFeedDocumentWithResponse call
func ParseGetFeedDocumentResp(rsp *http.Response) (*GetFeedDocumentResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetFeedDocumentResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetFeedDocumentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetFeedsResp parses an HTTP response from a GetFeedsWithResponse call
func ParseGetFeedsResp(rsp *http.Response) (*GetFeedsResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetFeedsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetFeedsResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCreateFeedResp parses an HTTP response from a CreateFeedWithResponse call
func ParseCreateFeedResp(rsp *http.Response) (*CreateFeedResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CreateFeedResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CreateFeedResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCancelFeedResp parses an HTTP response from a CancelFeedWithResponse call
func ParseCancelFeedResp(rsp *http.Response) (*CancelFeedResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CancelFeedResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CancelFeedResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetFeedResp parses an HTTP response from a GetFeedWithResponse call
func ParseGetFeedResp(rsp *http.Response) (*GetFeedResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetFeedResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetFeedResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}
