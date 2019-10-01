package enf

import (
	"context"
	"fmt"
	"net/http"
)

// EndpointService handles communication with all endpoint related methods within
// the ENF API.
type EndpointService service

// EndpointRateLimits represents the values of rate limits for an endpoint.
type EndpointRateLimits struct {
	PacketsPerSecond *int  `json:"packets_per_second"`
	PacketsBurstSize *int  `json:"packets_burst_size"`
	BytesPerSecond   *int  `json:"bytes_per_second"`
	BytesBurstSize   *int  `json:"bytes_burst_size"`
	Inherit          *bool `json:"inherit"`
}

// endpointRateLimitResponse represents the standard API response for
// all relevant endpoint rate limit endpoints.
type endpointRateLimitResponse struct {
	Data []*EndpointRateLimits  `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// GetCurrentRateLimits gets the current rate limits for the given endpoint IPv6 address.
func (s *EndpointService) GetCurrentRateLimits(ctx context.Context, endpointIPv6 string) (*EndpointRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/cxns/%v/ep_rate_limits/current", endpointIPv6)
	body, resp, err := s.client.get(ctx, path, new(endpointRateLimitResponse))
	if err != nil {
		return nil, resp, err
	}
	return body.(*endpointRateLimitResponse).Data[0], resp, nil
}

// GetMaxRateLimits gets the max rate limits for the given endpoint IPv6 address.
func (s *EndpointService) GetMaxRateLimits(ctx context.Context, endpointIPv6 string) (*EndpointRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/cxns/%v/ep_rate_limits/max", endpointIPv6)
	body, resp, err := s.client.get(ctx, path, new(endpointRateLimitResponse))
	if err != nil {
		return nil, resp, err
	}
	return body.(*endpointRateLimitResponse).Data[0], resp, nil
}

// SetCurrentRateLimits sets the current rate limits for the given endpoint IPv6 address and the specified limit.
func (s *EndpointService) SetCurrentRateLimits(ctx context.Context, values *EndpointRateLimits, endpointIPv6 string) (*EndpointRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/cxns/%v/ep_rate_limits/current", endpointIPv6)
	body, resp, err := s.client.put(ctx, path, new(endpointRateLimitResponse), values)
	if err != nil {
		return nil, resp, err
	}
	return body.(*endpointRateLimitResponse).Data[0], resp, nil
}

// SetMaxRateLimits sets the max rate limits for the given endpoint IPv6 address and the specified limit.
func (s *EndpointService) SetMaxRateLimits(ctx context.Context, values *EndpointRateLimits, endpointIPv6 string) (*EndpointRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/cxns/%v/ep_rate_limits/max", endpointIPv6)
	body, resp, err := s.client.put(ctx, path, new(endpointRateLimitResponse), values)
	if err != nil {
		return nil, resp, err
	}
	return body.(*endpointRateLimitResponse).Data[0], resp, nil
}
