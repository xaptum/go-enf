package enf

import (
	"context"
	"fmt"
	"net/http"
)

// DNSService handles communication with the DNS related
// methods of the ENF API. These methods are used to manage the DNS zones,
// DNS records, and service endpoints on the ENF.
type DNSService service

// Zone represents a DNS zone within the ENF.
type Zone struct {
	ZoneDomainName *string `json:"zone_domain_name"`
	Description    *string `json:"description"`
	EnfDomain      *string `json:"enf_domain"`
}

// ZoneRequest represents a request to create a DNS zone within the ENF.
// Uses the same fields as the Zone struct.
type ZoneRequest Zone

type zoneResponse struct {
	Data []*Zone                `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// CreateDNSZone creates a new DNS zone.
func (s *DNSService) CreateDNSZone(ctx context.Context, req *ZoneRequest) (*Zone, *http.Response, error) {
	path := "api/xdns/v1/zones"
	body, resp, err := s.client.post(ctx, path, new(zoneResponse), req)
	if err != nil {
		return nil, resp, err
	}

	return body.(*zoneResponse).Data[0], resp, nil
}

// ListDNSZones lists all the DNS zones for a given ENF domain (::/48 address).
func (s *DNSService) ListDNSZones(ctx context.Context, domain string) ([]*Zone, *http.Response, error) {
	path := fmt.Sprintf("xdns/v1/zones?enf_domain=%v", domain)
	body, resp, err := s.client.get(ctx, path, new(zoneResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*zoneResponse).Data, resp, nil
}
