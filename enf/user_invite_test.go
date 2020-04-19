package enf

import (
	"context"
	"net/http"
	"testing"
)

func TestUserService_ListInvitesForDomainAddress(t *testing.T) {
	path := "/api/xcr/v2/domains/1/invites"

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
		return client.User.ListInvitesForDomainAddress(context.Background(), "1")
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

func TestUserService_SendNewInvite(t *testing.T) {
	path := "/api/xcr/v2/domains/1/invites"

	requestBody := &SendInviteRequest{
		Email:    String("user@acme.com"),
		FullName: String("Xaptum User"),
		UserType: String("DOMAIN_USER"),
	}

	responseBodyMock := `{
		"data": [
			{
				"id": 1,
				"email": "user@acme.com",
				"name": "Xaptum User",
				"type": "DOMAIN_USER"
			}
		]
	}`

	expected := &Invite{
		ID:    Int(1),
		Email: String("user@acme.com"),
		Name:  String("Xaptum User"),
		Type:  String("DOMAIN_USER"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.User.SendNewInvite(context.Background(), "1", requestBody)
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
	path := "/api/xcr/v2/users/invites"

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
	path := "/api/xcr/v2/invites/user@acme.com"

	responseBodyMock := `[]`

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.User.ResendInvite(context.Background(), "user@acme.com")
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
	path := "/api/xcr/v2/invites/user@acme.com"

	responseBodyMock := `[]`

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.User.DeleteInvite(context.Background(), "user@acme.com")
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
