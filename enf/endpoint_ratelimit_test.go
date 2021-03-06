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

func TestNetworkService_GetEndpointRateLimit(t *testing.T) {
	currentPath := "/api/xcr/v2/cxns/N/n/1234:5678/ep_rate_limits/current"

	currentResponseBodyMock := `{
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
	expected := &EndpointRateLimits{
		PacketsPerSecond: Int(100000),
		PacketsBurstSize: Int(100000),
		BytesPerSecond:   Int(10000000),
		BytesBurstSize:   Int(10000000),
		Inherit:          Bool(true),
	}
	currentMethod := func(client *Client) (interface{}, *http.Response, error) {
		return client.Endpoint.GetCurrentRateLimits(context.Background(), "N/n/1234:5678")
	}

	currentTestParams := &TestParams{
		Path:             currentPath,
		RequestBody:      struct{}{},
		ResponseBodyMock: currentResponseBodyMock,
		Expected:         expected,
		Method:           currentMethod,
		T:                t,
	}

	getTest(currentTestParams)

	maxPath := "/api/xcr/v2/cxns/N/n/1234:5678/ep_rate_limits/max"

	maxResponseBodyMock := `{
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

	maxMethod := func(client *Client) (interface{}, *http.Response, error) {
		return client.Endpoint.GetMaxRateLimits(context.Background(), "N/n/1234:5678")
	}

	maxTestParams := &TestParams{
		Path:             maxPath,
		RequestBody:      struct{}{},
		ResponseBodyMock: maxResponseBodyMock,
		Expected:         expected,
		Method:           maxMethod,
		T:                t,
	}

	getTest(maxTestParams)
}

func TestNetworkService_SetEndpointRateLimit(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	// expected body for PUT endpoint
	input := &EndpointRateLimits{
		PacketsPerSecond: Int(100),
		PacketsBurstSize: Int(100),
		BytesPerSecond:   Int(10000),
		BytesBurstSize:   Int(10000),
		Inherit:          Bool(false),
	}

	mux.HandleFunc("/api/xcr/v2/cxns/N/n/1234:5678/ep_rate_limits/current", func(w http.ResponseWriter, r *http.Request) {
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

	newCurrentRateLimit, _, err := client.Endpoint.SetCurrentRateLimits(context.Background(), input, "N/n/1234:5678")
	if err != nil {
		t.Errorf("Endpoint.SetCurrentRateLimits returned error: %v", err)
	}

	want := &EndpointRateLimits{
		PacketsPerSecond: Int(100),
		PacketsBurstSize: Int(100),
		BytesPerSecond:   Int(10000),
		BytesBurstSize:   Int(10000),
		Inherit:          Bool(false),
	}

	if !reflect.DeepEqual(newCurrentRateLimit, want) {
		t.Errorf("Endpoint.SetCurrentRateLimits returned %+v, want %+v", newCurrentRateLimit, want)
	}

	// Now, let's test going over the max rate limit for PacketsPerSecond.
	illegalInput := &EndpointRateLimits{
		PacketsPerSecond: Int(10000),
		PacketsBurstSize: Int(100),
		BytesPerSecond:   Int(10000),
		BytesBurstSize:   Int(10000),
		Inherit:          Bool(false),
	}

	_, _, expectedErr := client.Endpoint.SetCurrentRateLimits(context.Background(), illegalInput, "N/n/1234")
	// We expect a 400 error here, so if the error is nil we'll fail the test.
	if expectedErr == nil {
		t.Errorf("Endpoint.SetCurrentRateLimits should have returned error by exceeding the rate limit max: %v", err)
	}

}
