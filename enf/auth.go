package enf

import (
	"context"
	"errors"
	"net/http"
)

var (
	ErrMissingUsername = errors.New("Missing required field 'Username'")
	ErrMissingPassword = errors.New("Missing required field 'Password'")
)

// AuthService handles communication with authentication related
// methods of the ENF API. These methods are used to obtain
// authentication tokens.
type AuthService service

// AuthRequest represents a request to authenticate with the
// API.
type AuthRequest struct {
	Username *string `json:"username"`
	Password *string `json:"token"`
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

type authResponse struct {
	Data []Credentials          `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// Authenticate authenticates the given authorization request.
func (s *AuthService) Authenticate(ctx context.Context, authReq *AuthRequest) (*Credentials, *http.Response, error) {
	if *authReq.Username == "" {
		return nil, nil, ErrMissingUsername
	}

	if *authReq.Password == "" {
		return nil, nil, ErrMissingPassword
	}

	endpoint := "/api/xcr/v3/xauth"

	body, resp, err := s.client.post(ctx, endpoint, new(authResponse), authReq)
	if err != nil {
		return nil, resp, err
	}
	return &body.(*authResponse).Data[0], resp, nil
}
