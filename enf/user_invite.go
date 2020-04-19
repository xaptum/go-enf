package enf

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

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

type inviteResponse struct {
	Data []*Invite              `json:"data"`
	Page map[string]interface{} `json:"page"`
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
