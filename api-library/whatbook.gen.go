// Package WhatbookV1 provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package WhatbookV1

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Book defines model for Book.
type Book struct {
	Author   string `json:"Author"`
	Genre    string `json:"Genre"`
	NumPages uint32 `json:"NumPages"`
	Rating   uint32 `json:"Rating"`
	Title    string `json:"Title"`
	Year     uint32 `json:"Year"`
}

// Error defines model for Error.
type Error struct {
	Message *string `json:"message,omitempty"`
}

// GetBooksParams defines parameters for GetBooks.
type GetBooksParams struct {
	Author   *string `json:"author,omitempty"`
	Genre    *string `json:"genre,omitempty"`
	NumPages *uint32 `json:"numPages,omitempty"`
	Era      *string `json:"era,omitempty"`
}

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
		client.Client = http.DefaultClient
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
	// GetBooks request
	GetBooks(ctx context.Context, params *GetBooksParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTest request
	GetTest(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetBooks(ctx context.Context, params *GetBooksParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetBooksRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetTest(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTestRequest(c.Server)
	if err != nil {
		return nil, err
	}
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetBooksRequest generates requests for GetBooks
func NewGetBooksRequest(server string, params *GetBooksParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/books")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.Author != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "author", *params.Author); err != nil {
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

	if params.Genre != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "genre", *params.Genre); err != nil {
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

	if params.NumPages != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "numPages", *params.NumPages); err != nil {
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

	if params.Era != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "era", *params.Era); err != nil {
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

// NewGetTestRequest generates requests for GetTest
func NewGetTestRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/test")
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

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	req = req.WithContext(ctx)
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
	// GetBooks request
	GetBooksWithResponse(ctx context.Context, params *GetBooksParams) (*GetBooksResponse, error)

	// GetTest request
	GetTestWithResponse(ctx context.Context) (*GetTestResponse, error)
}

type GetBooksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Book
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r GetBooksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetBooksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTestResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetTestResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTestResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetBooksWithResponse request returning *GetBooksResponse
func (c *ClientWithResponses) GetBooksWithResponse(ctx context.Context, params *GetBooksParams) (*GetBooksResponse, error) {
	rsp, err := c.GetBooks(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetBooksResponse(rsp)
}

// GetTestWithResponse request returning *GetTestResponse
func (c *ClientWithResponses) GetTestWithResponse(ctx context.Context) (*GetTestResponse, error) {
	rsp, err := c.GetTest(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetTestResponse(rsp)
}

// ParseGetBooksResponse parses an HTTP response from a GetBooksWithResponse call
func ParseGetBooksResponse(rsp *http.Response) (*GetBooksResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetBooksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Book
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetTestResponse parses an HTTP response from a GetTestWithResponse call
func ParseGetTestResponse(rsp *http.Response) (*GetTestResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetTestResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /books)
	GetBooks(ctx echo.Context, params GetBooksParams) error

	// (GET /test)
	GetTest(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetBooks converts echo context to params.
func (w *ServerInterfaceWrapper) GetBooks(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBooksParams
	// ------------- Optional query parameter "author" -------------

	err = runtime.BindQueryParameter("form", true, false, "author", ctx.QueryParams(), &params.Author)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter author: %s", err))
	}

	// ------------- Optional query parameter "genre" -------------

	err = runtime.BindQueryParameter("form", true, false, "genre", ctx.QueryParams(), &params.Genre)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter genre: %s", err))
	}

	// ------------- Optional query parameter "numPages" -------------

	err = runtime.BindQueryParameter("form", true, false, "numPages", ctx.QueryParams(), &params.NumPages)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter numPages: %s", err))
	}

	// ------------- Optional query parameter "era" -------------

	err = runtime.BindQueryParameter("form", true, false, "era", ctx.QueryParams(), &params.Era)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter era: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBooks(ctx, params)
	return err
}

// GetTest converts echo context to params.
func (w *ServerInterfaceWrapper) GetTest(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTest(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/books", wrapper.GetBooks)
	router.GET(baseURL+"/test", wrapper.GetTest)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/5RTTWvcMBD9K2bao1lvEnrRLYUQQqENSaCUsAdFO2srXX1kNHZZFv/3MvJ6PxpTNhdL",
	"lmbevHnztAUTXAwePSdQW0imQafz9msIv2WNFCISW8yn1y03gWTHm4igIDFZX0Nfwi16wsmb76271/UA",
	"sArkNIOC1nq+uoRyDLeesUaS+AfNknlm9JPl9XTdX6jpPJS+BMK31hIuQT2PXY7YY29HnezA91wXe8zw",
	"8oqGpfwN0SDVqYQOU9L1FOP+HYgcWb8KErzEZMhGtsGDgoebx6fi+v6uWAUquMHiZ6NZZla8yIfQBOfQ",
	"L7XEFwmps0Y64EEtGMOhhA4pDaAXs7kQDxG9jhYUXM3mM9Eram4y+UrQ865GnmCFTBY7TJlF+oeGyCZK",
	"5J+7JSi4xUwi5RKkHTJSAvW8BStwby3SBkrw2glnPY5lsOmkgtOZ9W5+H070h4Efcs/x0zQckj5BQt86",
	"MZxZ65Ss0WsowYUlkj9y1J7jQlyaYvBpcNLlfC6LCZ7R52noGNfWZH2r1yQj2R5Vs4wuJ34mXIGCT9Xh",
	"9Ve7p19lUxycqIn0ZjDi6ajz3IrcXzGykrwvHyT1Py7DC5oo/ojUIRX7+76EijHxkTHfGe1J7qcVPAX/",
	"8S1D9iWkXGYwZEtrUNAwx6Sq6k+jWSw+M8FV3QX0i/5vAAAA//8IyIVATAUAAA==",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

