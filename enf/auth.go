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

// AuthResponse represents the authentication credentials returned by
// the auth API.
type AuthResponse struct {
	Username      *string `json:"username"`
	Token         *string `json:"token"`
	UserID        *int64  `json:"user_id"`
	UserType      *string `json:"type"`
	DomainID      *int64  `json:"domain_id"`
	DomainNetwork *string `json:"domain_network"`
}

func (s *AuthService) Authenticate(ctx context.Context, authReq *AuthRequest) (*AuthResponse, *http.Response, error) {
	if *authReq.Username == "" {
		return nil, nil, ErrMissingUsername
	}

	if *authReq.Password == "" {
		return nil, nil, ErrMissingPassword
	}

	path := "/api/xcr/v2/xauth"
	req, err := s.client.NewRequest("POST", path, authReq)
	if err != nil {
		return nil, nil, err
	}

	body := &struct {
		Data []AuthResponse         `json:"data"`
		Page map[string]interface{} `json:"page"`
	}{}
	resp, err := s.client.Do(ctx, req, body)
	if err != nil {
		return nil, resp, err
	}

	return &body.Data[0], resp, nil
}
