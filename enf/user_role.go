package enf

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// UserRole represents an ENF user role.
type UserRole struct {
	CIDR *string `json:"cidr"`
	Role *string `json:"role"`
}

// DeleteUserRolesQuery represents the query parameters to delete some
// of a user's roles.
type DeleteUserRolesQuery struct {
	roles   []string
	network *string
}

type userRoleResponse struct {
	Data []*UserRole            `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// ListUserRoles get the list of roles for the given user.
func (s *UserService) ListUserRoles(ctx context.Context, userID int) ([]*UserRole, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v3/users/%v/roles", userID)
	body, resp, err := s.client.get(ctx, path, url.Values{}, new(userRoleResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*userRoleResponse).Data, resp, nil
}

// AppendUserRoles adds roles to the user
func (s *UserService) AppendUserRoles(ctx context.Context, userID int, roles []UserRole) ([]*UserRole, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v3/users/%v/roles", userID)
	body, resp, err := s.client.post(ctx, path, new(userRoleResponse), roles)
	if err != nil {
		return nil, resp, err
	}

	return body.(*userRoleResponse).Data, resp, nil
}

// ReplaceUserRoles replaces all roles for the user
func (s *UserService) ReplaceUserRoles(ctx context.Context, userID int, roles []UserRole) ([]*UserRole, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v3/users/%v/roles", userID)
	body, resp, err := s.client.put(ctx, path, new(userRoleResponse), roles)
	if err != nil {
		return nil, resp, err
	}

	return body.(*userRoleResponse).Data, resp, nil
}

// DeleteUserRolesQuery deletes the roles for given user that match the specified query
func (s *UserService) DeleteUserRolesWithQuery(ctx context.Context, userID int, query DeleteUserRolesQuery) (*http.Response, error) {
	path := fmt.Sprintf("api/xcr/v3/users/%v/roles", userID)

	queryParameters := url.Values{}
	queryParameters.Add("roles", strings.Join(query.roles, ","))
	if query.network != nil {
		queryParameters.Add("network_cidr", *query.network)
	}

	resp, err := s.client.delete(ctx, path, queryParameters)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteUserRoles deletes all roles from the given user
func (s *UserService) DeleteUserRoles(ctx context.Context, userID int) (*http.Response, error) {
	path := fmt.Sprintf("api/xcr/v3/users/%v/roles", userID)

	resp, err := s.client.delete(ctx, path, url.Values{})
	if err != nil {
		return resp, err
	}

	return resp, nil
}
