package enf

import (
	"context"
	"net/http"
	"testing"
)

func TestUserService_ListUsers(t *testing.T) {
	path := "/api/xcr/v3/users"

	responseBodyMock := `{
		"data": [
			{
				"user_id": 1,
				"username": "user@acme",
				"full_name": "Xaptum User",
				"roles": [
					{
						"role" : "DOMAIN_USER",
						"cidr" : "D/d0"
					}
				],
				"status": "ACTIVE"
			}
		]
	}`

	expected := []*User{
		{
			UserID:   Int(1),
			Username: String("user@acme"),
			FullName: String("Xaptum User"),
			Roles: []*UserRole{
				{
					Role: String("DOMAIN_USER"),
					CIDR: String("D/d0"),
				},
			},
			Status: String("ACTIVE"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.ListUsers(context.Background())
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

func TestUserService_ListUsersForDomain(t *testing.T) {
	path := "/api/xcr/v3/users"

	responseBodyMock := `{
		"data": [
			{
				"user_id": 1,
				"username": "user@acme",
				"full_name": "Xaptum User",
				"roles": [
					{
						"role" : "DOMAIN_USER",
						"cidr" : "D/d0"
					}
				],
				"status": "ACTIVE"
			}
		]
	}`

	expected := []*User{
		{
			UserID:   Int(1),
			Username: String("user@acme"),
			FullName: String("Xaptum User"),
			Roles: []*UserRole{
				{
					Role: String("DOMAIN_USER"),
					CIDR: String("D/d0"),
				},
			},
			Status: String("ACTIVE"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.ListUsersForDomain(context.Background(), "D/d0")
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

func TestUserService_ListUsersForNetwork(t *testing.T) {
	path := "/api/xcr/v3/users"

	responseBodyMock := `{
		"data": [
			{
				"user_id": 1,
				"username": "user@acme",
				"full_name": "Xaptum User",
				"roles": [
					{
						"role" : "NETWORK_ADMIN",
						"cidr" : "N/n0"
					}
				],
				"status": "ACTIVE"
			}
		]
	}`

	expected := []*User{
		{
			UserID:   Int(1),
			Username: String("user@acme"),
			FullName: String("Xaptum User"),
			Roles: []*UserRole{
				{
					Role: String("NETWORK_ADMIN"),
					CIDR: String("N/n0"),
				},
			},
			Status: String("ACTIVE"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.ListUsersForNetwork(context.Background(), "N/n0")
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

func TestUserService_GetUser(t *testing.T) {
	path := "/api/xcr/v3/users/1"

	responseBodyMock := `{
		"data": [
			{
				"user_id": 1,
				"username": "user@acme",
				"full_name": "Xaptum User",
				"roles": [
					{
						"role" : "DOMAIN_USER",
						"cidr" : "D/d0"
					},
					{
						"role" : "NETWORK_ADMIN",
						"cidr" : "N/n0"
					}
				],
				"status": "ACTIVE"
			}
		]
	}`

	expected := &User{
		UserID:   Int(1),
		Username: String("user@acme"),
		FullName: String("Xaptum User"),
		Roles: []*UserRole{
			{
				Role: String("DOMAIN_USER"),
				CIDR: String("D/d0"),
			},
			{
				Role: String("NETWORK_ADMIN"),
				CIDR: String("N/n0"),
			},
		},
		Status: String("ACTIVE"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.GetUser(context.Background(), 1)
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
	path := "/api/xcr/v3/users/61/status"

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
	path := "/api/xcr/v3/users/reset"

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
	path := "/api/xcr/v3/users/reset"

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
