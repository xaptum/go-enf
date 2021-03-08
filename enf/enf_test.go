package enf

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// assert fails the test if the condition is false.
/*func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
    }*/

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

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestNew(t *testing.T) {
	client, err := New()
	ok(t, err)
	equals(t, "https://api.xaptum.io", client.baseUrl)
	equals(t, "", client.authToken)
}

func TestNewWithHost(t *testing.T) {
	client, err := New("http://localhost:9090")
	ok(t, err)
	equals(t, "http://localhost:9090", client.baseUrl)
	equals(t, "", client.authToken)
}

func TestNewInvalidHost(t *testing.T) {
	_, err := New("hello", "world")
	assertError(t, err)
}
