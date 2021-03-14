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
	"encoding/json"
	"testing"

	"github.com/xaptum/go-enf/enf"
)

func TestDnsTxtRecordJson(t *testing.T) {
	// create a txt record
	txt := "v=spf; domain.com; hello"
	value := enf.MakeTxtValue(txt)

	domain := "fd00::/48"
	id := "uuid-uu-uuid"
	name := "srv.local"
	ttl := 300
	zone := "zone-uuid-uuid"
	typ := enf.DnsTXT

	rec1 := &enf.DnsRecord{
		Domain:     &domain,
		Id:         &id,
		Name:       &name,
		Privileged: false,
		Ttl:        &ttl,
		Type:       &typ,
		Value:      value,
		ZoneId:     &zone,
	}

	// able to marshall
	jsn1, err := json.Marshal(rec1)
	ok(t, err)

	// now check txt after unmarshal
	rec2 := &enf.DnsRecord{}
	err = json.Unmarshal(jsn1, rec2)
	ok(t, err)
	equals(t, txt, *rec2.TXT().Txt)

	// now marshall rec1 and see json is same
	jsn2, err := json.Marshal(rec2)
	ok(t, err)
	equals(t, string(jsn1), string(jsn2))

}

func TestDnsAaaaRecordJson(t *testing.T) {
	// create a aaaa record
	ipv6 := "fd00::beef"
	value := enf.MakeAaaaValue(ipv6)

	domain := "fd00::/48"
	id := "uuid-uu-uuid"
	name := "srv.local"
	ttl := 300
	zone := "zone-uuid-uuid"
	typ := enf.DnsAAAA

	rec1 := &enf.DnsRecord{
		Domain:     &domain,
		Id:         &id,
		Name:       &name,
		Privileged: false,
		Ttl:        &ttl,
		Type:       &typ,
		Value:      value,
		ZoneId:     &zone,
	}

	// able to marshall
	jsn1, err := json.Marshal(rec1)
	ok(t, err)

	// now check txt after unmarshal
	rec2 := &enf.DnsRecord{}
	err = json.Unmarshal(jsn1, rec2)
	ok(t, err)
	equals(t, ipv6, *rec2.AAAA().Ipv6)

	// now marshall rec1 and see json is same
	jsn2, err := json.Marshal(rec2)
	ok(t, err)
	equals(t, string(jsn1), string(jsn2))

}

func TestDnsCnameRecordJson(t *testing.T) {
	// create a cname record
	dname := "another.local"
	value := enf.MakeCnameValue(dname)

	domain := "fd00::/48"
	id := "uuid-uu-uuid"
	name := "srv.local"
	ttl := 300
	zone := "zone-uuid-uuid"
	typ := enf.DnsCNAME

	rec1 := &enf.DnsRecord{
		Domain:     &domain,
		Id:         &id,
		Name:       &name,
		Privileged: false,
		Ttl:        &ttl,
		Type:       &typ,
		Value:      value,
		ZoneId:     &zone,
	}

	// able to marshall
	jsn1, err := json.Marshal(rec1)
	ok(t, err)

	// now check txt after unmarshal
	rec2 := &enf.DnsRecord{}
	err = json.Unmarshal(jsn1, rec2)
	ok(t, err)
	equals(t, dname, *rec2.CNAME().Dname)

	// now marshall rec1 and see json is same
	jsn2, err := json.Marshal(rec2)
	ok(t, err)
	equals(t, string(jsn1), string(jsn2))

}

func TestDnsSrvRecordJson(t *testing.T) {
	// create a srv record
	priority := 10
	weight := 20
	port := 9009
	target := "target.local"
	value := enf.MakeSrvValue(priority, weight, port, target)

	domain := "fd00::/48"
	id := "uuid-uu-uuid"
	name := "srv.local"
	ttl := 300
	zone := "zone-uuid-uuid"
	typ := enf.DnsSRV

	rec1 := &enf.DnsRecord{
		Domain:     &domain,
		Id:         &id,
		Name:       &name,
		Privileged: false,
		Ttl:        &ttl,
		Type:       &typ,
		Value:      value,
		ZoneId:     &zone,
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

	// now marshall rec1 and see json is same
	jsn2, err := json.Marshal(rec2)
	ok(t, err)
	equals(t, string(jsn1), string(jsn2))

}
