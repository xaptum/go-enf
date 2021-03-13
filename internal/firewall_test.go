//-------------------------------------------------------------------------------------------
//
// XAPTUM CONFIDENTIAL
// __________________
//
//  2021(C) Xaptum, Inc.
//  All Rights Reserved.Patents Pending.
//
// NOTICE:  All information contained herein is, and remains
// the property of Xaptum, Inc.  The intellectual and technical concepts contained
// herein are proprietary to Xaptum, Inc and may be covered by U.S. and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Xaptum, Inc.
//
// @author Venkatakumar Srinivasan
// @since March 12, 2021
//
//-------------------------------------------------------------------------------------------
package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/xaptum/go-enf/enf"
)

func TestFirewallRuleJson(t *testing.T) {
	fixture := `{"id":"43e0ba98-a939-4417-aa77-59d88e2d1c0a","priority":10,"protocol":"ICMP6","direction":"EGRESS","source_ip":"fd08:8f80:8000:1::/64","source_port":0,"dest_ip":"*","dest_port":0,"action":"ACCEPT","ip_family":"IP6","network":"fd08:8f80:8000:1::/64"}`

	rule := &enf.FirewallRule{}

	// unmarshal
	err := json.Unmarshal([]byte(fixture), rule)
	ok(t, err)

	// check
	equals(t, "43e0ba98-a939-4417-aa77-59d88e2d1c0a", *rule.Id)
	equals(t, 10, *rule.Priority)
}

func TestDeleteRuleSuccess(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	network := "fd08:8f80:8000:1::/64"
	id := "43e0ba98-a939-4417-aa77-59d88e2d1c0a"
	path := enf.FirewallApiPath(fmt.Sprintf("/%s/rule/%s", network, id))
	httpmock.RegisterResponder("DELETE", "http://localhost"+path,
		httpmock.NewStringResponder(200, "[]"))

	// call the api
	err = client.FirewallSvc.DeleteRule(context.Background(), network, id)
	ok(t, err)
}

func TestDeleteRuleFail(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"error":{"code":"firewall_error","text":"Unable to delete firewall rule\n"}}`
	network := "fd08:8f80:8000:1::/64"
	id := "43e0ba98-a939-4417-aa77-59d88e2d1c0a"
	path := enf.FirewallApiPath(fmt.Sprintf("/%s/rule/%s", network, id))
	httpmock.RegisterResponder("DELETE", "http://localhost"+path,
		httpmock.NewStringResponder(400, fixture))

	// call the api
	err = client.FirewallSvc.DeleteRule(context.Background(), network, id)
	equals(t, "FIREWALL_ERROR: Unable to delete firewall rule\n", err.Error())
}

func TestListFirewallRules(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"id":"1f019638-af33-41bf-a296-6dc987ee960d","priority":10,"protocol":"UDP","direction":"EGRESS","source_ip":"fd08:8f80:8000:1::/64","source_port":0,"dest_ip":"*","dest_port":0,"action":"ACCEPT","ip_family":"IP6","network":"fd08:8f80:8000:1::/64"},{"id":"205a1af8-1b0c-4e62-a74e-de5ca02b02db","priority":10,"protocol":"UDP","direction":"INGRESS","source_ip":"*","source_port":0,"dest_ip":"fd08:8f80:8000:1::/64","dest_port":0,"action":"ACCEPT","ip_family":"IP6","network":"fd08:8f80:8000:1::/64"},{"id":"fa04a275-ff57-43a1-bcfd-a64857f0662a","priority":10,"protocol":"TCP","direction":"INGRESS","source_ip":"*","source_port":0,"dest_ip":"fd08:8f80:8000:1::/64","dest_port":0,"action":"ACCEPT","ip_family":"IP6","network":"fd08:8f80:8000:1::/64"}]}`
	network := "fd08:8f80:8000:1::/64"
	path := enf.FirewallApiPath(fmt.Sprintf("/%s/rule", network))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	rules, err := client.FirewallSvc.ListRules(context.Background(), network)
	ok(t, err)
	equals(t, 3, len(rules))
}

