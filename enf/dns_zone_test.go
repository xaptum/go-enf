package enf

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestDNSService_ListZones(t *testing.T) {
	path := "/api/xdns/2019-05-27/zones"

	responseBodyMock := `{
		"data": [
			{
				"description": "abc.def test domain",
				"enf_domain": "N/n0",
				"id": "1234",
				"zone_domain_name": "abc.def"
			},
			{
				"enf_domain": "N/n0",
				"id": "5678",
				"zone_domain_name": "def.ghi"
			}
		]
	}`

	expected := []*Zone{
		{
			Description:    String("abc.def test domain"),
			EnfDomain:      String("N/n0"),
			ID:             String("1234"),
			ZoneDomainName: String("abc.def"),
		},
		{
			EnfDomain:      String("N/n0"),
			ID:             String("5678"),
			ZoneDomainName: String("def.ghi"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.DNS.ListZones(context.Background())
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

func TestDNSService_GetZone(t *testing.T) {
	path := "/api/xdns/2019-05-27/zones/1234"

	responseBodyMock := `{
		"data": [
			{
				"created": "2019-10-21T19:48:21.961747Z",
				"description": "Test zone",
				"enf_domain": "N/n0",
				"id": "1234",
				"zone_domain_name": "abc.def",
				"modified": null
			}
		]
	}
	`

	createdTime, _ := time.Parse(time.RFC3339, "2019-10-21T19:48:21.961747Z")

	expected := &Zone{
		Created:        Time(createdTime),
		Description:    String("Test zone"),
		EnfDomain:      String("N/n0"),
		ID:             String("1234"),
		Modified:       nil,
		ZoneDomainName: String("abc.def"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.DNS.GetZone(context.Background(), "1234")
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

func TestDNSService_CreateZone(t *testing.T) {
	path := "/api/xdns/2019-05-27/zones"

	requestBody := &CreateZoneRequest{
		Description:    String("Test create zone"),
		ZoneDomainName: String("abc.def"),
	}

	responseBodyMock := `{
		"data": [
			{
				"description": "Test create zone",
				"enf_domain": "N/n0",
				"id": "1234",
				"zone_domain_name": "abc.def"
			}
		]
	}
	`

	expected := &Zone{
		Description:    String("Test create zone"),
		EnfDomain:      String("N/n0"),
		ID:             String("1234"),
		ZoneDomainName: String("abc.def"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.DNS.CreateZone(context.Background(), requestBody)
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

func TestDNSService_UpdateZone(t *testing.T) {
	path := "/api/xdns/2019-05-27/zones/1234"

	requestBody := &UpdateZoneRequest{
		Description: String("Test update zone"),
	}

	responseBodyMock := `{
		"data": [
			{
				"description": "Test update zone",
				"enf_domain": "N/n0",
				"id": "1234",
				"zone_domain_name": "abc.def"
			}
		]
	}`

	expected := &Zone{
		Description:    String("Test update zone"),
		EnfDomain:      String("N/n0"),
		ID:             String("1234"),
		ZoneDomainName: String("abc.def"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.DNS.UpdateZone(context.Background(), "1234", requestBody)
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      requestBody,
		ResponseBodyMock: responseBodyMock,
		Expected:         expected,
		Method:           method,
		T:                t,
	}

	putTest(testParams)
}

func TestDNSService_DeleteZone(t *testing.T) {
	path := "/api/xdns/2019-05-27/zones/1234"

	method := func(client *Client) (interface{}, *http.Response, error) {
		resp, err := client.DNS.DeleteZone(context.Background(), "1234")
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
