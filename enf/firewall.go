package enf

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var (
	ErrMissingNetwork = errors.New("Missing required network")
)

// FirewallService handles communication with the firewall related
// methods of the ENF API. These methods are used to manage the
// firewall rules for each network.
type FirewallService service

// FirewallRule represents a firewall rule for a network in the ENF.
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

// FirewallRuleRequest represents the body of the request for creating a firewall rule.
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

// ListRules gets all the firewall rules for the given network.
func (s *FirewallService) ListRules(ctx context.Context, network string) ([]*FirewallRule, *http.Response, error) {
	path := fmt.Sprintf("api/xfw/v1/%v/rule", network)
	body, resp, err := s.client.get(ctx, path, url.Values{}, new([]*FirewallRule))
	if err != nil {
		return nil, resp, err
	}
	return *(body.(*[]*FirewallRule)), resp, nil
}

// GetRule gets the information for the firewall rule with the given id within the given network.
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

// CreateRule creates a firewall rule for the given network.
func (s *FirewallService) CreateRule(ctx context.Context, network string, rule *FirewallRuleRequest) (*FirewallRule, *http.Response, error) {
	path := fmt.Sprintf("api/xfw/v1/%v/rule", network)
	body, resp, err := s.client.post(ctx, path, new(FirewallRule), rule)
	if err != nil {
		return nil, resp, err
	}
	return body.(*FirewallRule), resp, nil
}

// DeleteRule deletes the firewall rule associated with the given network address and ID.
func (s *FirewallService) DeleteRule(ctx context.Context, network string, id string) (*http.Response, error) {
	path := fmt.Sprintf("api/xfw/v1/%v/rule/%v", network, id)
	return s.client.delete(ctx, path, url.Values{})
}
