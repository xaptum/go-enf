package enf

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	ProjectURL     = "github.com/xaptum/go-enf"
	ProjectVersion = "0.4.0"

	HttpsScheme = "https"
	HttpScheme  = "http"

	HeaderToken       = "Authorization"
	HeaderTokenFormat = "Bearer %s"

	MediaTypeJSON = "application/json"

	DefaultHost = "https://api.xaptum.io"
)

var (
	defaultUserAgent = fmt.Sprintf("go-enf/%s (+%s; %s)", ProjectVersion, ProjectURL, runtime.Version())
)

// Client represents a wrapper for the HTTP client that communicates with the API.
type Client struct {
	// HTTP client used to communicate with the API.
	httpClient *resty.Client

	// Base URL
	baseUrl string

	// The API token for authenticating with the API
	authToken string

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

// NewClient returns a new ENF API client for the provided domain. If
// a nil httpClient is provided, a new http.Client will be used.  To
// use API methods which require authentication, provide an
// http.Client that will perform the authentication for you (such as
// that provided by TokenAuthClient in this library)
func New(host ...string) (*Client, error) {
	// create a http client with timeouts
	httpClient := &http.Client{
		Timeout: time.Second * 20,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
		},
	}

	// return new enf client
	return NewWithClient(httpClient, host...)
}

func NewWithClient(httpClient *http.Client, host ...string) (*Client, error) {
	// validate there is atmost on host value
	if len(host) > 1 {
		return nil, fmt.Errorf("Invalid host parameter")
	}

	// set baseurl
	var hostStr string
	if 0 == len(host) {
		hostStr = DefaultHost
	} else {
		hostStr = host[0]
	}
	baseUrl, err := url.Parse(hostStr)

	if nil != err {
		return nil, fmt.Errorf("Invalid host parameter: ")
	}

	// check if the scheme is http or https
	if HttpScheme != baseUrl.Scheme && HttpsScheme != baseUrl.Scheme {
		baseUrl.Scheme = HttpsScheme
	}

	// create a resty client
	restyClient := resty.NewWithClient(httpClient)
	restyClient.SetHostURL(baseUrl.String())

	// set headers for all requests
	restyClient.SetHeaders(map[string]string{
		"Accept":       MediaTypeJSON,
		"Content-Type": MediaTypeJSON,
		"User-Agent":   defaultUserAgent,
	})

	// create enf api client
	client := &Client{
		httpClient: restyClient,
		baseUrl:    baseUrl.String(),
	}
	client.service.client = client
	/*	c.Domains = (*DomainService)(&c.service)
		c.Endpoint = (*EndpointService)(&c.service)
		c.DNS = (*DNSService)(&c.service)
		c.Firewall = (*FirewallService)(&c.service)
		c.Network = (*NetworkService)(&c.service)
		c.User = (*UserService)(&c.service)*/
	return client, nil
}

// All the exported methods in this file are designed to be general-purpose HTTP helpers. These methods
// will accept any request struct, and support any struct type you want the response to be stored in.
// For usage examples, see the methods in network.go or firewall.go
