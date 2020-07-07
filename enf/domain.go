package enf

import (
	"context"
	"fmt"
)

// DomainService handles communication with the domain-related methods
// of the ENF API.
type DomainService Service

// Domain represents a domain in the ENF.
type Domain struct {
	Name    *string `json:"name"`
	Network *string `json:"network"`
	Status  *string `json:"status"`
}

// DomainRequest represents a request to provision a new domain.
type DomainRequest struct {
	Name       *string `json:"name"`
	Type       *string `json:"type"`
	AdminName  *string `json:"admin_name"`
	AdminEmail *string `json:"admin_email"`
}

type DomainResponse struct {
	Data []*Domain              `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// ListDomains lists all available domains on the ENF.
func (s *DomainService) ListDomains(ctx context.Context) ([]*Domain, error) {
	// create request
	req := NewRequest("/api/xcr/v3/domains", nil, nil, nil)

	// call api
	domainResponse := new(DomainResponse)
	_, err := s.client.get(ctx, req, domainResponse)
	if err != nil {
		return nil, err
	}

	// return data
	return domainResponse.Data, nil
}

// GetDomain gets the information of a specified domain.
func (s *DomainService) GetDomain(ctx context.Context, domain string) (*Domain, error) {
	// create request
	path := fmt.Sprintf("/api/xcr/v3/domains/%v", domain)
	req := NewRequest(path, nil, nil, nil)

	// call api
	domainResponse := new(DomainResponse)
	_, err := s.client.get(ctx, req, domainResponse)
	if err != nil {
		return nil, err
	}

	// return data
	return domainResponse.Data[0], nil
}

// CreateDomain provisions a domain on the ENF.
func (s *DomainService) CreateDomain(ctx context.Context, createDomainReq *DomainRequest) (*Domain, error) {
	// create request
	req := NewRequest("/api/xcr/v3/domains", nil, nil, createDomainReq)

	// call api
	domainResponse := new(DomainResponse)
	_, err := s.client.post(ctx, req, domainResponse)
	if err != nil {
		return nil, err
	}

	// return data
	return domainResponse.Data[0], nil
}

// ActivateDomain activates the given domain (sets the status field to ACTIVE)
func (s *DomainService) ActivateDomain(ctx context.Context, domain string) (*Domain, error) {
	return s.updateDomainStatus(ctx, domain, "ACTIVE")
}

// DeactivateDomain deactivates the given domain (sets the status field to READY)
func (s *DomainService) DeactivateDomain(ctx context.Context, domain string) (*Domain, error) {
	return s.updateDomainStatus(ctx, domain, "READY")
}

func (s *DomainService) updateDomainStatus(ctx context.Context, domain, status String) (*Domain, error) {
	// create request
	path := fmt.Sprintf("/api/xcr/v3/domains/%v/status", domain)
	req := NewRequest(path, nil, nil, status)

	// call api
	domainResponse := new(DomainResponse)
	_, err := s.client.put(ctx, req, domainResponse)
	if err != nil {
		return nil, err
	}

	return domainResponse.Data[0], nil
}
