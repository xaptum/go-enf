package enf

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestAuthService_Authenticate(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	input := &AuthRequest{
		Username: String("user"),
		Password: String("pass"),
	}

	wantAcceptHeaders := []string{mediaTypeJson}
	wantContentTypeHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xcr/v2/xauth", func(w http.ResponseWriter, r *http.Request) {
		v := new(AuthRequest)
		_ = json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		testHeader(t, r, "Content-Type", strings.Join(wantContentTypeHeaders, ", "))
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":[{"username":"user",
                                "token":"12345678",
                                "user_id":1}],
                        "page":{}}`)
	})

	auth, _, err := client.Auth.Authenticate(context.Background(), input)
	if err != nil {
		t.Errorf("Auth.Authenticate returned error: %v", err)
	}

	want := &AuthResponse{
		Username: String("user"),
		Token:    String("12345678"),
		UserID:   Int64(1),
	}
	if !reflect.DeepEqual(auth, want) {
		t.Errorf("Auth.Authenticate returned %+v, want %+v", auth, want)
	}
}
