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
)

func TestAuthenticateSuccess(t *testing.T) {
	// create client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate httpmock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"token":"authtoken","user_id":1,"username":"xap@admin"}],"page":{"curr":"","next":"","prev":""}}`
	httpmock.RegisterResponder("POST", "http://localhost/xauth/v1/authenticate",
		httpmock.NewStringResponder(200, fixture))

	// call the api
	credentials, err := client.AuthSvc.Authenticate(context.Background(), "xap@admin", "xxxxx")
	ok(t, err)
	equals(t, "xap@admin", *credentials.Username)
	assert(t, "" != *credentials.Token, "Check Token")
}

func TestAuthenticateFail(t *testing.T) {
	// create client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactive httpmock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"error":{"code":"authentication_failed","text":"Authentication Failed"}}`
	httpmock.RegisterResponder("POST", "http://localhost/xauth/v1/authenticate",
		httpmock.NewStringResponder(400, fixture))

	// call the api
	credentials, err := client.AuthSvc.Authenticate(context.Background(), "xap@admin", "xxxxx")
	assert(t, nil == credentials, "Verify credentials is nil")
	equals(t, "AUTHENTICATION_FAILED: Authentication Failed", err.Error())
}
