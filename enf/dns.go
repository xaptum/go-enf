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
package enf

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type txtValue struct {
	Txt *string `json:"txt"`
}

type aaaaValue struct {
	Ipv6 *string `json:"ipv6"`
}

type cnameValue struct {
	Dname *string `json:"dname"`
}

type srvValue struct {
	Priority *int    `json:"priority"`
	Weight   *int    `json:"weight"`
	Port     *int    `json:"port"`
	Target   *string `json:"target"`
}

type DnsZone struct {
	Id             *string    `json:"id"`
	Created        *time.Time `json:"created"`
	Modified       *time.Time `json:"modified"`
	ZoneDomainName *string    `json:"zone_domain_name"`
	Description    *string    `json:"description"`
	Domain         *string    `json:"enf_domain"`
	Privileged     bool       `json:"privileged"`
}

type NewDnsZoneReq struct {
	ZoneDomainName *string `json:"zone_domain_name"`
	Description    *string `json:"description"`
	Domain         *string `json:"enf_domain"`
	Network        *string `json:"enf_network"`
}

type UpdateDnsZoneReq struct {
	Description *string `json:"description"`
}

type AddOrReplaceNetworksReq struct {
	Networks []*string `json:"networks"`
}

type NewDnsRecordReq struct {
	Name  *string                 `json:"name"`
	Ttl   *int                    `json:"ttl"`
	Type  *string                 `json:"type"`
	Value *map[string]interface{} `json:"value"`
}

type DnsRecord struct {
	*NewDnsRecordReq
	Id         *string    `json:"id"`
	Domain     *string    `json:"enf_domain"`
	Created    *time.Time `json:"created"`
	Modified   *time.Time `json:"modified"`
	Privileged bool       `json:"privileged"`
	ZoneId     *string    `json:"zone_id"`
}

func (r *DnsRecord) TXT() *txtValue {
	res := &txtValue{}

	if nil != r.Value {
		res.Txt = toString((*r.Value)["txt"])
	}

	return res
}

func (r *DnsRecord) AAAA() *aaaaValue {
	res := &aaaaValue{}

	if nil != r.Value {
		res.Ipv6 = toString((*r.Value)["ipv6"])
	}

	return res

}

func (r *DnsRecord) CNAME() *cnameValue {
	res := &cnameValue{}

	if nil != r.Value {
		res.Dname = toString((*r.Value)["dname"])
	}

	return res
}

func (r *DnsRecord) SRV() *srvValue {
	res := &srvValue{}

	if nil != r.Value {
		res.Weight = toInt((*r.Value)["weight"])
		res.Priority = toInt((*r.Value)["priority"])
		res.Port = toInt((*r.Value)["port"])
		res.Target = toString((*r.Value)["target"])
	}

	return res
}

type NewDnsServerReq struct {
	Ipv6        *string `json:"ipv6"`
	Description *string `json:"description"`
}

type DnsServer struct {
	Id          *string    `json:"id"`
	Created     *time.Time `json:"created"`
	Modified    *time.Time `json:"modified"`
	Ipv6        *string    `json:"ipv6"`
	Description *string    `json:"description"`
	Domain      *string    `json:"enf_domain"`
	Network     *string    `json:"enf_network"`
}

type DnsZoneNetwork struct {
	Id       *int       `json:"rowid"`
	ZoneId   *string    `json:"zone_id"`
	Created  *time.Time `json:"created"`
	Modified *time.Time `json:"modified"`
	Domain   *string    `json:"enf_domain"`
	Network  *string    `json:"enf_network"`
}

type dnsZoneResp struct {
	Data []*DnsZone             `json:"data"`
	Page map[string]interface{} `json:"page"`
}

type dnsZoneNetworkResp struct {
	Data []*DnsZoneNetwork      `json:"data"`
	Page map[string]interface{} `json:"page"`
}

type dnsServerResp struct {
	Data []*DnsServer           `json:"data"`
	Page map[string]interface{} `json:"page"`
}

type dnsRecordResp struct {
	Data []*DnsRecord           `json:"data"`
	Page map[string]interface{} `json:"page"`
}

type DnsService Service

