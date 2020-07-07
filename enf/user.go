package enf

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// UserService handles communication with the user related methods of the
// ENF API. These methods include sending, resending, accepting, deleting, and
// listing new user invites, viewing the users for a domain, and resetting user passwords.
type UserService Service

// User represents an ENF user.
type User struct {
	UserID      *int        `json:"user_id"`
	Username    *string     `json:"username"`
	Description *string     `json:"description"`
	FullName    *string     `json:"full_name"`
	LastLogin   *time.Time  `json:"last_login"`
	Domain      *string     `json:"domain"`
	DomainID    *int        `json:"domain_id"`
	Roles       []*UserRole `json:"roles"`
	ResetCode   *string     `json:"reset_code"`
	ResetTime   *time.Time  `json:"reset_time"`
	Status      *string     `json:"status"`
}

// UpdateUserStatusRequest represents the body of the request to update a user's status.
type UpdateUserStatusRequest struct {
	Status *string `json:"status"`
}

// ResetPasswordRequest represents the body of the request to reset a user's password.
type ResetPasswordRequest struct {
	Email    *string `json:"email"`
	Code     *string `json:"code"`
	Password *string `json:"pwd"`
}

type userResponse struct {
	Data []*User                `json:"data"`
	Page map[string]interface{} `json:"page"`
}

type emptyResponse []interface{}

// ListUsers gets the list of all users.
func (s *UserService) ListUsers(ctx context.Context) ([]*User, *http.Response, error) {
	path := "api/xcr/v3/users"

	body, resp, err := s.client.get(ctx, path, url.Values{}, new(userResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*userResponse).Data, resp, nil
}

// ListUsersForDomain gets the list of users with roles in a given domain.
func (s *UserService) ListUsersForDomain(ctx context.Context, domain string) ([]*User, *http.Response, error) {
	path := "api/xcr/v3/users"

	queryParameters := url.Values{}
	queryParameters.Add("domain", domain)

	body, resp, err := s.client.get(ctx, path, queryParameters, new(userResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*userResponse).Data, resp, nil
}

// ListUsersForNetwork gets the list of users with roles in a given network.
func (s *UserService) ListUsersForNetwork(ctx context.Context, network string) ([]*User, *http.Response, error) {
	path := "api/xcr/v3/users"

	queryParameters := url.Values{}
	queryParameters.Add("network", network)

	body, resp, err := s.client.get(ctx, path, queryParameters, new(userResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*userResponse).Data, resp, nil
}

func (s *UserService) GetUser(ctx context.Context, userID int) (*User, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v3/users/%v", userID)
	body, resp, err := s.client.get(ctx, path, url.Values{}, new(userResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*userResponse).Data[0], resp, nil
}

// UpdateUserStatus updates the status of a user to "ACTIVE" or "INACTIVE".
func (s *UserService) UpdateUserStatus(ctx context.Context, userID int, updateUserStatusRequest *UpdateUserStatusRequest) (*http.Response, error) {
	path := fmt.Sprintf("api/xcr/v3/users/%v/status", userID)
	_, resp, err := s.client.put(ctx, path, new(emptyResponse), updateUserStatusRequest)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// EmailResetPasswordCode emails a reset password code to the given email address.
func (s *UserService) EmailResetPasswordCode(ctx context.Context, email string) (*http.Response, error) {
	path := "api/xcr/v3/users/reset"

	queryParameters := url.Values{}
	queryParameters.Add("email", email)
	_, resp, err := s.client.get(ctx, path, queryParameters, new(emptyResponse))
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ResetPassword resets a user's password.
func (s *UserService) ResetPassword(ctx context.Context, resetPasswordRequest *ResetPasswordRequest) (*http.Response, error) {
	path := "api/xcr/v3/users/reset"
	_, resp, err := s.client.post(ctx, path, new(emptyResponse), resetPasswordRequest)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
