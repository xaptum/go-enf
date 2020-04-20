package enf

import (
	"context"
	"net/http"
	"testing"
)

func TestUserService_ListUserRoles(t *testing.T) {
	path := "/api/xcr/v3/users/1/roles"

	responseBodyMock := `{
		"data": [
			{
				"cidr": "D/d0",
				"role": "DOMAIN_ADMIN"
			},
			{
				"cidr": "N/n0",
				"role": "NETWORK_USER"
			}
		]
	}`

	expected := []*UserRole{
		{
			CIDR: String("D/d0"),
			Role: String("DOMAIN_ADMIN"),
		},
		{
			CIDR: String("N/n0"),
			Role: String("NETWORK_USER"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.ListUserRoles(context.Background(), 1)
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

func TestUserService_AppendUserRoles(t *testing.T) {
	path := "/api/xcr/v3/users/1/roles"

	requestBody := []UserRole{
		{String("N/n0"), String("NETWORK_USER")},
	}

	responseBodyMock := `{
		"data": [
			{
				"cidr": "D/d0",
				"role": "DOMAIN_ADMIN"
			},
			{
				"cidr": "N/n0",
				"role": "NETWORK_USER"
			}
		]
	}`

	expected := []*UserRole{
		{
			CIDR: String("D/d0"),
			Role: String("DOMAIN_ADMIN"),
		},
		{
			CIDR: String("N/n0"),
			Role: String("NETWORK_USER"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.AppendUserRoles(context.Background(), 1, requestBody)
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      requestBody,
		ResponseBodyMock: responseBodyMock,
		Expected:         expected,
		Method:           method,
		T:                t,
	}

	postTest(testParams)
}

func TestUserService_ReplaceUserRoles(t *testing.T) {
	path := "/api/xcr/v3/users/1/roles"

	requestBody := []UserRole{
		{String("D/d0"), String("DOMAIN_ADMIN")},
		{String("N/n0"), String("NETWORK_USER")},
	}

	responseBodyMock := `{
		"data": [
			{
				"cidr": "D/d0",
				"role": "DOMAIN_ADMIN"
			},
			{
				"cidr": "N/n0",
				"role": "NETWORK_USER"
			}
		]
	}`

	expected := []*UserRole{
		{
			CIDR: String("D/d0"),
			Role: String("DOMAIN_ADMIN"),
		},
		{
			CIDR: String("N/n0"),
			Role: String("NETWORK_USER"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.ReplaceUserRoles(context.Background(), 1, requestBody)
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      requestBody,
		ResponseBodyMock: responseBodyMock,
		Expected:         expected,
		Method:           method,
		T:                t,
	}

	putTest(testParams)
}

func TestUserService_DeleteUserRoles(t *testing.T) {
	path := "/api/xcr/v3/users/1/roles"

	responseBodyMock := `[]`

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.User.DeleteUserRoles(context.Background(), 1)
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

	deleteTest(testParams)
}

func TestUserService_DeleteUserRolesWithQuery(t *testing.T) {
	path := "/api/xcr/v3/users/1/roles"

	responseBodyMock := `[]`

	query := DeleteUserRolesQuery{[]string{"DOMAIN_*", "*_ADMIN"}, nil}

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.User.DeleteUserRolesWithQuery(context.Background(), 1, query)
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

	deleteTest(testParams)
}
