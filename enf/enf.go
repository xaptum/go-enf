package enf

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
)

const (
	projectURL     = "github.com/xaptum/go-enf"
	projectVersion = "0.2.2"

	defaultScheme = "https"

	headerToken       = "Authorization"
	headerTokenFormat = "Bearer %s"

	mediaTypeJSON = "application/json"
)

var (
	defaultUserAgent       = fmt.Sprintf("go-enf/%s (+%s; %s)", projectVersion, projectURL, runtime.Version())
	wantAcceptHeaders      = []string{mediaTypeJSON}
	wantContentTypeHeaders = []string{mediaTypeJSON}
)

// Client represents a wrapper for the HTTP client that communicates with the API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// The ENF customer domain
	Domain string

	// The base URL for API requests parsed from the Domain
	BaseURL *url.URL

	// User agent used when communicating with the ENF API.
	UserAgent string

	// The API token for authenticating with the API
	APIToken string

	// Reuse a single struct instead of allocating one for each service on the heap
	common service

	// Services used for talking to different parts of the ENF API.
	Auth     *AuthService
	Firewall *FirewallService
	Network  *NetworkService
	DNS      *DNSService
	Domains  *DomainService
	Endpoint *EndpointService
}

type service struct {
	client *Client
}

// All the exported methods in this file are designed to be general-purpose HTTP helpers. These methods
// will accept any request struct, and support any struct type you want the response to be stored in.
// For usage examples, see the methods in network.go or firewall.go

// get makes a get request to the given path and stores the response in the given body object.
func (c *Client) get(ctx context.Context, path string, body interface{}) (interface{}, *http.Response, error) {
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	return c.makeRequest(ctx, req, body)
}

// post makes post requests to the given path with the given fields and stores the response in the given body object.
func (c *Client) post(ctx context.Context, path string, body interface{}, fields interface{}) (interface{}, *http.Response, error) {
	req, err := c.NewRequest("POST", path, fields)
	if err != nil {
		return nil, nil, err
	}
	return c.makeRequest(ctx, req, body)
}

// put makes put requests to the given path with the given fields and stores the response in the given body object.
func (c *Client) put(ctx context.Context, path string, body interface{}, fields interface{}) (interface{}, *http.Response, error) {
	req, err := c.NewRequest("PUT", path, fields)
	if err != nil {
		return nil, nil, err
	}
	return c.makeRequest(ctx, req, body)
}

// delete makes delete requests to the given path.
func (c *Client) delete(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(ctx, req, nil)
}

// makeRequest makes the given request, and stores the result into the given body.
func (c *Client) makeRequest(ctx context.Context, req *http.Request, body interface{}) (interface{}, *http.Response, error) {
	resp, err := c.Do(ctx, req, body)
	return body, resp, err
}

// NewClient returns a new ENF API client for the provided domain. If
// a nil httpClient is provided, a new http.Client will be used.  To
// use API methods which require authentication, provide an
// http.Client that will perform the authentication for you (such as
// that provided by TokenAuthClient in this library)
func NewClient(domain string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	if baseURL.Scheme == "" {
		baseURL.Scheme = defaultScheme
	}

	c := &Client{client: httpClient, Domain: domain, BaseURL: baseURL, UserAgent: defaultUserAgent}
	c.common.client = c
	c.Auth = (*AuthService)(&c.common)
	c.Firewall = (*FirewallService)(&c.common)
	c.Network = (*NetworkService)(&c.common)
	c.Domains = (*DomainService)(&c.common)
	c.Endpoint = (*EndpointService)(&c.common)
	//c.DNS = (*DNSService)(&c.common)
	return c, nil
}

// NewRequest creates a new HTTP request.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.APIToken != "" {
		req.Header.Set(headerToken, fmt.Sprintf(headerTokenFormat, c.APIToken))
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do executes an HTTP request.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
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

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

// ErrorResponse represents the error response from the API.
type ErrorResponse struct {
	Response *http.Response
	Errorr   struct {
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: [%d] %v - %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Errorr.Code, r.Errorr.Text)
}

// CheckResponse checks the HTTP response for an error.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 209 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		_ = json.Unmarshal(data, errorResponse)
	}

	return errorResponse
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
