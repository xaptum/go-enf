package enf

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

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

type domainResponse struct {
	Data []*Domain              `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// ListDomains lists all available domains on the ENF.
func (s *DomainService) ListDomains(ctx context.Context) ([]*Domain, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains")
	body, resp, err := s.client.get(ctx, path, url.Values{}, new(domainResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*domainResponse).Data, resp, nil
}

// GetDomain gets the information of a specified domain.
func (s *DomainService) GetDomain(ctx context.Context, domain string) (*Domain, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v", domain)
	body, resp, err := s.client.get(ctx, path, url.Values{}, new(domainResponse))
	if err != nil {
		return nil, resp, err
	}

	return body.(*domainResponse).Data[0], resp, nil
}

// CreateDomain provisions a domain on the ENF.
func (s *DomainService) CreateDomain(ctx context.Context, req *DomainRequest) (*Domain, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains")
	body, resp, err := s.client.post(ctx, path, new(domainResponse), req)
	if err != nil {
		return nil, resp, err
	}

	return body.(*domainResponse).Data[0], resp, nil
}

// ActivateDomain activates the given domain (sets the status field to ACTIVE)
func (s *DomainService) ActivateDomain(ctx context.Context, domain string) (*Domain, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/status", domain)
	body, resp, err := s.client.put(ctx, path, new(domainResponse), "ACTIVE")
	if err != nil {
		return nil, resp, err
	}

	return body.(*domainResponse).Data[0], resp, nil
}

// DeactivateDomain deactivates the given domain (sets the status field to READY)
func (s *DomainService) DeactivateDomain(ctx context.Context, domain string) (*Domain, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/status", domain)
	body, resp, err := s.client.put(ctx, path, new(domainResponse), "READY")
	if err != nil {
		return nil, resp, err
	}
	return body.(*domainResponse).Data[0], resp, nil
}
