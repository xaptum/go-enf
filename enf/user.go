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

// Invite represents an invite to a new user.
type Invite struct {
	ID           *int       `json:"id"`
	CreatedBy    *string    `json:"created_by"`
	InsertedDate *time.Time `json:"inserted_date"`
	ModifiedDate *time.Time `json:"modified_data"`
	Version      *int       `json:"version"`
	DomainID     *int       `json:"domain_id"`
	Email        *string    `json:"email"`
	InviteToken  *string    `json:"invite_token"`
	InvitedBy    *string    `json:"invited_by"`
	Name         *string    `json:"name"`
	Type         *string    `json:"type"`
}

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

// SendInviteRequest represents the body of the request to send an invite.
type SendInviteRequest struct {
	Email    *string `json:"email"`
	FullName *string `json:"full_name"`
	UserType *string `json:"user_type"`
}

// AcceptInviteRequest represents the body of the request to accept an invite.
type AcceptInviteRequest struct {
	Email    *string `json:"email"`
	Code     *string `json:"code"`
	Name     *string `json:"name"`
	Password *string `json:"password"`
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

type inviteResponse struct {
	Data []*Invite              `json:"data"`
	Page map[string]interface{} `json:"page"`
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

// ListInvitesForDomainAddress gets a list of active invites for a given domain address.
func (s *UserService) ListInvitesForDomainAddress(ctx context.Context, address string) ([]*Invite, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/invites", address)
	body, resp, err := s.client.get(ctx, path, url.Values{}, new(inviteResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*inviteResponse).Data, resp, nil
}

// SendNewInvite sends a new invite for a user to join the domain with the given address.
func (s *UserService) SendNewInvite(ctx context.Context, address string, inviteRequest *SendInviteRequest) (*Invite, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/invites", address)
	body, resp, err := s.client.post(ctx, path, new(inviteResponse), inviteRequest)
	if err != nil {
		return nil, resp, err
	}

	return body.(*inviteResponse).Data[0], resp, nil
}

// AcceptInvite accepts an invite.
func (s *UserService) AcceptInvite(ctx context.Context, acceptInviteRequest *AcceptInviteRequest) (*http.Response, error) {
	path := "api/xcr/v2/users/invites"
	_, resp, err := s.client.post(ctx, path, new(emptyResponse), acceptInviteRequest)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ResendInvite resends an invite to the given email address.
func (s *UserService) ResendInvite(ctx context.Context, email string) (*http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/invites/%v", email)
	_, resp, err := s.client.put(ctx, path, new(emptyResponse), struct{}{})
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteInvite deletes an invite to the given email address.
func (s *UserService) DeleteInvite(ctx context.Context, email string) (*http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/invites/%v", email)
	resp, err := s.client.delete(ctx, path)
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
