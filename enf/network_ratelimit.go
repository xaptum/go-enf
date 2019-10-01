package enf

import (
	"context"
	"fmt"
	"net/http"
)

// NetworkRateLimits represents the values of rate limits for a network.
type NetworkRateLimits struct {
	PacketsPerSecond *int  `json:"packets_per_second"`
	PacketsBurstSize *int  `json:"packets_burst_size"`
	BytesPerSecond   *int  `json:"bytes_per_second"`
	BytesBurstSize   *int  `json:"bytes_burst_size"`
	Inherit          *bool `json:"inherit"`
}

// standardRateLimitResponse represents the standard API response for
// all relevant network rate limit endpoints.
type networkRateLimitResponse struct {
	Data []*NetworkRateLimits   `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// GetDefaultEndpointRateLimits gets the default rate limits for endpoints in the given network.
func (s *NetworkService) GetDefaultEndpointRateLimits(ctx context.Context, network string) (*NetworkRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/nws/%v/ep_rate_limits/default", network)
	body, resp, err := s.client.get(ctx, path, new(networkRateLimitResponse))
	if err != nil {
		return nil, resp, err
	}
	return body.(*networkRateLimitResponse).Data[0], resp, nil
}

// GetMaxDefaultEndpointRateLimits gets the max default rate limits for endpoints in the given network.
func (s *NetworkService) GetMaxDefaultEndpointRateLimits(ctx context.Context, network string) (*NetworkRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/nws/%v/ep_rate_limits/max", network)
	body, resp, err := s.client.get(ctx, path, new(networkRateLimitResponse))
	if err != nil {
		return nil, resp, err
	}
	return body.(*networkRateLimitResponse).Data[0], resp, nil
}

// SetDefaultEndpointRateLimits sets the default rate limits for endpoints in the given network.
func (s *NetworkService) SetDefaultEndpointRateLimits(ctx context.Context, values *NetworkRateLimits, network string) (*NetworkRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/nws/%v/ep_rate_limits/default", network)
	body, resp, err := s.client.put(ctx, path, new(networkRateLimitResponse), values)
	if err != nil {
		return nil, resp, err
	}
	return body.(*networkRateLimitResponse).Data[0], resp, nil
}

// SetMaxDefaultEndpointRateLimits sets the max default rate limits for endpoints in the given network.
func (s *NetworkService) SetMaxDefaultEndpointRateLimits(ctx context.Context, values *NetworkRateLimits, network string) (*NetworkRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/nws/%v/ep_rate_limits/max", network)
	body, resp, err := s.client.put(ctx, path, new(networkRateLimitResponse), values)
	if err != nil {
		return nil, resp, err
	}
	return body.(*networkRateLimitResponse).Data[0], resp, nil
}
