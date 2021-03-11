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
package internal

import (
	"context"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/xaptum/go-enf/enf"
)

func TestMe(t *testing.T) {
	// create the client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// initialize httpmock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"id":66,"created_by":"root@xaptum","inserted_date":"2018-04-04T11:15:03Z","modified_date":"2020-08-22T18:00:00Z","description":"","full_name":"Xaptum Administrators","last_login":"2021-03-09T23:33:58Z","status":"ACTIVE","username":"xap@admin","domain_id":1,"domain":"fd00:8f80:8000::/48","roles":[{"cidr":"fd00:8f80:8000::/48","role":"DOMAIN_ADMIN"}]}],"page":{"curr":"","next":"","prev":""}}`
	httpmock.RegisterResponder("GET", "http://localhost"+enf.XcrApiPath("/me"),
		httpmock.NewStringResponder(200, fixture))

	// call the api
	user, err := client.UserSvc.Me(context.Background())
	ok(t, err)
	equals(t, "xap@admin", *user.Username)
	equals(t, enf.ActiveStatus, *user.Status)
}
