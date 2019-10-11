package enf

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestDomainService_ListDomains(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xcr/v2/domains", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		fmt.Fprint(w, `{
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
			`)
	})

	domains, _, err := client.Domains.ListDomains(context.Background())
	if err != nil {
		t.Errorf("Domains.ListDomains returned error: %v", err)
	}

	want := []*Domain{
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
	if !reflect.DeepEqual(domains, want) {
		t.Errorf("Domains.ListDomains returned %+v, want %+v", domains, want)
	}
}

func TestDomainService_GetDomain(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xcr/v2/domains/N/n0", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		fmt.Fprint(w, `{
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
				`)
	})

	domain, _, err := client.Domains.GetDomain(context.Background(), "N/n0")
	if err != nil {
		t.Errorf("Domains.GetDomain returned error: %v", err)
	}

	want := &Domain{
		Name:    String("test.domain.1"),
		Network: String("N/n0"),
		Status:  String("ACTIVE"),
	}
	if !reflect.DeepEqual(domain, want) {
		t.Errorf("Domains.GetDomain returned %+v, want %+v", domain, want)
	}

}

func TestDomainService_CreateDomain(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xcr/v2/domains", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		fmt.Fprint(w, `{
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
			}`)
	})

	newDomainRequest := &DomainRequest{
		Name:       String("test.domain.1"),
		Type:       String("CUSTOMER_SOURCE"),
		AdminName:  String("tester"),
		AdminEmail: String("tester@test.com"),
	}
	newDomain, _, err := client.Domains.CreateDomain(context.Background(), newDomainRequest)
	if err != nil {
		t.Errorf("Domains.CreateDomain returned error: %v", err)
	}

	want := &Domain{
		Name:    String("test.domain.1"),
		Network: String("N/n0"),
		Status:  String("READY"),
	}

	if !reflect.DeepEqual(newDomain, want) {
		t.Errorf("Domains.CreateDomain returned %+v, want %+v", newDomain, want)
	}

}

func TestDomainService_ActivateDomain(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xcr/v2/domains/N/n0/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		fmt.Fprint(w, `{
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
				`)
	})

	activatedDomain, _, err := client.Domains.ActivateDomain(context.Background(), "N/n0")
	if err != nil {
		t.Errorf("Domains.ActivateDomain returned error: %v", err)
	}

	want := &Domain{
		Name:    String("test.domain.1"),
		Network: String("N/n0"),
		Status:  String("ACTIVE"),
	}
	if !reflect.DeepEqual(activatedDomain, want) {
		t.Errorf("Domains.ActivateDomain returned %+v, want %+v", activatedDomain, want)
	}
}

func TestDomainService_DeactivateDomain(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{mediaTypeJson}
	mux.HandleFunc("/api/xcr/v2/domains/N/n0/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		fmt.Fprint(w, `{
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
				`)
	})

	_, err := client.Domains.DeactivateDomain(context.Background(), "N/n0")
	if err != nil {
		t.Errorf("Domains.DeactivateDomain returned error: %v", err)
	}
}
