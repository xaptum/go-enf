package enf

import (
	"context"
	"net/http"
	"testing"
)

func TestDomainService_ListDomains(t *testing.T) {
	path := "/api/xcr/v2/domains"

	responseBodyMock := `{
		"data": [
			{
				"network": "N/n0",
				"name": "test.domain.1",
				"status": "ACTIVE"
			},
			{
				"network": "N/n1",
				"name": "test.domain.2",
				"status": "ACTIVE"
			}
		],
		"page": {
			"curr": -1,
			"next": -1,
			"prev": -1
		}
	}
			`

	expected := []*Domain{
		{
			Name:    String("test.domain.1"),
			Network: String("N/n0"),
			Status:  String("ACTIVE"),
		},
		{
			Name:    String("test.domain.2"),
			Network: String("N/n1"),
			Status:  String("ACTIVE"),
		},
	}
	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Domains.ListDomains(context.Background())
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

func TestDomainService_GetDomain(t *testing.T) {
	path := "/api/xcr/v2/domains/N/n0"

	responseBodyMock := `{
		"data": [
			{
				"network": "N/n0",
				"name": "test.domain.1",
				"status": "ACTIVE"
			}
		],
		"page": {
			"curr": -1,
			"next": -1,
			"prev": -1
		}
	}
			`

	expected := &Domain{
		Name:    String("test.domain.1"),
		Network: String("N/n0"),
		Status:  String("ACTIVE"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Domains.GetDomain(context.Background(), "N/n0")
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

func TestDomainService_CreateDomain(t *testing.T) {
	path := "/api/xcr/v2/domains"

	requestBody := &DomainRequest{
		Name:       String("test.domain.1"),
		Type:       String("CUSTOMER_SOURCE"),
		AdminName:  String("tester"),
		AdminEmail: String("tester@test.com"),
	}

	responseBodyMock := `{
		"data": [
			{
				"network": "N/n0",
				"name": "test.domain.1",
				"status": "READY"
			}
		],
		"page": {
			"curr": -1,
			"next": -1,
			"prev": -1
		}
		}`

	expected := &Domain{
		Name:    String("test.domain.1"),
		Network: String("N/n0"),
		Status:  String("READY"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Domains.CreateDomain(context.Background(), requestBody)
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

func TestDomainService_ActivateDomain(t *testing.T) {
	path := "/api/xcr/v2/domains/N/n0/status"

	responseBodyMock := `{
		"data": [
			{
				"network": "N/n0",
				"name": "test.domain.1",
				"status": "ACTIVE"
			}
		],
		"page": {
			"curr": -1,
			"next": -1,
			"prev": -1
		}
	}
			`

	expected := &Domain{
		Name:    String("test.domain.1"),
		Network: String("N/n0"),
		Status:  String("ACTIVE"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Domains.ActivateDomain(context.Background(), "N/n0")
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      struct{}{},
		ResponseBodyMock: responseBodyMock,
		Expected:         expected,
		Method:           method,
		T:                t,
	}

	putTest(testParams)
}

func TestDomainService_DeactivateDomain(t *testing.T) {
	path := "/api/xcr/v2/domains/N/n0/status"

	responseBodyMock := `{
		"data": [
			{
				"network": "N/n0",
				"name": "test.domain.1",
				"status": "READY"
			}
		],
		"page": {
			"curr": -1,
			"next": -1,
			"prev": -1
		}
	}
			`

	expected := &Domain{
		Name:    String("test.domain.1"),
		Network: String("N/n0"),
		Status:  String("READY"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Domains.DeactivateDomain(context.Background(), "N/n0")
	}

	testParams := &TestParams{
		Path:             path,
		RequestBody:      struct{}{},
		ResponseBodyMock: responseBodyMock,
		Expected:         expected,
		Method:           method,
		T:                t,
	}

	putTest(testParams)
}
