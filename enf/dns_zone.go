package enf

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// DNSService handles communication with the DNS related
// methods of the ENF API. These methods are used to manage the DNS zones,
// DNS records, and service endpoints on the ENF.
type DNSService service

// Zone represents a DNS zone within the ENF.
type Zone struct {
	Created        *time.Time `json:"created"`
	Description    *string    `json:"description"`
	EnfDomain      *string    `json:"enf_domain"`
	ID             *string    `json:"id"`
	Modified       *time.Time `json:"modified"`
	ZoneDomainName *string    `json:"zone_domain_name"`
}

// CreateZoneRequest represents a request to create a DNS zone within the ENF.
// Uses the same fields as the Zone struct.
type CreateZoneRequest struct {
	Description    *string `json:"description"`
	ZoneDomainName *string `json:"zone_domain_name"`
}

// UpdateZoneRequest represents a request to update a DNS zone within the ENF.
type UpdateZoneRequest struct {
	Description *string `json:"description"`
}

type zoneResponse struct {
	Data []*Zone                `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// ListZones lists all the DNS zones for a given ENF domain (::/48 address).
func (s *DNSService) ListZones(ctx context.Context) ([]*Zone, *http.Response, error) {
	path := "api/xdns/2019-05-27/zones"
	body, resp, err := s.client.get(ctx, path, new(zoneResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*zoneResponse).Data, resp, nil
}

// GetZone gets a DNS zone given its UUID.
func (s *DNSService) GetZone(ctx context.Context, zoneUUID string) (*Zone, *http.Response, error) {
	path := fmt.Sprintf("api/xdns/2019-05-27/zones/%v", zoneUUID)
	body, resp, err := s.client.get(ctx, path, new(zoneResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*zoneResponse).Data[0], resp, nil
}

// CreateZone creates a new DNS zone.
func (s *DNSService) CreateZone(ctx context.Context, req *CreateZoneRequest) (*Zone, *http.Response, error) {
	path := "api/xdns/2019-05-27/zones"
	body, resp, err := s.client.post(ctx, path, new(zoneResponse), req)
	if err != nil {
		return nil, resp, err
	}

	return body.(*zoneResponse).Data[0], resp, nil
}

// UpdateZone updates a zone's description given its UUID.
func (s *DNSService) UpdateZone(ctx context.Context, zoneUUID string, req *UpdateZoneRequest) (*Zone, *http.Response, error) {
	path := fmt.Sprintf("api/xdns/2019-05-27/zones/%v", zoneUUID)
	body, resp, err := s.client.put(ctx, path, new(zoneResponse), req)
	if err != nil {
		return nil, resp, err
	}

	return body.(*zoneResponse).Data[0], resp, nil
}

// DeleteZone deletes a zone given its UUID.
func (s *DNSService) DeleteZone(ctx context.Context, zoneUUID string) (*http.Response, error) {
	path := fmt.Sprintf("api/xdns/2019-05-27/zones/%v", zoneUUID)
	resp, err := s.client.delete(ctx, path)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
