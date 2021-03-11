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
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Service
type Service struct {
	client *Client
}

// Client represents a wrapper for the HTTP client that communicates with the API.
type Client struct {
	// HTTP client used to communicate with the API.
	rst *resty.Client

	// Base URL
	baseUrl string

	// The API token for authenticating with the API
	authToken string

	// Reuse a single struct instead of allocating one for each service on the heap
	service Service

	// Services used for talking to different parts of the ENF API.
	AuthSvc     *AuthService
	DnsSvc      *DnsService
	DomainSvc   *DomainService
	EndpointSvc *EndpointService
	FirewallSvc *FirewallService
	NetworkSvc  *NetworkService
	UserSvc     *UserService
}

func (client *Client) BaseUrl() string {
	return client.baseUrl
}

func (client *Client) Get(ctx context.Context, path string, result interface{}) error {
	// call the api
	resp, err := client.rst.R().
		SetContext(ctx).
		SetAuthToken(client.authToken).
		Get(path)
	return client.processApiRespone(resp, err, result)
}

func (client *Client) Post(ctx context.Context, path string, request interface{}, result interface{}) error {
	// call the api
	resp, err := client.rst.R().
		SetContext(ctx).
		SetAuthToken(client.authToken).
		SetBody(request).
		Post(path)
	return client.processApiRespone(resp, err, result)
}

func (client *Client) Put(ctx context.Context, path string, request interface{}, result interface{}) error {
	// call the api
	resp, err := client.rst.R().
		SetContext(ctx).
		SetAuthToken(client.authToken).
		SetBody(request).
		Put(path)
	return client.processApiRespone(resp, err, result)
}

func (client *Client) Delete(ctx context.Context, path string, result interface{}) error {
	// call the api
	resp, err := client.rst.R().
		SetContext(ctx).
		SetAuthToken(client.authToken).
		Delete(path)
	return client.processApiRespone(resp, err, result)
}

func (client *Client) processApiRespone(resp *resty.Response, respErr error, result interface{}) error {
	// check for response error
	if nil != respErr {
		// wrap as api error
		msg := "Unable to create api request"
		apiErr := &EnfApiError{
			StatusCode:   0, // TODO: may be -1?
			ErrorMessage: &msg,
		}
		return apiErr
	}

	// handle response
	statusCode := resp.StatusCode()
	body := resp.Body()

	// create a place holder for api errors
	var apiErr = &EnfApiError{
		StatusCode: statusCode,
	}

	switch statusCode {
	case 200, 201:
		// request was successful.
		if err := json.Unmarshal(body, result); nil != err {
			// not a json response
			msg := "Invalid json response"
			apiErr.ErrorMessage = &msg
			return apiErr
		}
		// got back a json response
		return nil

	case 400, 401:
		// bad request
		if err := json.Unmarshal(body, apiErr); nil != err {
			// not a json error message
			msg := string(body)
			apiErr.ErrorMessage = &msg
		}

	case 403:
		// method not found error
		apiErr.SetCodeText("http_error", "Method Not Found")

	case 415:
		// media type error
		apiErr.SetCodeText("http_error", "Unsupported Media Type")

	case 404:
		// parse response body for error json
		if err := json.Unmarshal(body, apiErr); nil != err {
			// not a json error message. assume URL not found
			apiErr.SetCodeText("http_error", "Not Found")
		}

	case 500:
		// parse response body for error json
		if err := json.Unmarshal(body, apiErr); nil != err {
			// not a json error message. encode a genric error message
			apiErr.SetCodeText("server_error", "Server error received without details")
		}

	default:
		msg := fmt.Sprintf("unexpected status code %d", statusCode)
		apiErr.ErrorMessage = &msg
	}

	// if we reached here that means there was an api error
	return apiErr
}
