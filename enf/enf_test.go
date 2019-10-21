package enf

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type TestParams struct {

	// Path is the API path.
	Path string

	// RequestBody is the body of the request. For GET/DELETE requests,
	// RequestBody should be an empty struct.
	RequestBody interface{}

	// ResponseBodyMock is the body of the response from the mock server
	// for this test.
	ResponseBodyMock string

	// Expected is the expected result from calling Method.
	Expected interface{}

	// Method is the method being tested.
	Method func(*Client) (interface{}, *http.Response, error)

	// T is the testing type used for managing test state and logging.
	T *testing.T
}

// setup sets up a test HTTP server along with an enf.Client that is
// configured to talk to that test server. Test should register
// handlers on mux which provides mock responses for the API method
// being tested.
func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)
	client, _ = NewClient(server.URL, nil)

	return client, mux, server.Close
}

// Each of the following four functions provides an easy, abstracted
// way to test an API method. For examples of using these functions,
// look at any of the testing files.
func getTest(testingParameters *TestParams) interface{} {
	return checkMethod("GET", testingParameters)
}

func postTest(testingParameters *TestParams) interface{} {
	return checkMethod("POST", testingParameters)
}

func putTest(testingParameters *TestParams) interface{} {
	return checkMethod("PUT", testingParameters)
}

func deleteTest(testingParameters *TestParams) interface{} {
	return checkMethod("DELETE", testingParameters)
}

// checkMethod does the hard work for verifying that the method being
// tested gives the right output. checkMethod sets up the dependencies,
// mocks out the path, and checks for equality with the expected and actual responses.
func checkMethod(methodType string, testingParameters *TestParams) interface{} {
	client, mux, teardown := setup()
	defer teardown()

	mockPath(testingParameters, methodType, mux)

	result, _, err := testingParameters.Method(client)
	if err != nil {
		testingParameters.T.Error(err)
	}

	if !reflect.DeepEqual(result, testingParameters.Expected) {
		testingParameters.T.Errorf("Method returned %+v, want %+v", result, testingParameters.Expected)
	}

	return result
}

// mockPath creates a mock response on a mux for a specific path included in the testing
// parameters.
func mockPath(testingParameters *TestParams, methodType string, mux *http.ServeMux) {
	mux.HandleFunc(testingParameters.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(testingParameters.T, r, methodType)
		testHeaders(testingParameters.T, methodType, mux, r)
		fmt.Fprint(w, testingParameters.ResponseBodyMock)
	})
}

// testHeaders verifies that the headers are appropriate for the given request.
func testHeaders(t *testing.T, methodType string, mux *http.ServeMux, r *http.Request) {
	switch methodType {
	case "GET":
		testGetHeaders(t, r)
	case "POST":
		testPostPutHeaders(t, r)
	case "PUT":
		testPostPutHeaders(t, r)
	}
}

func testGetHeaders(t *testing.T, r *http.Request) {
	testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
}

func testPostPutHeaders(t *testing.T, r *http.Request) {
	testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
	testHeader(t, r, "Content-Type", strings.Join(wantContentTypeHeaders, ", "))
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}
