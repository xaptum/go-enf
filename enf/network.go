package enf

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// NetworkService handles communication with the network related
// methods of the ENF API. These methods are used to manage the networks
// under each domain.
type NetworkService service

// NetworkRequest is used to create a new network or update the information of an existing one.
type NetworkRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

// Network represents a network in the ENF.
type Network struct {
	Name        *string `json:"name"`
	Network     *string `json:"network"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

// networkResponse represents the typical API response for all
// the endpoints in the xcr namespace.
type networkResponse struct {
	Data []*Network             `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// ListNetworks gets a list of all the networks under a given domain.
func (s *NetworkService) ListNetworks(ctx context.Context, domain string) ([]*Network, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/nws", domain)
	body, resp, err := s.client.get(ctx, path, url.Values{}, new(networkResponse))
	if err != nil {
		return nil, resp, err
	}
	return body.(*networkResponse).Data, resp, nil
}

// GetNetwork gets the network object for a given network address of the form <prefix>/<prefix_length>.
func (s *NetworkService) GetNetwork(ctx context.Context, network string) (*Network, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/nws/%s", network)
	body, resp, err := s.client.get(ctx, path, url.Values{}, new(networkResponse))
	if err != nil {
		return nil, resp, err
	}
	return body.(*networkResponse).Data[0], resp, nil
}

// CreateNetwork creates a network with the given fields under the given domain.
func (s *NetworkService) CreateNetwork(ctx context.Context, domain string, fields *NetworkRequest) (*Network, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/nws", domain)
	body, resp, err := s.client.post(ctx, path, new(networkResponse), fields)
	if err != nil {
		return nil, resp, err
	}
	return body.(*networkResponse).Data[0], resp, nil
}

// UpdateNetwork updates the name and/or description of an existing network.
func (s *NetworkService) UpdateNetwork(ctx context.Context, network string, fields *NetworkRequest) (*Network, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/nws/%s", network)
	body, resp, err := s.client.put(ctx, path, new(networkResponse), fields)
	if err != nil {
		return nil, resp, err
	}
	return body.(*networkResponse).Data[0], resp, nil
}
