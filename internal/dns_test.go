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
// @since March 10, 2021
//
//-------------------------------------------------------------------------------------------
package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/xaptum/go-enf/enf"
)

func TestDnsTxtRecordJson(t *testing.T) {
	// create a txt record
	txt := "v=spf; domain.com; hello"

	name := "srv.local"
	ttl := 300
	typ := enf.DnsTXT
	value := enf.MakeDnsTxtValue(txt)

	rec1 := &enf.NewDnsRecordReq{
		Name:  &name,
		Ttl:   &ttl,
		Type:  &typ,
		Value: value,
	}

	// able to marshall
	jsn1, err := json.Marshal(rec1)
	ok(t, err)

	// now check txt after unmarshal as dns record
	rec2 := &enf.DnsRecord{}
	err = json.Unmarshal(jsn1, rec2)
	ok(t, err)
	equals(t, txt, *rec2.TXT().Txt)
}

func TestDnsAaaaRecordJson(t *testing.T) {
	// create a aaaa record
	ipv6 := "fd00::beef"

	name := "srv.local"
	ttl := 300
	value := enf.MakeDnsAaaaValue(ipv6)
	typ := enf.DnsAAAA

	rec1 := &enf.NewDnsRecordReq{
		Name:  &name,
		Ttl:   &ttl,
		Type:  &typ,
		Value: value,
	}

	// able to marshall
	jsn1, err := json.Marshal(rec1)
	ok(t, err)

	// now check txt after unmarshal as dns record
	rec2 := &enf.DnsRecord{}
	err = json.Unmarshal(jsn1, rec2)
	ok(t, err)
	equals(t, ipv6, *rec2.AAAA().Ipv6)
}

func TestDnsCnameRecordJson(t *testing.T) {
	// create a cname record
	dname := "another.local"

	name := "srv.local"
	ttl := 300
	value := enf.MakeDnsCnameValue(dname)
	typ := enf.DnsCNAME

	rec1 := &enf.NewDnsRecordReq{
		Name:  &name,
		Ttl:   &ttl,
		Type:  &typ,
		Value: value,
	}

	// able to marshall
	jsn1, err := json.Marshal(rec1)
	ok(t, err)

	// now check txt after unmarshal as dns record
	rec2 := &enf.DnsRecord{}
	err = json.Unmarshal(jsn1, rec2)
	ok(t, err)
	equals(t, dname, *rec2.CNAME().Dname)
}

func TestDnsSrvRecordJson(t *testing.T) {
	// create a srv record
	priority := 10
	weight := 20
	port := 9009
	target := "target.local"

	name := "srv.local"
	ttl := 300
	value := enf.MakeDnsSrvValue(priority, weight, port, target)
	typ := enf.DnsSRV

	rec1 := &enf.NewDnsRecordReq{
		Name:  &name,
		Ttl:   &ttl,
		Type:  &typ,
		Value: value,
	}

	// able to marshall
	jsn1, err := json.Marshal(rec1)
	ok(t, err)

	// now check txt after unmarshal
	rec2 := &enf.DnsRecord{}
	err = json.Unmarshal(jsn1, rec2)
	ok(t, err)
	equals(t, priority, *rec2.SRV().Priority)
	equals(t, weight, *rec2.SRV().Weight)
	equals(t, port, *rec2.SRV().Port)
	equals(t, target, *rec2.SRV().Target)
}

func TestListDnsZones(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:34:22.742Z","description":"A Test DNS Zone","enf_domain":"fd07:8f80:8080::/48","id":"a4ff8247-35ec-43d5-9d24-97b2a54971b3","modified":null,"privileged":true,"rowid":786,"version":0,"zone_domain_name":"test.abc.xyz"},{"created":"2021-03-17T02:34:22.742Z","description":"A Test DNS Zone","enf_domain":"fd07:8f80:8080::/48","id":"118cfa9e-0fa9-48dc-9ac9-13ad7353b236","modified":null,"privileged":false,"rowid":788,"version":0,"zone_domain_name":"no-priv.abc.xyz"}],"page":{"curr":"1","next":"","prev":""}}`
	path := enf.DnsApiPath("/zones")
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	zones, err := client.DnsSvc.ListZones(context.Background())
	ok(t, err)
	equals(t, 2, len(zones))

}

