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

func TestNetworkService_GetDomainRateLimit(t *testing.T) {
	defaultPath := "/api/xcr/v2/domains/N/ep_rate_limits/default"

	defaultResponseBodyMock := `{
		"data": [
			{
				"packets_per_second": 1000,
				"packets_burst_size": 1000,
				"bytes_per_second": 10000,
				"bytes_burst_size": 10000
			}
		],
		"page": {
			"curr": -1,
			"next": -1,
			"prev": -1
		}
	}`

	defaultExpected :=
		&DomainRateLimits{
			PacketsPerSecond: Int(1000),
			PacketsBurstSize: Int(1000),
			BytesPerSecond:   Int(10000),
			BytesBurstSize:   Int(10000),
		}

	defaultMethod := func(client *Client) (interface{}, *http.Response, error) {
		return client.Domains.GetDefaultEndpointRateLimits(context.Background(), "N")
	}

	defaultTestingParameters := &TestParams{
		Path:             defaultPath,
		RequestBody:      struct{}{},
		ResponseBodyMock: defaultResponseBodyMock,
		Expected:         defaultExpected,
		Method:           defaultMethod,
		T:                t,
	}

	getTest(defaultTestingParameters)

	maxPath := "/api/xcr/v2/domains/N/ep_rate_limits/max"

	maxResponseBodyMock := `{
		"data": [
			{
				"packets_per_second": 1000,
				"packets_burst_size": 1000,
				"bytes_per_second": 10000,
				"bytes_burst_size": 10000
			}
		],
		"page": {
			"curr": -1,
			"next": -1,
			"prev": -1
		}
	}`

	maxExpected :=
		&DomainRateLimits{
			PacketsPerSecond: Int(1000),
			PacketsBurstSize: Int(1000),
			BytesPerSecond:   Int(10000),
			BytesBurstSize:   Int(10000),
		}

	maxMethod := func(client *Client) (interface{}, *http.Response, error) {
		return client.Domains.GetMaxDefaultEndpointRateLimits(context.Background(), "N")
	}

	maxTestParams := &TestParams{
		Path:             maxPath,
		RequestBody:      struct{}{},
		ResponseBodyMock: maxResponseBodyMock,
		Expected:         maxExpected,
		Method:           maxMethod,
		T:                t,
	}

	getTest(maxTestParams)
}

func TestNetworkService_SetDomainRateLimit(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	// expected body for PUT endpoint
	input := &DomainRateLimits{
		PacketsPerSecond: Int(100),
		PacketsBurstSize: Int(100),
		BytesPerSecond:   Int(10000),
		BytesBurstSize:   Int(10000),
	}

	mux.HandleFunc("/api/xcr/v2/domains/N/ep_rate_limits/default", func(w http.ResponseWriter, r *http.Request) {
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
					"bytes_burst_size": 10000
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

	newDefaultRateLimits, _, err := client.Domains.SetDefaultEndpointRateLimits(context.Background(), input, "N")
	if err != nil {
		t.Errorf("Domains.SetDefaultRateLimits returned error: %v", err)
	}

	want := &DomainRateLimits{
		PacketsPerSecond: Int(100),
		PacketsBurstSize: Int(100),
		BytesPerSecond:   Int(10000),
		BytesBurstSize:   Int(10000),
	}

	if !reflect.DeepEqual(newDefaultRateLimits, want) {
		t.Errorf("Domains.SetDefaultRateLimits returned %+v, want %+v", newDefaultRateLimits, want)
	}

	// Now, let's test going over the max rate limit for PacketsPerSecond.
	illegalInput := &DomainRateLimits{
		PacketsPerSecond: Int(10000),
		PacketsBurstSize: Int(100),
		BytesPerSecond:   Int(10000),
		BytesBurstSize:   Int(10000),
	}

	_, _, expectedErr := client.Domains.SetDefaultEndpointRateLimits(context.Background(), illegalInput, "N")
	// We expect a 400 error here, so if the error is nil we'll fail the test.
	if expectedErr == nil {
		t.Errorf("Domains.SetDefaultRateLimits should have returned error by exceeding the rate limit max: %v", err)
	}

}
