//-------------------------------------------------------------------------------------------
//
// XAPTUM CONFIDENTIAL
// __________________
//
//  2021(C) Xaptum, Inc.
//  All Rights Reserved.Patents Pending.
//
// NOTICE:  All information contained herein is, and remains
// the property of Xaptum, Inc.  The intellectual and technical concepts contained
// herein are proprietary to Xaptum, Inc and may be covered by U.S. and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Xaptum, Inc.
//
// @author Venkatakumar Srinivasan
// @since March 09, 2021
//
//-------------------------------------------------------------------------------------------
package enf

import (
	"context"
	"time"
)

// UserService handles communication with the user related methods of the
// ENF API. These methods include sending, resending, accepting, deleting, and
// listing new user invites, viewing the users for a domain, and resetting user passwords.
type UserService Service

type UserRole struct {
	Cidr *string `json:"cidr"`
	Role *string `json:"role"`
}

type User struct {
	Id          *int        `json:"id"`
	Description *string     `json:"description"`
	FullName    *string     `json:"full_name"`
	LastLogin   *time.Time  `json:"last_login"`
	Status      *string     `json:"status"`
	Username    *string     `json:"username"`
	DomainId    *int        `json:"domain_id"`
	Domain      *string     `json:"domain"`
	Roles       []*UserRole `json:"roles"`
}

type Invite struct {
	Id          *int        `json:"id"`
	Email       *string     `json:"email"`
	InviteToken *string     `json:"invite_token"`
	Name        *string     `json:"name"`
	DomainId    *int        `json:"domain_id"`
	Domain      *string     `json:"domain"`
	Roles       []*UserRole `json:"roles"`
}

type userResponse struct {
	Data []*User                `json:"data"`
	Page map[string]interface{} `json:"page"`
}

/*type inviteResponse struct {
	Data []*Invite               `json:"data"`
	Page map[string]interface{} `json:"page"`
}

type userRolesResponse struct {
	Data []*UserRole             `json:"data"`
	Page map[string]interface{} `json:"page"`
    }*/

// Me
func (svc *UserService) Me(ctx context.Context) (*User, error) {
	// create result struct
	result := &userResponse{}

	// call the me api
	err := svc.client.Get(ctx, XcrApiPath("/me"), result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	// return user
	return result.Data[0], nil
}
