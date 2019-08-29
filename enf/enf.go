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
	projectVersion = "0.1.0"

	defaultScheme = "https"

	headerToken       = "Authorization"
	headerTokenFormat = "Bearer %s"

	mediaTypeJson = "application/json"
)

var (
	defaultUserAgent = fmt.Sprintf("go-enf/%s (+%s; %s)", projectVersion, projectURL, runtime.Version())
)

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
	ApiToken string

	// Reuse a single struct instead of allocating one for each service on the heap
	common service

	// Services used for talking to different parts of the ENF API.
	Auth     *AuthService
	Firewall *FirewallService
}

type service struct {
	client *Client
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
	return c, nil
}

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

	if c.ApiToken != "" {
		req.Header.Set(headerToken, fmt.Sprintf(headerTokenFormat, c.ApiToken))
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
