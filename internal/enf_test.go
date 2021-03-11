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
// @since March 08, 2021
//
//-------------------------------------------------------------------------------------------
package internal

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/jarcoal/httpmock"
	_ "github.com/jarcoal/httpmock"
	"github.com/xaptum/go-enf/enf"
)

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// AssertError fails the test if err is nil
func assertError(tb testing.TB, err error) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: Expected error but got nil \033[39m\n\n", filepath.Base(file), line)
		tb.FailNow()
	}
}

// Equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

// mock httpClient
func mockClient(host string) (*enf.Client, error) {
	// create http client
	hc := &http.Client{}

	// activate httpmockk
	httpmock.ActivateNonDefault(hc)

	// now create enf client with the mocked http client
	return enf.NewWithClient(hc, host)
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestNew(t *testing.T) {
	client, err := enf.New()
	ok(t, err)
	equals(t, "https://api.xaptum.io", client.BaseUrl())
}

func TestNewWithHost(t *testing.T) {
	client, err := enf.New("http://localhost:9090")
	ok(t, err)
	equals(t, "http://localhost:9090", client.BaseUrl())
}

func TestNewInvalidHost(t *testing.T) {
	_, err := enf.New("hello", "world")
	assertError(t, err)
}
