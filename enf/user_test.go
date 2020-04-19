package enf

import (
	"context"
	"net/http"
	"testing"
)

func TestUserService_ListUsersForDomainAddress(t *testing.T) {
	path := "/api/xcr/v2/domains/N/users"

	responseBodyMock := `{
		"data": [
			{
				"user_id": 1,
				"username": "user@acme",
				"full_name": "Xaptum User",
				"status": "ACTIVE"
			}
		]
	}`

	expected := []*User{
		{
			UserID:   Int(1),
			Username: String("user@acme"),
			FullName: String("Xaptum User"),
			Status:   String("ACTIVE"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.ListUsersForDomainAddress(context.Background(), "N")
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      struct{}{},
		ResponseBodyMock: responseBodyMock,
		Expected:         expected,
		Method:           method,
		T:                t,
	}

	getTest(testParams)
}

func TestUserService_ListUsersForDomainID(t *testing.T) {
	path := "/api/xcr/v2/domains/1/users"

	responseBodyMock := `{
		"data": [
			{
				"user_id": 1,
				"username": "user@acme",
				"full_name": "Xaptum User",
				"status": "ACTIVE"
			}
		]
	}`

	expected := []*User{
		{
			UserID:   Int(1),
			Username: String("user@acme"),
			FullName: String("Xaptum User"),
			Status:   String("ACTIVE"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.ListUsersForDomainID(context.Background(), "1")
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      struct{}{},
		ResponseBodyMock: responseBodyMock,
		Expected:         expected,
		Method:           method,
		T:                t,
	}

	getTest(testParams)
}

func TestUserService_UpdateUserStatus(t *testing.T) {
	path := "/api/xcr/v2/users/61/status"

	requestBody := &UpdateUserStatusRequest{
		Status: String("INACTIVE"),
	}

	responseBodyMock := `[]`

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.User.UpdateUserStatus(context.Background(), 61, requestBody)
		return struct{}{}, resp, err
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      requestBody,
		ResponseBodyMock: responseBodyMock,
		Expected:         struct{}{},
		Method:           method,
		T:                t,
	}

	putTest(testParams)
}

func TestUserService_EmailResetPasswordCode(t *testing.T) {
	path := "/api/xcr/v2/users/reset"

	responseBodyMock := `[]`

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.User.EmailResetPasswordCode(context.Background(), "user@acme.com")
		return struct{}{}, resp, err
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      struct{}{},
		ResponseBodyMock: responseBodyMock,
		Expected:         struct{}{},
		Method:           method,
		T:                t,
	}

	getTest(testParams)
}

func TestUserService_ResetPassword(t *testing.T) {
	path := "/api/xcr/v2/users/reset"

	requestBody := &ResetPasswordRequest{
		Email:    String("user@acme.com"),
		Code:     String("token1234"),
		Password: String("User1234!"),
	}

	responseBodyMock := `[]`

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.User.ResetPassword(context.Background(), requestBody)
		return struct{}{}, resp, err
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      requestBody,
		ResponseBodyMock: responseBodyMock,
		Expected:         struct{}{},
		Method:           method,
		T:                t,
	}

	postTest(testParams)
}