func (svc *DnsService) ListZones(ctx context.Context, domain ...string) ([]*DnsZone, error) {
	// create result struct
	result := &dnsZoneResp{}

	// create api path
	path := listZonesPath(domain...)

	// call the api
	err := svc.client.Get(ctx, path, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data, nil
}

func (svc *DnsService) CreateZone(ctx context.Context, newZone *NewDnsZoneReq) (*DnsZone, error) {
	// create result struct
	result := &dnsZoneResp{}

	// create api path
	path := DnsApiPath("/zones")

	// call the api
	err := svc.client.Post(ctx, path, newZone, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data[0], nil
}

func (svc *DnsService) UpdateZone(ctx context.Context, zoneId string, updateZoneReq *UpdateDnsZoneReq) (*DnsZone, error) {
	// create result struct
	result := &dnsZoneResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/zones/%s", zoneId))

	// call the api
	err := svc.client.Put(ctx, path, updateZoneReq, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data[0], nil
}

func (svc *DnsService) GetZone(ctx context.Context, zoneId string) (*DnsZone, error) {
	// create result struct
	result := &dnsZoneResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/zones/%s", zoneId))

	// call the api
	err := svc.client.Get(ctx, path, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data[0], nil
}

func (svc *DnsService) DeleteZone(ctx context.Context, zoneId string) error {
	// create api path
	path := DnsApiPath(fmt.Sprintf("/zones/%s", zoneId))

	// call the api
	err := svc.client.Delete(ctx, path)

	// check if request failed
	if nil != err {
		return err
	}

	return nil
}

func (svc *DnsService) AddNetworksToZone(ctx context.Context, zoneId string, networksReq *AddOrReplaceNetworksReq) ([]*DnsZoneNetwork, error) {
	// create result struct
	result := &dnsZoneNetworkResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/zones/%s/networks", zoneId))

	// call the api
	err := svc.client.Post(ctx, path, networksReq, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data, nil
}

func (svc *DnsService) ReplaceNetworksInZone(ctx context.Context, zoneId string, networksReq *AddOrReplaceNetworksReq) ([]*DnsZoneNetwork, error) {
	// create result struct
	result := &dnsZoneNetworkResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/zones/%s/networks", zoneId))

	// call the api
	err := svc.client.Put(ctx, path, networksReq, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data, nil
}

func (svc *DnsService) ListNetworksInZone(ctx context.Context, zoneId string) ([]*DnsZoneNetwork, error) {
	// create result struct
	result := &dnsZoneNetworkResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/zones/%s/networks", zoneId))

	// call the api
	err := svc.client.Get(ctx, path, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data, nil
}

func (svc *DnsService) DeleteNetworksFromZone(ctx context.Context, zoneId string, networks ...string) error {
	// create networks csv parameter
	networksCsv := strings.Join(networks, ",")
	// create api path
	path := DnsApiPath(fmt.Sprintf("/zones/%s/networks?delete=%s", zoneId, networksCsv))

	// call the api
	err := svc.client.Delete(ctx, path)

	// check if request failed
	if nil != err {
		return err
	}

	return nil
}

func (svc *DnsService) ListZonesInNetwork(ctx context.Context, network string) ([]*DnsZone, error) {
	// create result struct
	result := &dnsZoneResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/networks/%s/zones", network))

	// call the api
	err := svc.client.Get(ctx, path, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data, nil
}

func (svc *DnsService) ProvisionServer(ctx context.Context, network string, newServer *NewDnsServerReq) (*DnsServer, error) {
	// create result struct
	result := &dnsServerResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/networks/%s/servers", network))

	// call the api
	err := svc.client.Post(ctx, path, newServer, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data[0], nil
}

func (svc *DnsService) GetServer(ctx context.Context, network, ipv6 string) (*DnsServer, error) {
	// create result struct
	result := &dnsServerResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/networks/%s/servers/%s", network, ipv6))

	// call the api
	err := svc.client.Get(ctx, path, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data[0], nil
}

func (svc *DnsService) ListServers(ctx context.Context, network string) ([]*DnsServer, error) {
	// create result struct
	result := &dnsServerResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/networks/%s/servers", network))

	// call the api
	err := svc.client.Get(ctx, path, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data, nil
}

func (svc *DnsService) DeleteServer(ctx context.Context, network, ipv6 string) error {
	// create api path
	path := DnsApiPath(fmt.Sprintf("/networks/%s/servers/%s", network, ipv6))

	// call the api
	err := svc.client.Delete(ctx, path)

	// check if request failed
	if nil != err {
		return err
	}

	return nil
}

func (svc *DnsService) ListRecords(ctx context.Context, zoneId string) ([]*DnsRecord, error) {
	// create result struct
	result := &dnsRecordResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/zones/%s/records", zoneId))

	// call the api
	err := svc.client.Get(ctx, path, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data, nil
}

func (svc *DnsService) GetRecord(ctx context.Context, id string) (*DnsRecord, error) {
	// create result struct
	result := &dnsRecordResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/records/%s", id))

	// call the api
	err := svc.client.Get(ctx, path, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data[0], nil
}

func (svc *DnsService) DeleteRecord(ctx context.Context, id string) error {
	// create api path
	path := DnsApiPath(fmt.Sprintf("/records/%s", id))

	// call the api
	err := svc.client.Delete(ctx, path)

	// check if request failed
	if nil != err {
		return err
	}

	return nil
}

func (svc *DnsService) Query(ctx context.Context, network, recordType, recordName string) ([]*DnsRecord, error) {
	// create result struct
	result := &dnsRecordResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/networks/%s/query/%s/%s", network, recordType, recordName))

	// call the api
	err := svc.client.Get(ctx, path, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data, nil
}

func (svc *DnsService) CreateRecord(ctx context.Context, zoneId string, newRecord *NewDnsRecordReq) (*DnsRecord, error) {
	// create result struct
	result := &dnsRecordResp{}

	// create api path
	path := DnsApiPath(fmt.Sprintf("/zones/%s/records", zoneId))

	// call the api
	err := svc.client.Post(ctx, path, newRecord, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	return result.Data[0], nil
}

func MakeDnsTxtValue(txt string) *map[string]interface{} {
	m := make(map[string]interface{})
	m["txt"] = txt
	return &m
}

func MakeDnsAaaaValue(ipv6 string) *map[string]interface{} {
	m := make(map[string]interface{})
	m["ipv6"] = ipv6
	return &m
}

func MakeDnsCnameValue(dname string) *map[string]interface{} {
	m := make(map[string]interface{})
	m["dname"] = dname
	return &m
}

func MakeDnsSrvValue(priority, weight, port int, target string) *map[string]interface{} {
	m := make(map[string]interface{})
	m["weight"] = weight
	m["priority"] = priority
	m["port"] = port
	m["target"] = target
	return &m
}

func toString(v interface{}) *string {
	if nil == v {
		return nil
	}
	str := v.(string)
	return &str
}

func toInt(v interface{}) *int {
	if nil == v {
		return nil
	}
	i := int(v.(float64))
	return &i
}

func listZonesPath(domain ...string) string {
	if 0 == len(domain) {
		return DnsApiPath("/zones")
	} else {
		return DnsApiPath(fmt.Sprintf("/zones?enf_domain=%s", domain[0]))
	}
}
