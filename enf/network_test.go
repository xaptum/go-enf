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

func TestNetworkService_GetNetwork(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xcr/v2/nws/N/n", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		fmt.Fprint(w, `{
				"data": [
				  {
					"name": "TestNetwork 1",
					"network": "N/n",
					"description": "This is a network.",
					"status": "ACTIVE"
				  }
				],
				"page": {
				  "curr": -1,
				  "next": -1,
				  "prev": -1
				}
			  }
			`)
	})
	network, _, err := client.Network.GetNetwork(context.Background(), "N/n")
	if err != nil {
		t.Errorf("Network.GetNetwork returned error: %v", err)
	}

	want := &Network{
		Name:        String("TestNetwork 1"),
		Network:     String("N/n"),
		Description: String("This is a network."),
		Status:      String("ACTIVE"),
	}
	if !reflect.DeepEqual(network, want) {
		t.Errorf("Network.GetNetwork returned %+v, want %+v", network, want)
	}
}

func TestNetworkService_CreateNetwork(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	input := &NetworkRequest{
		Name:        String("TestNetwork 1"),
		Description: String("This is a new network."),
	}

	wantAcceptHeaders := []string{mediaTypeJson}
	wantContentTypeHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xcr/v2/domains/1/nws", func(w http.ResponseWriter, r *http.Request) {
		v := new(NetworkRequest)
		_ = json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		testHeader(t, r, "Content-Type", strings.Join(wantContentTypeHeaders, ", "))
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{
			"data": [
			  {
				"name": "TestNetwork 1",
				"network": "N/n",
				"description": "This is a network.",
				"status": "ACTIVE"
				}
			],
			"page": {
			  "curr": -1,
			  "next": -1,
			  "prev": -1
			}
		  }
			`)
	})

	network, _, err := client.Network.CreateNetwork(context.Background(), "1", input)
	if err != nil {
		t.Errorf("Network.GetNetwork returned error: %v", err)
	}

	want := &Network{
		Name:        String("TestNetwork 1"),
		Network:     String("N/n"),
		Description: String("This is a network."),
		Status:      String("ACTIVE"),
	}
	if !reflect.DeepEqual(network, want) {
		t.Errorf("Network.GetNetwork returned %+v, want %+v", network, want)
	}
}

func TestNetworkService_UpdateNetwork(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{mediaTypeJson}
	wantContentTypeHeaders := []string{mediaTypeJson}

	// mock out creating the network
	mux.HandleFunc("/api/xcr/v2/domains/1/nws", func(w http.ResponseWriter, r *http.Request) {
		v := new(NetworkRequest)
		_ = json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		testHeader(t, r, "Content-Type", strings.Join(wantContentTypeHeaders, ", "))

		fmt.Fprint(w, `{
			"data": [
			  {
				"name": "TestNetwork 1",
				"network": "N/n",
				"description": "This is a network.",
				"status": "ACTIVE"
				}
			],
			"page": {
			  "curr": -1,
			  "next": -1,
			  "prev": -1
			}
		  }
			`)
	})

	// expected body for PUT endpoint
	input := &NetworkRequest{
		Name:        String("TestNetwork 334"),
		Description: String("Trying to update the network.."),
	}

	mux.HandleFunc("/api/xcr/v2/nws/N/n", func(w http.ResponseWriter, r *http.Request) {
		v := new(NetworkRequest)
		_ = json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ","))
		testHeader(t, r, "Content-Type", strings.Join(wantContentTypeHeaders, ", "))
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprintf(w, `{
			"data": [
				{
					"name": "TestNetwork 334",
					"network": "N/n",
					"description":  "Trying to update the network..",
					"status": "ACTIVE"
				}
			],
			"page": {
				"curr": -1,
				"next": -1,
				"prev": -1
			}
		}
	`)
	})

	networkReq := &NetworkRequest{
		Name:        String("TestNetwork 1"),
		Description: String("This is a network."),
	}

	// first, "create" a new network
	network, _, _ := client.Network.CreateNetwork(context.Background(), "1", networkReq)

	fields := &NetworkRequest{
		Name:        String("TestNetwork 334"),
		Description: String("Trying to update the network.."),
	}

	// now, update the network we just created
	updatedNetwork, _, err := client.Network.UpdateNetwork(context.Background(), *network.Network, fields)
	if err != nil {
		t.Errorf("Network.UpdateNetwork returned error: %v", err)
	}

	want := &Network{
		Name:        String("TestNetwork 334"),
		Network:     String("N/n"),
		Description: String("Trying to update the network.."),
		Status:      String("ACTIVE"),
	}
	if !reflect.DeepEqual(updatedNetwork, want) {
		t.Errorf("Network.UpdateNetwork returned %+v, want %+v", updatedNetwork, want)
	}

}
