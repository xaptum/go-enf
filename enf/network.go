package enf

import (
	"context"
	"fmt"
	"net/http"
)

type NetworkService service

type NetworkRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type Network struct {
	Name        *string `json:"name"`
	Network     *string `json:"network"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

func (s *NetworkService) ListNetworks(ctx context.Context, domain string) ([]*Network, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/domains/%v/nws", domain)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	body := &struct {
		Data []*Network             `json:"data"`
		Page map[string]interface{} `json:"page"`
	}{}

	resp, err := s.client.Do(ctx, req, body)
	if err != nil {
		return nil, resp, err
	}

	return body.Data, resp, nil
}

func (s *NetworkService) GetNetwork(ctx context.Context, network string) (*Network, *http.Response, error) {
	path := fmt.Sprintf("api/xcr/v2/nws/%s", network)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	body := &struct {
		Data []*Network             `json:"data"`
		Page map[string]interface{} `json:"page"`
	}{}

	resp, err := s.client.Do(ctx, req, body)
	if err != nil {
		return nil, resp, err
	}

	return body.Data[0], resp, nil
}
