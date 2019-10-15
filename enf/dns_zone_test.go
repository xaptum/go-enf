package enf

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestDNSService_CreateDNSZone(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xdns/v1/zones", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		fmt.Fprint(w, `{
			"data": [
			  {
				"id": "test1",
				"zone_domain_name": "abc.def.xyz",
				"description": "This is a test zone",
				"enf_domain": "N/n",
				"modified": null
			  }
			],
			"page": {
			  "curr": "1",
			  "next": null,
			  "prev": null
			}
		  }
		  `)
	})

	req := &ZoneRequest{
		ZoneDomainName: String("abc.def.xyz"),
		Description:    String("This is a test zone"),
		EnfDomain:      String("N/n"),
	}
	newZone, _, err := client.DNS.CreateDNSZone(context.Background(), req)
	if err != nil {
		t.Errorf("DNS.CreateDNSZone returned error: %v", err)
	}

	want := &Zone{
		ZoneDomainName: req.ZoneDomainName,
		Description:    req.Description,
		EnfDomain:      req.EnfDomain,
	}

	if !reflect.DeepEqual(newZone, want) {
		t.Errorf("DNS.CreateDNSZone returned %+v, want %+v", newZone, want)
	}

}
