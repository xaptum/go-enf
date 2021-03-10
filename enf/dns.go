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

type DnsRecordValue map[string]interface{}

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
	Id             *string `json:"id"`
	ZoneDomainName *string `json:"zone_domain_name"`
	Description    *string `json:"description"`
	Domain         *string `json:"enf_domain"`
	Privileged     bool    `json:"privileged"`
}

// DnsRecord uses a custom unmarshaljson function to
// unmarshall json string into the struct. This is to
// support different dns record types
type DnsRecord struct {
	Domain     *string                 `json:"enf_domain"`
	Id         *string                 `json:"id"`
	Name       *string                 `json:"name"`
	Privileged bool                    `json:"privileged"`
	Ttl        *int                    `json:"ttl"`
	Type       *string                 `json:"type"`
	Value      *map[string]interface{} `json:"value"`
	ZoneId     *string                 `json:"zone_id"`
}

func (r *DnsRecord) TXT() *txtValue {
	res := &txtValue{}

	if nil != r.Value {
		res.Txt = nil
	}

	return res
}

func (r *DnsRecord) AAAA() *aaaaValue {
	res := &aaaaValue{}

	if nil != r.Value {
		res.Ipv6 = nil
	}

	return res

}

func (r *DnsRecord) CNAME() *cnameValue {
	res := &cnameValue{}

	if nil == r.Value {
		res.Dname = nil
	}

	return res
}

func (r *DnsRecord) SRV() *srvValue {
	res := &srvValue{}

	if nil == r.Value {
		res.Weight = nil
		res.Priority = nil
		res.Port = nil
		res.Target = nil
	}

	return res
}

type DnsServer struct {
	Id          *string `json:"id"`
	Ipv6        *string `json:"ipv6"`
	Description *string `json:"description"`
	Domain      *string `json:"enf_domain"`
	Network     *string `json:"enf_network"`
}

func MakeTxtValue(txt string) *map[string]interface{} {
	m := make(map[string]interface{})
	m["txt"] = txt
	return &m
}

/*func toString(v interface{}) *string {
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
    }*/
