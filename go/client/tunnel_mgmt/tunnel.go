// Package tunnel_mgmt_client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package tunnel_mgmt_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	externalRef1 "github.com/Codename-Uranium/api/go/server/common"

	externalRef0 "github.com/Codename-Uranium/api/go/client/federation"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

const (
	Federation_keyScopes = "Federation_key.Scopes"
)

// Contains short statistic from the node
type PingResponse struct {
	IfRxBytes        int `json:"if_rx_bytes"`
	IfRxErrors       int `json:"if_rx_errors"`
	IfRxPackets      int `json:"if_rx_packets"`
	IfTxBytes        int `json:"if_tx_bytes"`
	IfTxErrors       int `json:"if_tx_errors"`
	IfTxPackets      int `json:"if_tx_packets"`
	PeersTotal       int `json:"peers_total"`
	PeersWithTraffic int `json:"peers_with_traffic"`
}

// TrustedKey defines model for TrustedKey.
type TrustedKey string

// TrustedKeyRecord defines model for TrustedKeyRecord.
type TrustedKeyRecord struct {
	Id  string     `json:"id"`
	Key TrustedKey `json:"key"`
}

// FederationSetAuthorizerKeysJSONBody defines parameters for FederationSetAuthorizerKeys.
type FederationSetAuthorizerKeysJSONBody []externalRef0.PublicKeyRecord

// FederationSetAuthorizerKeysJSONRequestBody defines body for FederationSetAuthorizerKeys for application/json ContentType.
type FederationSetAuthorizerKeysJSONRequestBody FederationSetAuthorizerKeysJSONBody

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
	// FederationPing request
	FederationPing(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FederationSetAuthorizerKeys request with any body
	FederationSetAuthorizerKeysWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FederationSetAuthorizerKeys(ctx context.Context, body FederationSetAuthorizerKeysJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AdminListTrustedKeys request
	AdminListTrustedKeys(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AdminDeleteTrustedKey request
	AdminDeleteTrustedKey(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AdminGetTrustedKey request
	AdminGetTrustedKey(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AdminAddTrustedKey request with any body
	AdminAddTrustedKeyWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AdminUpdateTrustedKey request with any body
	AdminUpdateTrustedKeyWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) FederationPing(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFederationPingRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) FederationSetAuthorizerKeysWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFederationSetAuthorizerKeysRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) FederationSetAuthorizerKeys(ctx context.Context, body FederationSetAuthorizerKeysJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFederationSetAuthorizerKeysRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AdminListTrustedKeys(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAdminListTrustedKeysRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AdminDeleteTrustedKey(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAdminDeleteTrustedKeyRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AdminGetTrustedKey(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAdminGetTrustedKeyRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AdminAddTrustedKeyWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAdminAddTrustedKeyRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AdminUpdateTrustedKeyWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAdminUpdateTrustedKeyRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewFederationPingRequest generates requests for FederationPing
func NewFederationPingRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tunnel/federation/ping")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewFederationSetAuthorizerKeysRequest calls the generic FederationSetAuthorizerKeys builder with application/json body
func NewFederationSetAuthorizerKeysRequest(server string, body FederationSetAuthorizerKeysJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewFederationSetAuthorizerKeysRequestWithBody(server, "application/json", bodyReader)
}

// NewFederationSetAuthorizerKeysRequestWithBody generates requests for FederationSetAuthorizerKeys with any type of body
func NewFederationSetAuthorizerKeysRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tunnel/federation/set-authorizer-keys")
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

	return req, nil
}

// NewAdminListTrustedKeysRequest generates requests for AdminListTrustedKeys
func NewAdminListTrustedKeysRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tunnel/trusted-keys")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAdminDeleteTrustedKeyRequest generates requests for AdminDeleteTrustedKey
func NewAdminDeleteTrustedKeyRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tunnel/trusted-keys/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAdminGetTrustedKeyRequest generates requests for AdminGetTrustedKey
func NewAdminGetTrustedKeyRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tunnel/trusted-keys/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAdminAddTrustedKeyRequestWithBody generates requests for AdminAddTrustedKey with any type of body
func NewAdminAddTrustedKeyRequestWithBody(server string, id string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tunnel/trusted-keys/%s", pathParam0)
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

	return req, nil
}

// NewAdminUpdateTrustedKeyRequestWithBody generates requests for AdminUpdateTrustedKey with any type of body
func NewAdminUpdateTrustedKeyRequestWithBody(server string, id string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tunnel/trusted-keys/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

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
	// FederationPing request
	FederationPingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*FederationPingResponse, error)

	// FederationSetAuthorizerKeys request with any body
	FederationSetAuthorizerKeysWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*FederationSetAuthorizerKeysResponse, error)

	FederationSetAuthorizerKeysWithResponse(ctx context.Context, body FederationSetAuthorizerKeysJSONRequestBody, reqEditors ...RequestEditorFn) (*FederationSetAuthorizerKeysResponse, error)

	// AdminListTrustedKeys request
	AdminListTrustedKeysWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*AdminListTrustedKeysResponse, error)

	// AdminDeleteTrustedKey request
	AdminDeleteTrustedKeyWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*AdminDeleteTrustedKeyResponse, error)

	// AdminGetTrustedKey request
	AdminGetTrustedKeyWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*AdminGetTrustedKeyResponse, error)

	// AdminAddTrustedKey request with any body
	AdminAddTrustedKeyWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AdminAddTrustedKeyResponse, error)

	// AdminUpdateTrustedKey request with any body
	AdminUpdateTrustedKeyWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AdminUpdateTrustedKeyResponse, error)
}

type FederationPingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PingResponse
	JSON401      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r FederationPingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FederationPingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type FederationSetAuthorizerKeysResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *externalRef1.Error
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r FederationSetAuthorizerKeysResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FederationSetAuthorizerKeysResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AdminListTrustedKeysResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]TrustedKeyRecord
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r AdminListTrustedKeysResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AdminListTrustedKeysResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AdminDeleteTrustedKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON404      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r AdminDeleteTrustedKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AdminDeleteTrustedKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AdminGetTrustedKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON404      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r AdminGetTrustedKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AdminGetTrustedKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AdminAddTrustedKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *externalRef1.Error
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON409      *externalRef1.Error
	JSON500      *externalRef1.Error
	JSON507      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r AdminAddTrustedKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AdminAddTrustedKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AdminUpdateTrustedKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *externalRef1.Error
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON404      *externalRef1.Error
	JSON409      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r AdminUpdateTrustedKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AdminUpdateTrustedKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// FederationPingWithResponse request returning *FederationPingResponse
func (c *ClientWithResponses) FederationPingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*FederationPingResponse, error) {
	rsp, err := c.FederationPing(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFederationPingResponse(rsp)
}

// FederationSetAuthorizerKeysWithBodyWithResponse request with arbitrary body returning *FederationSetAuthorizerKeysResponse
func (c *ClientWithResponses) FederationSetAuthorizerKeysWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*FederationSetAuthorizerKeysResponse, error) {
	rsp, err := c.FederationSetAuthorizerKeysWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFederationSetAuthorizerKeysResponse(rsp)
}

func (c *ClientWithResponses) FederationSetAuthorizerKeysWithResponse(ctx context.Context, body FederationSetAuthorizerKeysJSONRequestBody, reqEditors ...RequestEditorFn) (*FederationSetAuthorizerKeysResponse, error) {
	rsp, err := c.FederationSetAuthorizerKeys(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFederationSetAuthorizerKeysResponse(rsp)
}

// AdminListTrustedKeysWithResponse request returning *AdminListTrustedKeysResponse
func (c *ClientWithResponses) AdminListTrustedKeysWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*AdminListTrustedKeysResponse, error) {
	rsp, err := c.AdminListTrustedKeys(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAdminListTrustedKeysResponse(rsp)
}

// AdminDeleteTrustedKeyWithResponse request returning *AdminDeleteTrustedKeyResponse
func (c *ClientWithResponses) AdminDeleteTrustedKeyWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*AdminDeleteTrustedKeyResponse, error) {
	rsp, err := c.AdminDeleteTrustedKey(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAdminDeleteTrustedKeyResponse(rsp)
}

// AdminGetTrustedKeyWithResponse request returning *AdminGetTrustedKeyResponse
func (c *ClientWithResponses) AdminGetTrustedKeyWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*AdminGetTrustedKeyResponse, error) {
	rsp, err := c.AdminGetTrustedKey(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAdminGetTrustedKeyResponse(rsp)
}

// AdminAddTrustedKeyWithBodyWithResponse request with arbitrary body returning *AdminAddTrustedKeyResponse
func (c *ClientWithResponses) AdminAddTrustedKeyWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AdminAddTrustedKeyResponse, error) {
	rsp, err := c.AdminAddTrustedKeyWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAdminAddTrustedKeyResponse(rsp)
}

// AdminUpdateTrustedKeyWithBodyWithResponse request with arbitrary body returning *AdminUpdateTrustedKeyResponse
func (c *ClientWithResponses) AdminUpdateTrustedKeyWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AdminUpdateTrustedKeyResponse, error) {
	rsp, err := c.AdminUpdateTrustedKeyWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAdminUpdateTrustedKeyResponse(rsp)
}

// ParseFederationPingResponse parses an HTTP response from a FederationPingWithResponse call
func ParseFederationPingResponse(rsp *http.Response) (*FederationPingResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &FederationPingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PingResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseFederationSetAuthorizerKeysResponse parses an HTTP response from a FederationSetAuthorizerKeysWithResponse call
func ParseFederationSetAuthorizerKeysResponse(rsp *http.Response) (*FederationSetAuthorizerKeysResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &FederationSetAuthorizerKeysResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseAdminListTrustedKeysResponse parses an HTTP response from a AdminListTrustedKeysWithResponse call
func ParseAdminListTrustedKeysResponse(rsp *http.Response) (*AdminListTrustedKeysResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AdminListTrustedKeysResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []TrustedKeyRecord
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseAdminDeleteTrustedKeyResponse parses an HTTP response from a AdminDeleteTrustedKeyWithResponse call
func ParseAdminDeleteTrustedKeyResponse(rsp *http.Response) (*AdminDeleteTrustedKeyResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AdminDeleteTrustedKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseAdminGetTrustedKeyResponse parses an HTTP response from a AdminGetTrustedKeyWithResponse call
func ParseAdminGetTrustedKeyResponse(rsp *http.Response) (*AdminGetTrustedKeyResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AdminGetTrustedKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseAdminAddTrustedKeyResponse parses an HTTP response from a AdminAddTrustedKeyWithResponse call
func ParseAdminAddTrustedKeyResponse(rsp *http.Response) (*AdminAddTrustedKeyResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AdminAddTrustedKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 507:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON507 = &dest

	}

	return response, nil
}

// ParseAdminUpdateTrustedKeyResponse parses an HTTP response from a AdminUpdateTrustedKeyWithResponse call
func ParseAdminUpdateTrustedKeyResponse(rsp *http.Response) (*AdminUpdateTrustedKeyResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AdminUpdateTrustedKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
