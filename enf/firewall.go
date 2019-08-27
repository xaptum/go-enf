package enf

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	//	"github.com/google/uuid"
)

var (
	ErrMissingNetwork = errors.New("Missing required network")
)

// FirewallService handles communication with the firewall related
// methods of the ENF API. These methods are used to manage the
// firewall rules for each network.
type FirewallService service

type FirewallRule struct {
	ID      *string `json:"id"`
	Network *string `json:"network"`

	Priority   *int    `json:"priority"`
	Action     *string `json:"action"`
	Direction  *string `json:"direction"`
	IPFamily   *string `json:"ip_family"`
	Protocol   *string `json:"protocol"`
	SourceIP   *string `json:"source_ip"`
	SourcePort *int    `json:"source_port"`
	DestIP     *string `json:"dest_ip"`
	DestPort   *int    `json:"dest_port"`
}

type FirewallRuleRequest struct {
	Priority   *int    `json:"priority"`
	Action     *string `json:"action"`
	Direction  *string `json:"direction"`
	IPFamily   *string `json:"ip_family"`
	Protocol   *string `json:"protocol"`
	SourceIP   *string `json:"source_ip"`
	SourcePort *int    `json:"source_port"`
	DestIP     *string `json:"dest_ip"`
	DestPort   *int    `json:"dest_port"`
}

func (s *FirewallService) ListRules(ctx context.Context, network string) ([]*FirewallRule, *http.Response, error) {
	path := fmt.Sprintf("api/xfw/v1/%v/rule", network)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var rules []*FirewallRule
	resp, err := s.client.Do(ctx, req, &rules)
	if err != nil {
		return nil, resp, err
	}

	return rules, resp, nil
}

func (s *FirewallService) GetRule(ctx context.Context, network string, id string) (*FirewallRule, *http.Response, error) {
	rules, resp, err := s.ListRules(ctx, network)
	if err != nil {
		return nil, resp, err
	}

	for _, r := range rules {
		if *r.ID == id {
			return r, resp, nil
		}
	}

	resp.StatusCode = 404
	return nil, resp, fmt.Errorf("Rule not found")
}

func (s *FirewallService) CreateRule(ctx context.Context, network string, rule *FirewallRuleRequest) (*FirewallRule, *http.Response, error) {
	path := fmt.Sprintf("api/xfw/v1/%v/rule", network)
	req, err := s.client.NewRequest("POST", path, rule)
	if err != nil {
		return nil, nil, err
	}

	r := new(FirewallRule)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (s *FirewallService) DeleteRule(ctx context.Context, network string, id string) (*http.Response, error) {
	path := fmt.Sprintf("api/xfw/v1/%v/rule/%v", network, id)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
