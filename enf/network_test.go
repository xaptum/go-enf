package enf

import (
	"context"
	"net/http"
	"testing"
)

func TestNetworkService_ListNetworks(t *testing.T) {
	path := "/api/xcr/v2/domains/N/nws"

	responseBodyMock := `{
		"data": [
			{
				"name": "TestNetwork 1",
				"network": "fd00:8f80:8000:0000::/64",
				"description": "This is a network.",
				"status": "ACTIVE"
			},
			{
				"name": "TestNetwork 2",
				"network": "fd00:8f80:8000:1::/64",
				"description": "This is another network.",
				"status": "ACTIVE"
			}
		],
		"page": 
		{
			"curr": -1,
			"next": -1,
			"prev": -1
		}
	}
			`

	expected := []*Network{
		{
			Name:        String("TestNetwork 1"),
			Network:     String("fd00:8f80:8000:0000::/64"),
			Description: String("This is a network."),
			Status:      String("ACTIVE"),
		},
		{
			Name:        String("TestNetwork 2"),
			Network:     String("fd00:8f80:8000:1::/64"),
			Description: String("This is another network."),
			Status:      String("ACTIVE"),
		},
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Network.ListNetworks(context.Background(), "N")
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

func TestNetworkService_GetNetwork(t *testing.T) {
	path := "/api/xcr/v2/nws/N/n"

	responseBodyMock := `{
		"data": [
		  {
			"name": "TestNetwork 1",
			"network": "N/n",
			"description": "This is a network.",
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

	expected := &Network{
		Name:        String("TestNetwork 1"),
		Network:     String("N/n"),
		Description: String("This is a network."),
		Status:      String("ACTIVE"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Network.GetNetwork(context.Background(), "N/n")
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

func TestNetworkService_CreateNetwork(t *testing.T) {
	path := "/api/xcr/v2/domains/1/nws"

	requestBody := &NetworkRequest{
		Name:        String("TestNetwork 1"),
		Description: String("This is a new network."),
	}

	responseBodyMock := `{
		"data": [
		  {
			"name": "TestNetwork 1",
			"network": "N/n",
			"description": "This is a network.",
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

	expected := &Network{
		Name:        String("TestNetwork 1"),
		Network:     String("N/n"),
		Description: String("This is a network."),
		Status:      String("ACTIVE"),
	}

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Network.CreateNetwork(context.Background(), "1", requestBody)
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

func TestNetworkService_UpdateNetwork(t *testing.T) {
	path := "/api/xcr/v2/nws/N/n"

	requestBody := &NetworkRequest{
		Name:        String("TestNetwork 334"),
		Description: String("Trying to update the network.."),
	}

	responseBodyMock := `{
		"data": [
			{
				"name": "TestNetwork 334",
				"network": "N/n",
				"description":  "Trying to update the network..",
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

	method := func(client *Client) (interface{}, *http.Response, error) {
		return client.Network.UpdateNetwork(context.Background(), "N/n", requestBody)
	}

	expected := &Network{
		Name:        String("TestNetwork 334"),
		Network:     String("N/n"),
		Description: String("Trying to update the network.."),
		Status:      String("ACTIVE"),
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
