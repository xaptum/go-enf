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

func TestNetworkService_GetNetworkRateLimit(t *testing.T) {
	path := "/api/xcr/v2/nws/N/n/ep_rate_limits/default"

	responseBodyMock := `{
		"data": [
			{
				"packets_per_second": 100000,
				"packets_burst_size": 100000,
				"bytes_per_second": 10000000,
				"bytes_burst_size": 10000000,
				"inherit": true 
			}
		],
		"page": {
			"curr": -1,
			"next": -1,
			"prev": -1
		}
	}
		`

	expected := &NetworkRateLimits{
		PacketsPerSecond: Int(100000),
		PacketsBurstSize: Int(100000),
		BytesPerSecond:   Int(10000000),
		BytesBurstSize:   Int(10000000),
		Inherit:          Bool(true),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Network.GetDefaultEndpointRateLimits(context.Background(), "N/n")
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      struct{}{},
		ResponseBodyMock: responseBodyMock,
		Expected:         expected,
		Method:           method,
		T:                t,
	}

	getTest(testParams)
}

func TestNetworkService_SetNetworkRateLimit(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	// expected body for PUT endpoint
	input := &NetworkRateLimits{
		PacketsPerSecond: Int(100),
		PacketsBurstSize: Int(100),
		BytesPerSecond:   Int(10000),
		BytesBurstSize:   Int(10000),
		Inherit:          Bool(false),
	}

	mux.HandleFunc("/api/xcr/v2/nws/N/n/ep_rate_limits/default", func(w http.ResponseWriter, r *http.Request) {
		v := new(DomainRateLimits)
		_ = json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		testHeader(t, r, "Content-Type", strings.Join(wantContentTypeHeaders, ", "))

		if *v.PacketsPerSecond < 1000 {
			w.WriteHeader(200)
			fmt.Fprint(w, `{
			"data": [
				{
					"packets_per_second": 100,
					"packets_burst_size": 100,
					"bytes_per_second": 10000,
					"bytes_burst_size": 10000,
					"inherit": false
				}
			],
			"page": {
				"curr": -1,
				"next": -1,
				"prev": -1
			}
		}
			`)
		} else {
			w.WriteHeader(400)
			fmt.Fprint(w, `{
			"error": {
				"code": "validation_error",
				"text": "rate limit exceeds allowed max"
			}
		}`)
		}
	})

	newDefaultRateLimit, _, err := client.Network.SetDefaultEndpointRateLimits(context.Background(), input, "N/n")
	if err != nil {
		t.Errorf("Network.SetDefaultRateLimit returned error: %v", err)
	}

	want := &NetworkRateLimits{
		PacketsPerSecond: Int(100),
		PacketsBurstSize: Int(100),
		BytesPerSecond:   Int(10000),
		BytesBurstSize:   Int(10000),
		Inherit:          Bool(false),
	}

	if !reflect.DeepEqual(newDefaultRateLimit, want) {
		t.Errorf("Network.SetDefaultRateLimit returned %+v, want %+v", newDefaultRateLimit, want)
	}

	// Now, let's test going over the max rate limit for PacketsPerSecond.
	illegalInput := &NetworkRateLimits{
		PacketsPerSecond: Int(10000),
		PacketsBurstSize: Int(100),
		BytesPerSecond:   Int(10000),
		BytesBurstSize:   Int(10000),
		Inherit:          Bool(false),
	}

	_, _, expectedErr := client.Network.SetDefaultEndpointRateLimits(context.Background(), illegalInput, "N/n")
	// We expect a 400 error here, so if the error is nil we'll fail the test.
	if expectedErr == nil {
		t.Errorf("Network.SetDefaultRateLimit should have returned error by exceeding the rate limit max: %v", err)
	}

}
