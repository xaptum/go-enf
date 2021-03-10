//-------------------------------------------------------------------------------------------
//
// XAPTUM CONFIDENTIAL
// __________________
//
//  2021(C) Xaptum, Inc.
//  All Rights Reserved.Patents Pending.
//
// NOTICE:  All information contained herein is, and remains
// the property of Xaptum, Inc.  The intellectual and technical concepts contained
// herein are proprietary to Xaptum, Inc and may be covered by U.S. and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Xaptum, Inc.
//
// @author Venkatakumar Srinivasan
// @since March 08, 2021
//
//-------------------------------------------------------------------------------------------
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

	XcrBasePath      = "/api/xcr/v3"
	FirewallBasePath = "/api/xfw/v2"
	DnsBasePath      = "/api/xdns/v1"
	IamBasePath      = "/api/xiam/v1"
	CaptiveBasePath  = "/api/captive/v1"
)

var (
	ActiveStatus   = "ACTIVE"
	InactiveStatus = "INACTIVE"

	XaptumAdmin  = "XAPTUM_ADMIN"
	CaptiveAdmin = "CAPTIVE_ADMIN"
	IamAdmin     = "IAM_ADMIN"
	DomainAdmin  = "DOMAIN_ADMIN"
	DomainUser   = "DOMAIN_USER"
	NetworkAdmin = "NETWORK_ADMIN"
	NetworkUser  = "NETWORK_USER"

	DnsAAAA  = "AAAA"
	DnsCNAME = "CNAME"
	DnsTXT   = "TXT"
	DnsSRV   = "SRV"
)

var (
	defaultUserAgent = fmt.Sprintf("go-enf/%s (+%s; %s)", ProjectVersion, ProjectURL, runtime.Version())
)

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
	rst := resty.NewWithClient(httpClient)
	rst.SetHostURL(baseUrl.String())

	// set headers for all requests
	rst.SetHeaders(map[string]string{
		"Accept":       MediaTypeJSON,
		"Content-Type": MediaTypeJSON,
		"User-Agent":   defaultUserAgent,
	})

	// create enf api client
	client := &Client{
		rst:       rst,
		baseUrl:   baseUrl.String(),
		authToken: "", // initialize to empty string
	}
	client.service.client = client
	client.AuthSvc = (*AuthService)(&client.service)
	/*	client.DomainSvc = (*DomainService)(&client.service)
		client.EndpointSvc = (*EndpointService)(&client.service)
		client.DnsSvc = (*DnsService)(&client.service)
		client.FirewallSvc = (*FirewallService)(&client.service)
		client.NetworkSvc = (*NetworkService)(&client.service)*/
	client.UserSvc = (*UserService)(&client.service)
	return client, nil
}

func xcrApiPath(path string) string {
	return XcrBasePath + path
}

/*func firewallApiPath(path string) string {
	return FirewallBasePath + path
}

func dnsApiPath(path string) string {
	return DnsBasePath + path
}

func iamApiPath(path string) string {
	return IamBasePath + path
}

func captiveApiPath(path string) string {
	return CaptiveBasePath + path
    }*/

// All the exported methods in this file are designed to be general-purpose HTTP helpers. These methods
// will accept any request struct, and support any struct type you want the response to be stored in.
// For usage examples, see the methods in network.go or firewall.go
