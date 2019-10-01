package enf

import (
	"context"
	"fmt"
	"net/http"
)

// DomainService handles communication with the domain-related methods
// of the ENF API.
type DomainService service

// DomainRateLimits represents the values of rate limits for a domain.
type DomainRateLimits struct {
	PacketsPerSecond *int `json:"packets_per_second"`
	PacketsBurstSize *int `json:"packets_burst_size"`
	BytesPerSecond   *int `json:"bytes_per_second"`
	BytesBurstSize   *int `json:"bytes_burst_size"`
}

// domainRateLimitResponse represents the standard API response for
// all relevant network or endpoint rate limit endpoints.
type domainRateLimitResponse struct {
	Data []*DomainRateLimits    `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// GetDefaultEndpointRateLimits gets the default rate limits for an endpoint in the given domain.
func (s *DomainService) GetDefaultEndpointRateLimits(ctx context.Context, domain string) (*DomainRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/ep_rate_limits/default", domain)
	body, resp, err := s.client.get(ctx, path, new(domainRateLimitResponse))
	if err != nil {
		return nil, resp, err
	}
	return body.(*domainRateLimitResponse).Data[0], resp, nil
}

// GetMaxDefaultEndpointRateLimits gets the max rate limits for an endpoint in the given domain.
func (s *DomainService) GetMaxDefaultEndpointRateLimits(ctx context.Context, domain string) (*DomainRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/ep_rate_limits/max", domain)
	body, resp, err := s.client.get(ctx, path, new(domainRateLimitResponse))
	if err != nil {
		return nil, resp, err
	}
	return body.(*domainRateLimitResponse).Data[0], resp, nil
}

// SetDefaultEndpointRateLimits sets the default rate limits for an endpoint in the given domain.
func (s *DomainService) SetDefaultEndpointRateLimits(ctx context.Context, values *DomainRateLimits, domain string) (*DomainRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/ep_rate_limits/default", domain)
	body, resp, err := s.client.put(ctx, path, new(domainRateLimitResponse), values)
	if err != nil {
		return nil, resp, err
	}
	return body.(*domainRateLimitResponse).Data[0], resp, nil
}

// SetMaxDefaultEndpointRateLimits sets the max default rate limits for an endpoint in the given domain.
func (s *DomainService) SetMaxDefaultEndpointRateLimits(ctx context.Context, values *DomainRateLimits, domain string) (*DomainRateLimits, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/ep_rate_limits/max", domain)
	body, resp, err := s.client.put(ctx, path, new(domainRateLimitResponse), values)
	if err != nil {
		return nil, resp, err
	}
	return body.(*domainRateLimitResponse).Data[0], resp, nil
}
