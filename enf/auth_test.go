package enf

import (
	"context"
	"net/http"
	"testing"
)

func TestAuthService_Authenticate(t *testing.T) {
	path := "/api/xcr/v2/xauth"

	requestBody := &AuthRequest{
		Username: String("user"),
		Password: String("pass"),
	}

	responseBodyMock := `{
		"data": [
			{
				"username":"user",
				"token":"12345678",
				"user_id":1
			}
		],
		"page": {

			}
		}`

	expected := &Credentials{
		Username: String("user"),
		Token:    String("12345678"),
		UserID:   Int64(1),
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
