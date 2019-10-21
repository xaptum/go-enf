package enf

import (
	"context"
	"net/http"
	"testing"
)

func TestFirewallService_ListRules(t *testing.T) {
	path := "/api/xfw/v1/N/rule"

	responseBodyMock := `[
		{
			"id": "00000000-0000-4000-2000-000000000000"
		},
		{
			"id":"00000000-0000-4000-2000-000000000001"
		}
	]`

	expected := []*FirewallRule{
		{ID: String("00000000-0000-4000-2000-000000000000")},
		{ID: String("00000000-0000-4000-2000-000000000001")},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Firewall.ListRules(context.Background(), "N")
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      struct{}{},
		ResponseBodyMock: responseBodyMock,
		Expected:         expected,
		Method:           method,
		T:                t,
	}

	getTest(testParams)
}

func TestFirewallService_GetRule(t *testing.T) {
	path := "/api/xfw/v1/N/rule"

	responseBodyMock := `[
		{
			"id":"00000000-0000-4000-2000-000000000000"
		},
		{
			"id":"00000000-0000-4000-2000-000000000001"
		}
	]`

	expected := &FirewallRule{
		ID: String("00000000-0000-4000-2000-000000000001"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Firewall.GetRule(context.Background(), "N", "00000000-0000-4000-2000-000000000001")
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      struct{}{},
		ResponseBodyMock: responseBodyMock,
		Expected:         expected,
		Method:           method,
		T:                t,
	}

	getTest(testParams)
}

func TestFirewallService_CreateRule(t *testing.T) {
	path := "/api/xfw/v1/N/rule"

	requestBody := &FirewallRuleRequest{
		Priority:  Int(1),
		Action:    String("ACCEPT"),
		Direction: String("INGRESSS"),
	}

	responseBodyMock := `{
		"ID":"00000000-0000-4000-2000-000000000002"
		}
		`

	expected := &FirewallRule{
		ID: String("00000000-0000-4000-2000-000000000002"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Firewall.CreateRule(context.Background(), "N", requestBody)
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      requestBody,
		ResponseBodyMock: responseBodyMock,
		Expected:         expected,
		Method:           method,
		T:                t,
	}

	postTest(testParams)
}

func TestFirewallService_DeleteRule(t *testing.T) {
	path := "/api/xfw/v1/N/rule/00000000-0000-4000-2000-000000000000"

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.Firewall.DeleteRule(context.Background(), "N", "00000000-0000-4000-2000-000000000000")
		return struct{}{}, resp, err
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      struct{}{},
		ResponseBodyMock: "",
		Expected:         struct{}{},
		Method:           method,
		T:                t,
	}

	deleteTest(testParams)
}
