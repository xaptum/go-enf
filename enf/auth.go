package enf

import (
	"context"
	"errors"
)

var (
	ErrMissingHost     = errors.New("Missing required parameter 'host'")
	ErrMissingUsername = errors.New("Missing required parameter 'username'")
	ErrMissingPassword = errors.New("Missing required paramter 'password'")
)

// AuthRequest represents a request to authenticate with the
// API.
type AuthRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

// Credentials represents the authentication credentials returned by
// the auth API.
type Credentials struct {
	Username *string     `json:"username"`
	Token    *string     `json:"token"`
	UserID   *int64      `json:"user_id"`
	Roles    []*UserRole `json:"roles"`
	DomainID *int64      `json:"domain_id"`
	Domain   *string     `json:"domain"`
}

type AuthResponse struct {
	Data []Credentials          `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// Authenticate authenticates the given authorization request.
func Authenticate(ctx context.Context, host, username, password string) (*Credentials, error) {
	if "" == host {
		return nil, ErrMissingHost
	}

	if username == "" {
		return nil, ErrMissingUsername
	}

	if password == "" {
		return nil, ErrMissingPassword
	}

	// create api client
	http, err := NewClient(host, nil, nil)
	if nil != err {
		return nil, err
	}

	// create request object
	req := NewRequest("/api/xcr/v3/xauth", nil, nil, &AuthRequest{Username: String(username), Password: String(password)})

	authResponse := new(AuthResponse)
	_, err := http.post(ctx, req, authResponse)
	if nil != err {
		return nil, err
	}

	var cred Credentials = authResponse.Data[0]
	return &cred, nil
}
