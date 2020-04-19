package enf

import (
	"context"
	"net/http"
	"testing"
)

func TestUserService_ListInvites(t *testing.T) {
	path := "/api/xcr/v3/invites"

	responseBodyMock := `{
		"data": [
			{
				"id": 1,
				"email": "user@acme.com",
				"name": "Xaptum User",
				"invite_token": "token1234"
			}
		]
	}`

	expected := []*Invite{
		{
			ID:          Int(1),
			Email:       String("user@acme.com"),
			Name:        String("Xaptum User"),
			InviteToken: String("token1234"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.ListInvites(context.Background())
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

func TestUserService_SendInvite(t *testing.T) {
	path := "/api/xcr/v3/invites"

	requestBody := &SendInviteRequest{
		Email:    String("user@acme.com"),
		FullName: String("Xaptum User"),
		Roles: []*UserRole{
			{
				CIDR: String("D/d0"),
				Role: String("DOMAIN_USER"),
			},
		},
	}

	responseBodyMock := `{
		"data": [
			{
				"id": 1,
				"email": "user@acme.com",
				"name": "Xaptum User",
				"role": {
					"cidr" : "D/d0",
					"role" : "DOMAIN_USER"
				}
			}
		]
	}`

	expected := &Invite{
		ID:    Int(1),
		Email: String("user@acme.com"),
		Name:  String("Xaptum User"),
		Role: &UserRole{
			CIDR: String("D/d0"),
			Role: String("DOMAIN_USER"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.SendInvite(context.Background(), requestBody)
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

func TestUserService_AcceptInvite(t *testing.T) {
	path := "/api/xcr/v3/invites"

	requestBody := &AcceptInviteRequest{
		Email:    String("user@acme.com"),
		Code:     String("token1234"),
		Name:     String("Xaptum User"),
		Password: String("User1234!"),
	}
	responseBodyMock := `[]`

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.User.AcceptInvite(context.Background(), requestBody)
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

func TestUserService_ResendInvite(t *testing.T) {
	path := "/api/xcr/v3/invites/1"

	responseBodyMock := `[]`

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.User.ResendInvite(context.Background(), 1)
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

	putTest(testParams)
}

func TestUserService_DeleteInvite(t *testing.T) {
	path := "/api/xcr/v3/invites/1"

	responseBodyMock := `[]`

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.User.DeleteInvite(context.Background(), 1)
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
