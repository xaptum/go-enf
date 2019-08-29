package enf

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
