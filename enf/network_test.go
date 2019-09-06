package enf

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestNetworkService_ListNetworks(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xcr/v2/domains/N/nws", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		fmt.Fprint(w, `{
		"data": [
			{
				"name": "TestNetwork 1",
				"network": "fd00:8f80:8000:0000::/64",
				"description": "This is a network.",
				"status": "ACTIVE"
			},
			{
				"name": "TestNetwork 2",
				"network": "fd00:8f80:8000:1::/64",
				"description": "This is another network.",
				"status": "ACTIVE"
			}
		],
		"page": 
		{
			"curr": -1,
			"next": -1,
			"prev": -1
		}
	}
			`)
	})

	networks, _, err := client.Network.ListNetworks(context.Background(), "N")
	if err != nil {
		t.Errorf("Network.ListNetworks returned error: %v", err)
	}

	want := []*Network{
		{
			Name:        String("TestNetwork 1"),
			Network:     String("fd00:8f80:8000:0000::/64"),
			Description: String("This is a network."),
			Status:      String("ACTIVE"),
		},
		{
			Name:        String("TestNetwork 2"),
			Network:     String("fd00:8f80:8000:1::/64"),
			Description: String("This is another network."),
			Status:      String("ACTIVE"),
		},
	}
	if !reflect.DeepEqual(networks, want) {
		t.Errorf("Network.ListNetworks returned %+v, want %+v", networks, want)
	}
}
