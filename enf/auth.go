package enf

import (
	"context"
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

type authResponse struct {
	Data []Credentials          `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// Authenticate authenticates the given authorization request.
func (svc *AuthService) Authenticate(ctx context.Context, username, password string) (*Credentials, error) {
	// create result struct
	result := &authResponse{}

	// call the authentication api
	err := svc.client.Post(ctx, "/xauth/v1/authenticate",
		AuthRequest{
			Username: username,
			Password: password,
		}, result)

	// Check if request failed
	if nil != err {
		return nil, err
	}

	// update the client with auth token
	credentials := &result.Data[0]
	svc.client.authToken = credentials.Token

	// return credentials
	return credentials, nil
}
