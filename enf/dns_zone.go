package enf

import (
	"context"
	"fmt"
	"time"
)

// DNSService handles communication with the DNS related
// methods of the ENF API. These methods are used to manage the DNS zones,
// DNS records, and service endpoints on the ENF.
type DNSService Service

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

type ZoneResponse struct {
	Data []*Zone                `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// ListZones lists all the DNS zones for a given ENF domain (::/48 address).
func (s *DNSService) ListZones(ctx context.Context) ([]*Zone, error) {
	// create request
	req := NewRequest("/api/xdns/2019-05-27/zones", nil, nil, nil)

	// call the api
	zoneResponse := new(ZoneResponse)
	_, err := s.client.get(ctx, req, zoneResponse)
	if err != nil {
		return nil, err
	}

	// return data
	return zoneResponse.Data, nil
}

// GetZone gets a DNS zone given its UUID.
func (s *DNSService) GetZone(ctx context.Context, zoneUUID string) (*Zone, error) {
	// create request
	path := fmt.Sprintf("/api/xdns/2019-05-27/zones/%v", zoneUUID)
	req := NewRequest(path, nil, nil, nil)

	// call api
	zoneResponse := new(ZoneResponse)
	_, err := s.client.get(ctx, req, zoneResponse)
	if err != nil {
		return nil, err
	}

	// return data
	return zoneResponse.Data[0], nil
}

// CreateZone creates a new DNS zone.
func (s *DNSService) CreateZone(ctx context.Context, createZoneReq *CreateZoneRequest) (*Zone, error) {
	// create request
	req := NewRequest("/api/xdns/2019-05-27/zones", nil, nil, createZoneReq)

	// call api
	zoneResponse := new(ZoneResponse)
	_, err := s.client.post(ctx, req, zoneResponse)
	if err != nil {
		return nil, err
	}

	// return data
	return zoneResponse.Data[0], nil
}

// UpdateZone updates a zone's description given its UUID.
func (s *DNSService) UpdateZone(ctx context.Context, zoneUUID string, updateZoneReq *UpdateZoneRequest) (*Zone, error) {
	// create request
	path := fmt.Sprintf("/api/xdns/2019-05-27/zones/%v", zoneUUID)
	req := NewRequest(path, nil, nil, updateZoneReq)

	// call api
	zoneResponse := new(ZoneResponse)
	_, err := s.client.put(ctx, req, zoneResponse)
	if err != nil {
		return nil, err
	}

	// return data
	return zoneResponse.Data[0], nil
}

// DeleteZone deletes a zone given its UUID.
func (s *DNSService) DeleteZone(ctx context.Context, zoneUUID string) error {
	// create request
	path := fmt.Sprintf("/api/xdns/2019-05-27/zones/%v", zoneUUID)
	req := NewRequest(path, nil, nil, nil)

	// call api
	_, err := s.client.delete(ctx, req, nil)
	if err != nil {
		return err
	}

	// return success
	return nil
}
