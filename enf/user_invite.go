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
	Domain       *string    `json:"domain"`
	Role         *UserRole  `json:"role"`
	Email        *string    `json:"email"`
	InviteToken  *string    `json:"invite_token"`
	InvitedBy    *string    `json:"invited_by"`
	Name         *string    `json:"name"`
}

// SendInviteRequest represents the body of the request to send an invite.
type SendInviteRequest struct {
	Email    *string     `json:"email"`
	FullName *string     `json:"full_name"`
	Domain   *string     `json:"domain"`
	Roles    []*UserRole `json:"roles"`
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

// ListInvites gets a list of active invites.
func (s *UserService) ListInvites(ctx context.Context) ([]*Invite, *http.Response, error) {
	path := "api/xcr/v3/invites"
	body, resp, err := s.client.get(ctx, path, url.Values{}, new(inviteResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*inviteResponse).Data, resp, nil
}

// SendInvite sends an invite for a new user.
func (s *UserService) SendInvite(ctx context.Context, inviteRequest *SendInviteRequest) (*Invite, *http.Response, error) {
	path := "api/xcr/v3/invites"
	body, resp, err := s.client.post(ctx, path, new(inviteResponse), inviteRequest)
	if err != nil {
		return nil, resp, err
	}

	return body.(*inviteResponse).Data[0], resp, nil
}

// AcceptInvite accepts an invite.
func (s *UserService) AcceptInvite(ctx context.Context, acceptInviteRequest *AcceptInviteRequest) (*http.Response, error) {
	path := "api/xcr/v3/invites"
	_, resp, err := s.client.post(ctx, path, new(emptyResponse), acceptInviteRequest)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ResendInvite resends the given invite.
func (s *UserService) ResendInvite(ctx context.Context, id int) (*http.Response, error) {
	path := fmt.Sprintf("api/xcr/v3/invites/%v", id)
	_, resp, err := s.client.put(ctx, path, new(emptyResponse), struct{}{})
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteInvite deletes the given invite
func (s *UserService) DeleteInvite(ctx context.Context, id int) (*http.Response, error) {
	path := fmt.Sprintf("api/xcr/v3/invites/%v", id)
	resp, err := s.client.delete(ctx, path, url.Values{})
	if err != nil {
		return resp, err
	}

	return resp, nil
}
