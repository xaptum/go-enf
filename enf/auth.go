package enf

import (
	"context"
	"fmt"
)

// AuthService handles communication with authentication related
// methods of the ENF API. These methods are used to obtain
// authentication tokens.
type AuthService Service

// AuthRequest represents a request to authenticate with the
// API.
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Credentials represents the authentication credentials returned by
// the auth API.
type Credentials struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	UserID   int64  `json:"user_id"`
}

/*type authResponse struct {
	Data []Credentials          `json:"data"`
	Page map[string]interface{} `json:"page"`
    }*/

// Authenticate authenticates the given authorization request.
func (svc *AuthService) Authenticate(ctx context.Context, username, password string) (*Credentials, error) {
	// call the authentication api
	resp, err := svc.client.httpClient.R().
		SetContext(ctx).
		SetBody(AuthRequest{
			Username: username,
			Password: password,
		}).
		Post("/api/xauth/v1/authenticate")

	if err != nil {
		return nil, err
	}

	fmt.Println("Resp: " + string(resp.Body()))
	//	fmt.Println("Resp Status: " + string(resp.StatusCode()))
	str := resp.String()
	return &Credentials{Username: str}, nil
}