func TestListFirewallRulesEmpty(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data": []}`
	network := "fd08:8f80:8000:1::/64"
	path := enf.FirewallApiPath(fmt.Sprintf("/%s/rule", network))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	rules, err := client.FirewallSvc.ListRules(context.Background(), network)
	ok(t, err)
	equals(t, 0, len(rules))
}

func TestListFirewallRuleServerError(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"error":{"code":"firewall_error","text":"Internal Server Error"}}`
	network := "fd08:8f80:8000:1::/64"
	path := enf.FirewallApiPath(fmt.Sprintf("/%s/rule", network))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(500, fixture))

	// call the api
	rules, err := client.FirewallSvc.ListRules(context.Background(), network)
	equals(t, "FIREWALL_ERROR: Internal Server Error", err.Error())
	equals(t, true, nil == rules)
}

func TestAddRuleWithNetworkNotAuthorized(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := "{\"error\":{\"code\":\"firewall_error\",\"text\":\"Not Authorized\"}}"
	network := "fd08:8f80:8000:1::/64"
	path := enf.FirewallApiPath(fmt.Sprintf("/%s/rule", network))
	httpmock.RegisterResponder("POST", "http://localhost"+path,
		httpmock.NewStringResponder(401, fixture))

	// create new rule request
	newRule := &enf.FirewallRuleRequest{}

	// call the api
	rules, err := client.FirewallSvc.AddRule(context.Background(), newRule, network)
	equals(t, "FIREWALL_ERROR: Not Authorized", err.Error())
	equals(t, true, nil == rules)
}

func TestAddRuleNotAuthorized(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := "{\"error\":{\"code\":\"firewall_error\",\"text\":\"Not Authorized\"}}"
	path := enf.FirewallApiPath("/rule")
	httpmock.RegisterResponder("POST", "http://localhost"+path,
		httpmock.NewStringResponder(401, fixture))

	// create new rule request
	newRule := &enf.FirewallRuleRequest{}

	// call the api
	rules, err := client.FirewallSvc.AddRule(context.Background(), newRule)
	equals(t, "FIREWALL_ERROR: Not Authorized", err.Error())
	equals(t, true, nil == rules)
}

func TestAddRuleFail(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := "{\"error\":{\"code\":\"firewall_error\",\"text\":\"Invalid Request\"}}"
	path := enf.FirewallApiPath("/rule")
	httpmock.RegisterResponder("POST", "http://localhost"+path,
		httpmock.NewStringResponder(400, fixture))

	// create new rule request
	newRule := &enf.FirewallRuleRequest{}

	// call the api
	rules, err := client.FirewallSvc.AddRule(context.Background(), newRule)
	equals(t, "FIREWALL_ERROR: Invalid Request", err.Error())
	equals(t, true, nil == rules)
}

func TestAddRuleSuccess(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data": [{"id":"43e0ba98-a939-4417-aa77-59d88e2d1c0a","priority":10,"protocol":"ICMP6","direction":"EGRESS","source_ip":"fd08:8f80:8000:1::/64","source_port":0,"dest_ip":"*","dest_port":0,"action":"ACCEPT","ip_family":"IP6","network":"fd08:8f80:8000:1::/64"}]}`
	path := enf.FirewallApiPath("/rule")
	httpmock.RegisterResponder("POST", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// create new rule request
	family := enf.IP6
	priority := 10
	protocol := enf.ICMP6
	sourceIp := "*"
	sourcePort := 0
	destIp := "*"
	destPort := 0
	direction := enf.Egress
	action := enf.Accept
	newRule := &enf.FirewallRuleRequest{
		IpFamily:   &family,
		Priority:   &priority,
		Protocol:   &protocol,
		SourceIp:   &sourceIp,
		SourcePort: &sourcePort,
		DestIp:     &destIp,
		DestPort:   &destPort,
		Direction:  &direction,
		Action:     &action,
	}

	// call the api
	rule, err := client.FirewallSvc.AddRule(context.Background(), newRule)
	ok(t, err)
	equals(t, true, nil != rule)
}
