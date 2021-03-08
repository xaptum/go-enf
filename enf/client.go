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
	Auth *AuthService
	/*DNS      *DNSService
	Domain   *DomainService
	Endpoint *EndpointService
	Firewall *FirewallService
	Network  *NetworkService
	User     *UserService*/
}

func (client *Client) Get(ctx context.Context, path string, result interface{}) error {
	// call the api
	resp, err := client.rst.R().
		SetContext(ctx).
		Get(path)
	return client.processApiRespone(resp, err, result)
}

func (client *Client) Post(ctx context.Context, path string, request interface{}, result interface{}) error {
	// call the api
	resp, err := client.rst.R().
		SetContext(ctx).
		SetBody(request).
		Post(path)
	return client.processApiRespone(resp, err, result)
}

func (client *Client) Put(ctx context.Context, path string, request interface{}, result interface{}) error {
	// call the api
	resp, err := client.rst.R().
		SetContext(ctx).
		SetBody(request).
		Put(path)
	return client.processApiRespone(resp, err, result)
}

func (client *Client) Delete(ctx context.Context, path string, result interface{}) error {
	// call the api
	resp, err := client.rst.R().
		SetContext(ctx).
		Get(path)
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
		apiErr.CodeError.Code = "http_error"
		apiErr.CodeError.Text = "Method Not Found"

	case 415:
		// media type error
		apiErr.CodeError.Code = "http_error"
		apiErr.CodeError.Text = "Unsupported Media Type"

	case 404:
		// parse response body for error json
		if err := json.Unmarshal(body, apiErr); nil != err {
			// not a json error message. assume URL not found
			apiErr.CodeError.Code = "http_error"
			apiErr.CodeError.Text = "Not Found"
		}

	case 500:
		// parse response body for error json
		if err := json.Unmarshal(body, apiErr); nil != err {
			// not a json error message. encode a genric error message
			apiErr.CodeError.Code = "server_error"
			apiErr.CodeError.Text = "Server error received without details"
		}

	default:
		msg := fmt.Sprintf("unexpected status code %d", statusCode)
		apiErr.ErrorMessage = &msg
	}

	// if we reached here that means there was an api error
	return apiErr
}
