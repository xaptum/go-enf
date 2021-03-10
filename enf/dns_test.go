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
	"encoding/json"
	"fmt"
	"testing"
)

func TestDnsRecordJson(t *testing.T) {
	// create a txt record
	value := MakeTxtValue("v=spf; domain.com; hello")

	domain := "fd00::/48"
	id := "uuid-uu-uuid"
	name := "srv.local"
	ttl := 300
	zone := "zone-uuid-uuid"

	rec := &DnsRecord{
		Domain:     &domain,
		Id:         &id,
		Name:       &name,
		Privileged: false,
		Ttl:        &ttl,
		Type:       &DnsTXT,
		Value:      value,
		ZoneId:     &zone,
	}

	// able to marshall
	jsn, err := json.Marshal(rec)
	ok(t, err)

	// now check unmarshal
	rec1 := &DnsRecord{}
	err = json.Unmarshal(jsn, rec1)
	ok(t, err)

	fmt.Println(rec1.TXT().Txt)
	equals(t, "", string(jsn))
}
