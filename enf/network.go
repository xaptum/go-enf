package enf

import (
	"context"
	"fmt"
	"net/http"
)

// NetworkService handles communication with the network related
// methods of the ENF API. These methods are used to manage the networks
// under each domain.
type NetworkService service

// NetworkRequest is used to create a new network in Network.CreateNetwork.
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

// networkresponse represents the typical API response for all
// the endpoints in the xcr namespace
type networkresponse struct {
	Data []*Network             `json:"data"`
	Page map[string]interface{} `json:"page"`
}

// ListNetworks gets a list of all the networks under a given domain.
func (s *NetworkService) ListNetworks(ctx context.Context, domain string) ([]*Network, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/nws", domain)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	body := new(networkresponse)

	resp, err := s.client.Do(ctx, req, body)
	if err != nil {
		return nil, resp, err
	}

	return body.Data, resp, nil
}

// GetNetwork gets the network object for a given network address
// of the form <prefix>/<prefix_length>.
func (s *NetworkService) GetNetwork(ctx context.Context, network string) (*Network, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/nws/%s", network)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	body := new(networkresponse)
	resp, err := s.client.Do(ctx, req, body)
	if err != nil {
		return nil, resp, err
	}

	return body.Data[0], resp, nil
}

// CreateNetwork creates a network under the given domain.
func (s *NetworkService) CreateNetwork(ctx context.Context, domain string, network *NetworkRequest) (*Network, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/nws", domain)
	req, err := s.client.NewRequest("POST", path, network)
	if err != nil {
		return nil, nil, err
	}

	body := new(networkresponse)

	resp, err := s.client.Do(ctx, req, body)
	if err != nil {
		return nil, resp, err
	}

	return body.Data[0], resp, nil
}
