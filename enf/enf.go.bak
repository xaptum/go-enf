package enf

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

const (
	ProjectURL     = "github.com/xaptum/go-enf"
	ProjectVersion = "0.4.0"

	HttpsScheme = "https"

	HeaderToken       = "Authorization"
	HeaderTokenFormat = "Bearer %s"

	MediaTypeJSON = "application/json"
)

var (
	defaultUserAgent = fmt.Sprintf("go-enf/%s (+%s; %s)", ProjectVersion, ProjectURL, runtime.Version())
)

type HttpMethod string

const (
	Get    HttpMethod = "GET"
	Put    HttpMethod = "PUT"
	Post   HttpMethod = "POST"
	Delete HttpMethod = "DELETE"
)

type QueryParams map[string]string
type RequestHeaders map[string]string

type Request struct {
	url         string
	headers     RequestHeaders
	queryParams QueryParams
	body        interface{}
}

type Response struct {
	StatusCode int
	Method     HttpMethod
	Url        string
	Headers    map[string][]string
}

type ErrorCodeText struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type ErrorReason struct {
	Reason string `json:"reason"`
}

// ErrorResponse represents the error response from the API.
type ErrorResponse struct {
	Response         *Response      `json:"-"`
	ErrorMessage     *ErrorCodeText `json:"error"`
	XiamErrorMessage *ErrorReason   `json:"xiam_error"`
}

func (e *ErrorResponse) Error() string {
	var msg string

	if nil == e.ErrorMessage {
		msg = fmt.Sprintf("%v %v: [%d] %v - %v",
			e.Response.Method, e.Response.Url,
			e.Response.StatusCode, e.ErrorMessage.Code, e.ErrorMessage.Text)
	} else if nil == e.XiamErrorMessage {
		msg = fmt.Sprintf("%v %v: [%d] %v",
			e.Response.Method, e.Response.Url,
			e.Response.StatusCode, e.XiamErrorMessage.Reason)
	} else {
		msg = "UNKNOWN_ERROR: server did not respond with properly formatted error message."
	}

	return msg
}

type TokenSource interface {
	Token() string
}

type StaticTokenSource struct {
	token string
}

func (ts *StaticTokenSource) Token() string {
	return ts.token
}

// Client represents a wrapper for the HTTP client that communicates with the API.
type Client struct {
	// HTTP client used to communicate with the API.
	httpClient *http.Client

	// The base URL for API requests parsed from the Domain
	baseUrl *url.URL

	// The API token for authenticating with the API
	authToken string

	// User agent used when communicating with the ENF API.
	userAgent string

	// Reuse a single struct instead of allocating one for each service on the heap
	service Service

	// Services used for talking to different parts of the ENF API.
	/*DNS      *DNSService
	Domain   *DomainService
	Endpoint *EndpointService
	Firewall *FirewallService
	Network  *NetworkService
	User     *UserService*/
}

type Service struct {
	client *Client
}

// Bool is a helper function that creates a new value and returns a
// pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper function that creates a new value and returns a
// pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper function that creates a new value and returns a
// pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper function that creates a new value and returns a
// pointer to it.
func String(v string) *string { return &v }

// Time is a helper function that creates a new value and returns a
// pointer to it.
func Time(v time.Time) *time.Time { return &v }

// NewClient returns a new ENF API client for the provided domain. If
// a nil httpClient is provided, a new http.Client will be used.  To
// use API methods which require authentication, provide an
// http.Client that will perform the authentication for you (such as
// that provided by TokenAuthClient in this library)
func NewClient(host string, ts TokenSource, httpClient *http.Client) (*Client, error) {
	// parse host
	baseUrl, err := url.Parse(host)
	if nil != err {
		return nil, err
	}
	if "" == baseUrl.Scheme {
		baseUrl.Scheme = HttpsScheme
	}

	// create http client
	if nil == httpClient {
		httpClient = &http.Client{
			Timeout: time.Second * 20,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).Dial,
			},
		}
	}

	// get api token
	var token string
	if nil != ts {
		token = ts.Token()
	}

	// create enf api client
	c := &Client{
		httpClient: httpClient,
		baseUrl:    baseUrl,
		authToken:  token,
		userAgent:  defaultUserAgent,
	}
	c.service.client = c
	/*	c.Domains = (*DomainService)(&c.service)
		c.Endpoint = (*EndpointService)(&c.service)
		c.DNS = (*DNSService)(&c.service)
		c.Firewall = (*FirewallService)(&c.service)
		c.Network = (*NetworkService)(&c.service)
		c.User = (*UserService)(&c.service)*/
	return c, nil
}

func NewRequest(url string, headers RequestHeaders, queryParams QueryParams, body interface{}) Request {
	return Request{
		url:         url,
		headers:     headers,
		queryParams: queryParams,
		body:        body,
	}
}

func (c *Client) get(ctx context.Context, request Request, v interface{}) (*Response, error) {
	// send request
	return c.do(ctx, Get, request, v)
}

