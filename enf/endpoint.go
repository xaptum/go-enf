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

import "time"

type EndpointService Service

type EndpointEvent struct {
	Id         *int       `json:"id"`
	Network    *string    `json:"network"`
	Ipv6       *string    `json:"ipv6"`
	Type       *string    `json:"type"`
	AsnOrg     *string    `json:"asn_org"`
	InstanceId *int       `json:"instance_id"`
	PopId      *int       `json:"pop_id"`
	RemoteIp   *string    `json:"remote_ip"`
	RemotePort *string    `json:"remote_port"`
	SessionId  *string    `json:"session_id"`
	Source     *string    `json:"source"`
	Timestamp  *time.Time `json:"timestamp"`
}

type Endpoint struct {
	Id        *int           `json:"id"`
	Domain    *string        `json:"domain"`
	DomainId  *int           `json:"domain_id"`
	Ipv6      *string        `json:"ipv6"`
	LastEvent *EndpointEvent `json:"last_event"`
	Name      *string        `json:"name"`
	Network   *string        `json:"network"`
	NetworkId *int           `json:"network_id"`
	State     *string        `json:"endpoint_state"`
	Tags      *string        `json:"tags"`
}

type EndpointRatelimit struct {
	Current *Ratelimit `json:"current"`
	Max     *Ratelimit `json:"max"`
}
