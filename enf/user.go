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
type UserService service

// User represents an ENF user.
type User struct {
	UserID      *int       `json:"user_id"`
	Username    *string    `json:"username"`
	Description *string    `json:"description"`
	FullName    *string    `json:"full_name"`
	LastLogin   *time.Time `json:"last_login"`
	DomainID    *int       `json:"domain_id"`
	Type        *string    `json:"type"`
	ResetCode   *string    `json:"reset_code"`
	ResetTime   *time.Time `json:"reset_time"`
	Status      *string    `json:"status"`
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

// ListUsersForDomainAddress gets the list of users for a given domain address.
func (s *UserService) ListUsersForDomainAddress(ctx context.Context, address string) ([]*User, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/users", address)
	body, resp, err := s.client.get(ctx, path, url.Values{}, new(userResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*userResponse).Data, resp, nil
}

// ListUsersForDomainID gets a list of users for a given unique domain identifier.
func (s *UserService) ListUsersForDomainID(ctx context.Context, id string) ([]*User, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/users", id)
	body, resp, err := s.client.get(ctx, path, url.Values{}, new(userResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*userResponse).Data, resp, nil
}

// UpdateUserStatus updates the status of a user to "ACTIVE" or "INACTIVE".
func (s *UserService) UpdateUserStatus(ctx context.Context, userID int, updateUserStatusRequest *UpdateUserStatusRequest) (*http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/users/%v/status", userID)
	_, resp, err := s.client.put(ctx, path, new(emptyResponse), updateUserStatusRequest)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// EmailResetPasswordCode emails a reset password code to the given email address.
func (s *UserService) EmailResetPasswordCode(ctx context.Context, email string) (*http.Response, error) {
	path := "api/xcr/v2/users/reset"

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
	path := "api/xcr/v2/users/reset"
	_, resp, err := s.client.post(ctx, path, new(emptyResponse), resetPasswordRequest)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