func (c *Client) post(ctx context.Context, request Request, v interface{}) (*Response, error) {
	// send request
	return c.do(ctx, Post, request, v)
}

func (c *Client) put(ctx context.Context, request Request, v interface{}) (*Response, error) {
	// send request
	return c.do(ctx, Put, request, v)
}

func (c *Client) delete(ctx context.Context, request Request, v interface{}) (*Response, error) {
	// send request
	return c.do(ctx, Delete, request, v)
}

// Do executes an HTTP request.
func (c *Client) do(ctx context.Context, method HttpMethod, request Request, v interface{}) (*Response, error) {
	// make http.Request
	req, err := c.newHttpRequest(method, request)
	if nil != err {
		return nil, err
	}

	// create request with context
	req = req.WithContext(ctx)

	// send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,.
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, err
		}
	}
	defer resp.Body.Close()

	return c.processHttpResponse(resp, v)
}

func (c *Client) newHttpRequest(method HttpMethod, request Request) (*http.Request, error) {
	// add query parameters
	request.url = c.addQueryParams(request.url, request.queryParams)

	// encode request body into json
	body, err := c.json(request.body)
	if nil != err {
		return nil, err
	}

	// create http requeuest
	httpRequest, err := http.NewRequest(string(method), request.url, bytes.NewBuffer(body))
	if nil != err {
		return nil, err
	}

	// add headers
	for key, value := range request.headers {
		httpRequest.Header.Set(key, value)
	}

	// add missing content type header
	_, exists := httpRequest.Header["Content-Type"]
	if len(body) > 0 && !exists {
		httpRequest.Header.Set("Content-Type", "application/json")
	}

	// add missing accept header
	_, exists = httpRequest.Header["Accept"]
	if !exists {
		httpRequest.Header.Set("Accept", "application/json")
	}

	// add authorization header
	if "" != c.authToken {
		httpRequest.Header.Set(HeaderToken, fmt.Sprintf(HeaderTokenFormat, c.authToken))
	}

	// add user-agent header
	httpRequest.Header.Set("User-Agent", c.userAgent)

	return httpRequest, nil
}

// process http response
func (c *Client) processHttpResponse(httpResp *http.Response, v interface{}) (*Response, error) {
	// create response object
	response := &Response{
		StatusCode: httpResp.StatusCode,
		Headers:    httpResp.Header,
		Method:     HttpMethod(httpResp.Request.Method),
		Url:        httpResp.Request.URL.String(),
	}

	// create an error response object
	var errResp = &ErrorResponse{
		Response: response,
	}

	switch httpResp.StatusCode {
	case 200, 201:
		// parse response body for api response json
		if err := c.parseJson(httpResp.Body, v); nil != err {
			return nil, err
		}
		return response, nil

	case 400, 401:
		// parse response body for error json
		if err := c.parseJson(httpResp.Body, errResp); nil != err {
			// not a json error message
			return nil, err
		}
		// return error response
		return nil, errResp

	case 403:
		// method not found error
		errResp.ErrorMessage.Code = "http_error"
		errResp.ErrorMessage.Text = "Method Not Found"
		return nil, errResp

	case 415:
		// method not found error
		errResp.ErrorMessage.Code = "http_error"
		errResp.ErrorMessage.Text = "Unsupported Media Type"
		return nil, errResp

	case 404:
		// parse response body for error json
		if err := c.parseJson(httpResp.Body, errResp); nil != err {
			// not a json error message. assume URL not found
			errResp.ErrorMessage.Code = "http_error"
			errResp.ErrorMessage.Text = "Not Found"
		}
		// return error response
		return nil, errResp

	case 500:
		// parse response body for error json
		if err := c.parseJson(httpResp.Body, errResp); nil != err {
			// not a json error message. encode a genric error message
			errResp.ErrorMessage.Code = "server_error"
			errResp.ErrorMessage.Text = "Server error received without details"
		}
		// return error response
		return nil, errResp

	default:
		return nil, fmt.Errorf("unexpected status code %d", httpResp.StatusCode)

	}
}

func (c *Client) addQueryParams(baseUrl string, queryParams QueryParams) string {
	if 0 == len(queryParams) {
		return baseUrl
	}

	baseUrl += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseUrl + params.Encode()
}

func (c *Client) json(body interface{}) ([]byte, error) {
	// create buffer
	buf := new(bytes.Buffer)

	// encode to json if needed
	if nil != body {
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	// return json
	return buf.Bytes(), nil
}

func (c *Client) parseJson(buf io.Reader, v interface{}) error {
	// check if v is not nil
	if nil == v {
		// no need to parse. just return
		return nil
	}

	var err error

	if w, ok := v.(io.Writer); ok {
		_, err = io.Copy(w, buf)
	} else {
		err = json.NewDecoder(buf).Decode(v)
		if io.EOF == err {
			err = nil // ignore ERO errors caused by empty response body
		}
	}

	return err
}

// All the exported methods in this file are designed to be general-purpose HTTP helpers. These methods
// will accept any request struct, and support any struct type you want the response to be stored in.
// For usage examples, see the methods in network.go or firewall.go
