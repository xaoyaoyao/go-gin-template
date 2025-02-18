// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

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
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
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

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// InitializationData request
	InitializationData(ctx context.Context, params *InitializationDataParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SignupWithBody request with any body
	SignupWithBody(ctx context.Context, params *SignupParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Signup(ctx context.Context, params *SignupParams, body SignupJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) InitializationData(ctx context.Context, params *InitializationDataParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewInitializationDataRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SignupWithBody(ctx context.Context, params *SignupParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSignupRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Signup(ctx context.Context, params *SignupParams, body SignupJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSignupRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewInitializationDataRequest generates requests for InitializationData
func NewInitializationDataRequest(server string, params *InitializationDataParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/initialize")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.QueryParams != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "queryParams", runtime.ParamLocationQuery, *params.QueryParams); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewSignupRequest calls the generic Signup builder with application/json body
func NewSignupRequest(server string, params *SignupParams, body SignupJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSignupRequestWithBody(server, params, "application/json", bodyReader)
}

// NewSignupRequestWithBody generates requests for Signup with any type of body
func NewSignupRequestWithBody(server string, params *SignupParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/users/signup")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "GO-API-KEY", runtime.ParamLocationHeader, params.GOAPIKEY)
		if err != nil {
			return nil, err
		}

		req.Header.Set("GO-API-KEY", headerParam0)

	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
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
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// InitializationDataWithResponse request
	InitializationDataWithResponse(ctx context.Context, params *InitializationDataParams, reqEditors ...RequestEditorFn) (*InitializationDataResponse, error)

	// SignupWithBodyWithResponse request with any body
	SignupWithBodyWithResponse(ctx context.Context, params *SignupParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SignupResponse, error)

	SignupWithResponse(ctx context.Context, params *SignupParams, body SignupJSONRequestBody, reqEditors ...RequestEditorFn) (*SignupResponse, error)
}

type InitializationDataResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ResponseEntity
}

// Status returns HTTPResponse.Status
func (r InitializationDataResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r InitializationDataResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SignupResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ResponseEntity
	JSON404      *ResponseEntity
	JSON500      *ResponseEntity
}

// Status returns HTTPResponse.Status
func (r SignupResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SignupResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// InitializationDataWithResponse request returning *InitializationDataResponse
func (c *ClientWithResponses) InitializationDataWithResponse(ctx context.Context, params *InitializationDataParams, reqEditors ...RequestEditorFn) (*InitializationDataResponse, error) {
	rsp, err := c.InitializationData(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseInitializationDataResponse(rsp)
}

// SignupWithBodyWithResponse request with arbitrary body returning *SignupResponse
func (c *ClientWithResponses) SignupWithBodyWithResponse(ctx context.Context, params *SignupParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SignupResponse, error) {
	rsp, err := c.SignupWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSignupResponse(rsp)
}

func (c *ClientWithResponses) SignupWithResponse(ctx context.Context, params *SignupParams, body SignupJSONRequestBody, reqEditors ...RequestEditorFn) (*SignupResponse, error) {
	rsp, err := c.Signup(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSignupResponse(rsp)
}

// ParseInitializationDataResponse parses an HTTP response from a InitializationDataWithResponse call
func ParseInitializationDataResponse(rsp *http.Response) (*InitializationDataResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &InitializationDataResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ResponseEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseSignupResponse parses an HTTP response from a SignupWithResponse call
func ParseSignupResponse(rsp *http.Response) (*SignupResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SignupResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ResponseEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ResponseEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ResponseEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
