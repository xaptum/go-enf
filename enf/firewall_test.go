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

func TestFirewallService_ListRules(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xfw/v1/N/rule", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		fmt.Fprint(w, `[{"id":"00000000-0000-4000-2000-000000000000"},
                        {"id":"00000000-0000-4000-2000-000000000001"}]`)
	})

	rules, _, err := client.Firewall.ListRules(context.Background(), "N")
	if err != nil {
		t.Errorf("Firewall.ListRules returned error: %v", err)
	}

	want := []*FirewallRule{
		{ID: String("00000000-0000-4000-2000-000000000000")},
		{ID: String("00000000-0000-4000-2000-000000000001")},
	}
	if !reflect.DeepEqual(rules, want) {
		t.Errorf("Firewall.ListRules returned %+v, want %+v", rules, want)
	}
}

func TestFirewallService_GetRule(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xfw/v1/N/rule", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		fmt.Fprint(w, `[{"id":"00000000-0000-4000-2000-000000000000"},
                        {"id":"00000000-0000-4000-2000-000000000001"}]`)
	})

	rule, _, err := client.Firewall.GetRule(context.Background(), "N", "00000000-0000-4000-2000-000000000001")
	if err != nil {
		t.Errorf("Firewall.GetRule returned error: %v", err)
	}

	want := &FirewallRule{
		ID: String("00000000-0000-4000-2000-000000000001"),
	}
	if !reflect.DeepEqual(rule, want) {
		t.Errorf("Firewall.GetRule returned %+v, want %+v", rule, want)
	}
}

func TestFirewallService_CreateRule(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	input := &FirewallRuleRequest{
		Priority:  Int(1),
		Action:    String("ACCEPT"),
		Direction: String("INGRESSS"),
	}

	wantAcceptHeaders := []string{mediaTypeJson}
	wantContentTypeHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xfw/v1/N/rule", func(w http.ResponseWriter, r *http.Request) {
		v := new(FirewallRuleRequest)
		_ = json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		testHeader(t, r, "Content-Type", strings.Join(wantContentTypeHeaders, ", "))
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"ID":"00000000-0000-4000-2000-000000000002"}`)
	})

	rule, _, err := client.Firewall.CreateRule(context.Background(), "N", input)
	if err != nil {
		t.Errorf("Firewall.CreateRule returned error: %v", err)
	}

	want := &FirewallRule{
		ID: String("00000000-0000-4000-2000-000000000002"),
	}
	if !reflect.DeepEqual(rule, want) {
		t.Errorf("Firewall.CreateRule returned %+v, want %+v", rule, want)
	}
}

func TestFirewallService_DeleteRule(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/xfw/v1/N/rule/00000000-0000-4000-2000-000000000000", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Firewall.DeleteRule(context.Background(), "N", "00000000-0000-4000-2000-000000000000")
	if err != nil {
		t.Errorf("FirewallRule.DeleteLabel returned error: %v", err)
	}
}