func TestListDnsZonesWithDomain(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:34:22.742Z","description":"A Test DNS Zone","enf_domain":"fd07:8f80:8080::/48","id":"a4ff8247-35ec-43d5-9d24-97b2a54971b3","modified":null,"privileged":true,"rowid":786,"version":0,"zone_domain_name":"test.abc.xyz"},{"created":"2021-03-17T02:34:22.742Z","description":"A Test DNS Zone","enf_domain":"fd07:8f80:8080::/48","id":"118cfa9e-0fa9-48dc-9ac9-13ad7353b236","modified":null,"privileged":false,"rowid":788,"version":0,"zone_domain_name":"no-priv.abc.xyz"}],"page":{"curr":"1","next":"","prev":""}}`
	domain := "fd07:8f80:8080::/48"
	path := enf.DnsApiPath(fmt.Sprintf("/zones?enf_domain=%s", domain))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	zones, err := client.DnsSvc.ListZones(context.Background(), domain)
	ok(t, err)
	equals(t, 2, len(zones))
}

func TestCreateDnsZone(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:34:21.742Z","description":"A DNS domain","enf_domain":"fd07:8f80:8080::/48","id":"0599e48d-836d-452a-a978-1101e22dce47","modified":null,"privileged":true,"rowid":725,"version":0,"zone_domain_name":"abc.def.xyz"}],"page":{"curr":"1","next":"","prev":""}}`
	path := enf.DnsApiPath("/zones")
	httpmock.RegisterResponder("POST", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	newZone := &enf.NewDnsZoneReq{}
	zone, err := client.DnsSvc.CreateZone(context.Background(), newZone)
	ok(t, err)
	equals(t, "abc.def.xyz", *zone.ZoneDomainName)
	equals(t, true, zone.Privileged)
}

func TestUpdateDnsZone(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:34:24.044Z","description":"Updated Description","enf_domain":"fd07:8f80:8080::/48","id":"295ed928-add9-42d6-97a0-bacd7bfd167c","modified":"2021-03-17T02:34:24.058Z","privileged":true,"rowid":849,"version":1,"zone_domain_name":"test.abc.xyz"}],"page":{"curr":"1","next":"","prev":""}}`
	zoneId := "295ed928-add9-42d6-97a0-bacd7bfd167c"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s", zoneId))
	httpmock.RegisterResponder("PUT", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	updZone := &enf.UpdateDnsZoneReq{}
	zone, err := client.DnsSvc.UpdateZone(context.Background(), zoneId, updZone)
	ok(t, err)
	equals(t, "Updated Description", *zone.Description)
	equals(t, "test.abc.xyz", *zone.ZoneDomainName)
	equals(t, true, zone.Privileged)
}

func TestUpdateDnsZoneFail(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"error":{"code":"validation_error","text":"DNS Zone not found"}}`
	zoneId := "295ed928-add9-42d6-97a0-bacd7bfd167c"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s", zoneId))
	httpmock.RegisterResponder("PUT", "http://localhost"+path,
		httpmock.NewStringResponder(404, fixture))

	// call the api
	updZone := &enf.UpdateDnsZoneReq{}
	zone, err := client.DnsSvc.UpdateZone(context.Background(), zoneId, updZone)
	equals(t, true, nil == zone)
	equals(t, "VALIDATION_ERROR: DNS Zone not found", err.Error())
}

func TestGetDnsZone(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:34:22.910Z","description":"A Test DNS Zone","enf_domain":"fd07:8f80:8080::/48","id":"96682f23-5ee2-418c-8077-145da14d61e0","modified":null,"privileged":true,"rowid":795,"version":0,"zone_domain_name":"test.abc.xyz"}],"page":{"curr":"1","next":"","prev":""}}`
	zoneId := "96682f23-5ee2-418c-8077-145da14d61e0"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s", zoneId))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	zone, err := client.DnsSvc.GetZone(context.Background(), zoneId)
	ok(t, err)
	equals(t, "A Test DNS Zone", *zone.Description)
	equals(t, "test.abc.xyz", *zone.ZoneDomainName)
	equals(t, true, zone.Privileged)
}

func TestGetDnsZoneFail(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"error":{"code":"get_zone_error","text":"Zone Not Found"}}`
	zoneId := "b1fc5b2b-05f8-443c-8eb5-ffc4cdb70303"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s", zoneId))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(404, fixture))

	// call the api
	zone, err := client.DnsSvc.GetZone(context.Background(), zoneId)
	equals(t, true, nil == zone)
	equals(t, "GET_ZONE_ERROR: Zone Not Found", err.Error())
}

func TestDeleteDnsZone(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `[]`
	zoneId := "b1fc5b2b-05f8-443c-8eb5-ffc4cdb70303"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s", zoneId))
	httpmock.RegisterResponder("DELETE", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	err = client.DnsSvc.DeleteZone(context.Background(), zoneId)
	ok(t, err)
}

func TestAddNetworksToZone(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:34:25.535Z","enf_domain":"fd07:8f80:8080::/48","enf_network":"fd07:8f80:8080:2::/64","modified":null,"rowid":937,"version":0,"zone_id":"b517cb51-27cf-4eca-8393-d98c03c1669e"},{"created":"2021-03-17T02:34:25.535Z","enf_domain":"fd07:8f80:8080::/48","enf_network":"fd07:8f80:8080::/64","modified":null,"rowid":938,"version":0,"zone_id":"b517cb51-27cf-4eca-8393-d98c03c1669e"}],"page":{"curr":"1","next":"","prev":""}}`
	zoneId := "b517cb51-27cf-4eca-8393-d98c03c1669e"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s/networks", zoneId))
	httpmock.RegisterResponder("POST", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	newNetworks := &enf.AddOrReplaceNetworksReq{}
	zoneNetworks, err := client.DnsSvc.AddNetworksToZone(context.Background(), zoneId, newNetworks)
	ok(t, err)
	equals(t, 2, len(zoneNetworks))

	equals(t, zoneId, *zoneNetworks[0].ZoneId)
	equals(t, zoneId, *zoneNetworks[1].ZoneId)
}

func TestReplaceNetworksInZone(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:34:27.854Z","enf_domain":"fd07:8f80:8080::/48","enf_network":"fd07:8f80:8080:1::/64","modified":null,"rowid":1065,"version":0,"zone_id":"651c5757-b5c9-4607-a0a3-85f2046c6199"}],"page":{"curr":"1","next":"","prev":""}}`
	zoneId := "651c5757-b5c9-4607-a0a3-85f2046c6199"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s/networks", zoneId))
	httpmock.RegisterResponder("PUT", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	newNetworks := &enf.AddOrReplaceNetworksReq{}
	zoneNetworks, err := client.DnsSvc.ReplaceNetworksInZone(context.Background(), zoneId, newNetworks)
	ok(t, err)
	equals(t, 1, len(zoneNetworks))

	equals(t, zoneId, *zoneNetworks[0].ZoneId)
}

func TestListNetworksInZone(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:34:26.247Z","enf_domain":"fd07:8f80:8080::/48","enf_network":"fd07:8f80:8080:200::/64","modified":null,"rowid":977,"version":0,"zone_id":"a1127ddc-ca0e-469c-9096-582d0e56e907"},{"created":"2021-03-17T02:34:26.263Z","enf_domain":"fd07:8f80:8080::/48","enf_network":"fd07:8f80:8080:2::/64","modified":null,"rowid":980,"version":0,"zone_id":"a1127ddc-ca0e-469c-9096-582d0e56e907"},{"created":"2021-03-17T02:34:26.263Z","enf_domain":"fd07:8f80:8080::/48","enf_network":"fd07:8f80:8080::/64","modified":null,"rowid":981,"version":0,"zone_id":"a1127ddc-ca0e-469c-9096-582d0e56e907"}],"page":{"curr":"1","next":"","prev":""}}`
	zoneId := "a1127ddc-ca0e-469c-9096-582d0e56e907"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s/networks", zoneId))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	zoneNetworks, err := client.DnsSvc.ListNetworksInZone(context.Background(), zoneId)
	ok(t, err)
	equals(t, 3, len(zoneNetworks))

	equals(t, zoneId, *zoneNetworks[0].ZoneId)
	equals(t, zoneId, *zoneNetworks[1].ZoneId)
	equals(t, zoneId, *zoneNetworks[2].ZoneId)
}

func TestDeleteNetworksFromZone(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `[]`
	zoneId := "4846dfce-4545-4f32-b65d-36e55f8bc827/networks"
	networks := []string{"fd07:8f80:8080::/64", "fd07:8f80:8080:2::/64"}
	networksCsv := strings.Join(networks, ",")
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s/networks?delete=%s", zoneId, networksCsv))
	httpmock.RegisterResponder("DELETE", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	err = client.DnsSvc.DeleteNetworksFromZone(context.Background(), zoneId, networks...)
	ok(t, err)
}

func TestListZonesInNetwork(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:43:40.380Z","description":"A Test DNS Zone","enf_domain":"fd07:8f80:8080::/48","id":"b4105b27-8d96-42e9-8b36-f7731b50c033","modified":null,"privileged":true,"rowid":1059,"version":0,"zone_domain_name":"test.abc.xyz"}],"page":{"curr":"1","next":"","prev":""}}`
	network := "fd07:8f80:8080:200::/64"
	path := enf.DnsApiPath(fmt.Sprintf("/networks/%s/zones", network))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	zones, err := client.DnsSvc.ListZonesInNetwork(context.Background(), network)
	ok(t, err)
	equals(t, 1, len(zones))

	equals(t, "b4105b27-8d96-42e9-8b36-f7731b50c033", *zones[0].Id)
}

func TestProvisionDnsServer(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:43:41.286Z","description":null,"enf_domain":"fd07:8f80:8080::/48","enf_network":"fd07:8f80:8080:200::/64","id":"9e4df6a7-cff9-43eb-ba7e-58ec9ba38720","ipv6":"fd00:8f80:8000:3:e152:b47a:b656:cc4d","modified":"2021-03-17T02:43:41.286Z","rowid":35,"version":1}],"page":{"curr":"1","next":"","prev":""}}`
	network := "fd07:8f80:8080:200::/64"
	path := enf.DnsApiPath(fmt.Sprintf("/networks/%s/servers", network))
	httpmock.RegisterResponder("POST", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	newServer := &enf.NewDnsServerReq{}
	server, err := client.DnsSvc.ProvisionServer(context.Background(), network, newServer)
	ok(t, err)

	equals(t, network, *server.Network)
	equals(t, true, len(*server.Ipv6) > 0)
}

func TestGetDnsServer(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:43:42.775Z","description":null,"enf_domain":"fd07:8f80:8080::/48","enf_network":"fd07:8f80:8080:200::/64","id":"55088b95-31f1-4abf-89bd-dab9f11f82e4","ipv6":"fd00:8f80:8000:3:e152:b47a:b656:cc4d","modified":"2021-03-17T02:43:42.775Z","rowid":47,"version":1}],"page":{"curr":"1","next":"","prev":""}}`
	network := "fd07:8f80:8080:200::/64"
	ipv6 := "fd00:8f80:8000:3:e152:b47a:b656:cc4d"
	path := enf.DnsApiPath(fmt.Sprintf("/networks/%s/servers/%s", network, ipv6))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	server, err := client.DnsSvc.GetServer(context.Background(), network, ipv6)
	ok(t, err)

	equals(t, network, *server.Network)
	equals(t, ipv6, *server.Ipv6)
}

func TestListDnsServers(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:43:42.348Z","description":null,"enf_domain":"fd07:8f80:8080::/48","enf_network":"fd07:8f80:8080:200::/64","id":"a845bf0b-63a3-4d6c-874f-a0d6f3c423dd","ipv6":"fd00:8f80:8000:3:e152:b47a:b656:cc4d","modified":"2021-03-17T02:43:42.348Z","rowid":42,"version":1}],"page":{"curr":"1","next":"","prev":""}}`
	network := "fd07:8f80:8080:200::/64"
	path := enf.DnsApiPath(fmt.Sprintf("/networks/%s/servers", network))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	servers, err := client.DnsSvc.ListServers(context.Background(), network)
	ok(t, err)

	equals(t, 1, len(servers))
	equals(t, network, *servers[0].Network)
}

func TestDeleteDnsServer(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `[]`
	network := "fd07:8f80:8080:200::/64"
	ipv6 := "fd00:8f80:8000:3:e152:b47a:b656:cc4d"
	path := enf.DnsApiPath(fmt.Sprintf("/networks/%s/servers/%s", network, ipv6))
	httpmock.RegisterResponder("DELETE", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	err = client.DnsSvc.DeleteServer(context.Background(), network, ipv6)
	ok(t, err)
}

func TestListDnsRecords(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:51:13.689Z","enf_domain":"fd07:8f80:8080::/48","id":"e87ab6fa-23ff-403e-b182-9f85979d611a","modified":null,"name":"hvac-01.test.abc.xyz","privileged":true,"rowid":1215,"ttl":300,"type":"AAAA","value":{"ipv6":"fd07:2607:6560:23f1:2::fd87"},"version":0,"zone_id":"8e20dc05-0128-44ed-9abb-30981f473c01"}],"page":{"curr":"1","next":"","prev":""}}`
	zoneId := "8e20dc05-0128-44ed-9abb-30981f473c01"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s/records", zoneId))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	records, err := client.DnsSvc.ListRecords(context.Background(), zoneId)
	ok(t, err)

	equals(t, 1, len(records))
	equals(t, zoneId, *records[0].ZoneId)
}

func TestGetDnsRecord(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:51:14.715Z","enf_domain":"fd07:8f80:8080::/48","id":"b368fcee-4780-463d-9eb9-c6c715c39a21","modified":null,"name":"hvac-01.test.abc.xyz","privileged":true,"rowid":1254,"ttl":300,"type":"AAAA","value":{"ipv6":"fd07:2607:6560:23f1:2::fd87"},"version":0,"zone_id":"52d75c3b-3767-402c-84ba-c8d20ddb18c8"}],"page":{"curr":"1","next":"","prev":""}}`
	recordId := "b368fcee-4780-463d-9eb9-c6c715c39a21"
	path := enf.DnsApiPath(fmt.Sprintf("/records/%s", recordId))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	record, err := client.DnsSvc.GetRecord(context.Background(), recordId)
	ok(t, err)
	equals(t, recordId, *record.Id)
}

func TestDeleteDnsRecord(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `[]`
	recordId := "b368fcee-4780-463d-9eb9-c6c715c39a21"
	path := enf.DnsApiPath(fmt.Sprintf("/records/%s", recordId))
	httpmock.RegisterResponder("DELETE", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	err = client.DnsSvc.DeleteRecord(context.Background(), recordId)
	ok(t, err)
}

func TestQueryDnsRecord(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:51:16.362Z","enf_domain":"fd07:8f80:8080::/48","id":"dd362f82-841b-4631-bdd4-73451779a38c","modified":null,"name":"hvac-01.test.abc.xyz","privileged":true,"rowid":1323,"ttl":300,"type":"AAAA","value":{"ipv6":"fd07:2607:6560:23f1:2::fd87"},"version":0,"zone_id":"18c01f52-c22b-41c7-a956-e17d0afee88f"}],"page":{"curr":"1","next":"","prev":""}}`
	network := "fd07:8f80:8080:200::/64"
	recordType := "AAAA"
	recordName := "hvac-01.test.abc.xyz"
	path := enf.DnsApiPath(fmt.Sprintf("/networks/%s/query/%s/%s", network, recordType, recordName))
	httpmock.RegisterResponder("GET", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	records, err := client.DnsSvc.Query(context.Background(), network, recordType, recordName)
	ok(t, err)
	equals(t, 1, len(records))

	record := records[0]
	equals(t, "AAAA", *record.Type)
	equals(t, "fd07:2607:6560:23f1:2::fd87", *record.AAAA().Ipv6)
}

func TestCreateTxtDnsRecord(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:51:32.772Z","enf_domain":"fd07:8f80:8080::/48","id":"5f46c115-8789-4455-8e2d-baef0d75a545","modified":null,"name":"hvac-03.test.abc.xyz","privileged":true,"rowid":1396,"ttl":300,"type":"TXT","value":{"txt":"v=spf;mail;txt"},"version":0,"zone_id":"ce3282d9-bba2-40f6-ad69-2b212742e50b"}],"page":{"curr":"1","next":"","prev":""}}`
	zoneId := "ce3282d9-bba2-40f6-ad69-2b212742e50b"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s/records", zoneId))
	httpmock.RegisterResponder("POST", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	newRecord := &enf.NewDnsRecordReq{}
	record, err := client.DnsSvc.CreateRecord(context.Background(), zoneId, newRecord)
	ok(t, err)
	equals(t, zoneId, *record.ZoneId)
	equals(t, "TXT", *record.Type)
	equals(t, "v=spf;mail;txt", *record.TXT().Txt)
	equals(t, true, nil == record.AAAA().Ipv6)
	equals(t, true, nil == record.CNAME().Dname)
	equals(t, true, nil == record.SRV().Priority)
	equals(t, true, nil == record.SRV().Weight)
	equals(t, true, nil == record.SRV().Port)
	equals(t, true, nil == record.SRV().Target)
}

func TestCreateSrvDnsRecord(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:51:32.834Z","enf_domain":"fd07:8f80:8080::/48","id":"4254da87-f7b0-4e28-8b8a-dd61cfd2efb0","modified":null,"name":"hvac-05.test.abc.xyz","privileged":true,"rowid":1400,"ttl":300,"type":"SRV","value":{"port":9090,"priority":10,"target":"google.com","weight":20},"version":0,"zone_id":"2f35329f-0db9-465b-93be-50f0450fca73"}],"page":{"curr":"1","next":"","prev":""}}`
	zoneId := "2f35329f-0db9-465b-93be-50f0450fca73"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s/records", zoneId))
	httpmock.RegisterResponder("POST", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	newRecord := &enf.NewDnsRecordReq{}
	record, err := client.DnsSvc.CreateRecord(context.Background(), zoneId, newRecord)
	ok(t, err)
	equals(t, zoneId, *record.ZoneId)
	equals(t, "SRV", *record.Type)
	equals(t, true, nil == record.TXT().Txt)
	equals(t, true, nil == record.AAAA().Ipv6)
	equals(t, true, nil == record.CNAME().Dname)
	equals(t, 10, *record.SRV().Priority)
	equals(t, 20, *record.SRV().Weight)
	equals(t, 9090, *record.SRV().Port)
	equals(t, "google.com", *record.SRV().Target)
}

func TestCreateCnameDnsRecord(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:51:32.875Z","enf_domain":"fd07:8f80:8080::/48","id":"b051d647-56fe-4e63-812a-4634de396745","modified":null,"name":"hvac-04.test.abc.xyz","privileged":true,"rowid":1404,"ttl":300,"type":"CNAME","value":{"dname":"google.com"},"version":0,"zone_id":"ca7ce03e-72b6-47f9-a076-adf255b9326f"}],"page":{"curr":"1","next":"","prev":""}}`
	zoneId := "ca7ce03e-72b6-47f9-a076-adf255b9326f"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s/records", zoneId))
	httpmock.RegisterResponder("POST", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	newRecord := &enf.NewDnsRecordReq{}
	record, err := client.DnsSvc.CreateRecord(context.Background(), zoneId, newRecord)
	ok(t, err)
	equals(t, zoneId, *record.ZoneId)
	equals(t, "CNAME", *record.Type)
	equals(t, true, nil == record.TXT().Txt)
	equals(t, true, nil == record.AAAA().Ipv6)
	equals(t, "google.com", *record.CNAME().Dname)
	equals(t, true, nil == record.SRV().Priority)
	equals(t, true, nil == record.SRV().Weight)
	equals(t, true, nil == record.SRV().Port)
	equals(t, true, nil == record.SRV().Target)
}

func TestCreateAaaaDnsRecord(t *testing.T) {
	// client client
	client, err := mockClient("http://localhost")
	ok(t, err)

	// deactivate http mock
	defer httpmock.DeactivateAndReset()

	// setup mocks
	fixture := `{"data":[{"created":"2021-03-17T02:51:33.047Z","enf_domain":"fd07:8f80:8080::/48","id":"c0999a46-e711-4412-94fd-644d4fb6ca86","modified":null,"name":"test.abc.xyz","privileged":true,"rowid":1415,"ttl":300,"type":"AAAA","value":{"ipv6":"fd07:2607:6560:23f1:2::fd88"},"version":0,"zone_id":"4af0ac09-db8c-43fd-950c-4e2af2681b73"}],"page":{"curr":"1","next":"","prev":""}}`
	zoneId := "4af0ac09-db8c-43fd-950c-4e2af2681b73"
	path := enf.DnsApiPath(fmt.Sprintf("/zones/%s/records", zoneId))
	httpmock.RegisterResponder("POST", "http://localhost"+path,
		httpmock.NewStringResponder(200, fixture))

	// call the api
	newRecord := &enf.NewDnsRecordReq{}
	record, err := client.DnsSvc.CreateRecord(context.Background(), zoneId, newRecord)
	ok(t, err)
	equals(t, zoneId, *record.ZoneId)
	equals(t, "AAAA", *record.Type)
	equals(t, true, nil == record.TXT().Txt)
	equals(t, "fd07:2607:6560:23f1:2::fd88", *record.AAAA().Ipv6)
	equals(t, true, nil == record.CNAME().Dname)
	equals(t, true, nil == record.SRV().Priority)
	equals(t, true, nil == record.SRV().Weight)
	equals(t, true, nil == record.SRV().Port)
	equals(t, true, nil == record.SRV().Target)
}
