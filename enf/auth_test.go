package enf

import (
	"context"
	"net/http"
	"testing"
)

func TestAuthService_Authenticate(t *testing.T) {
	path := "/api/xcr/v3/xauth"

	requestBody := &AuthRequest{
		Username: String("user"),
		Password: String("pass"),
	}

	responseBodyMock := `{
		"data": [
			{
				"username":"user",
				"token":"12345678",
				"user_id":1,
				"roles": [
					{
						"cidr" : "N/n0",
						"role" : "NETWORK_USER"
					}
				]
			}
		],
		"page": {
		}
		}`

	expected := &Credentials{
		Username: String("user"),
		Token:    String("12345678"),
		UserID:   Int64(1),
		Roles: []*UserRole{
			{
				CIDR: String("N/n0"),
				Role: String("NETWORK_USER"),
			},
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Auth.Authenticate(context.Background(), requestBody)
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
