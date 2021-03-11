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

type FirewallService Service

type FirewallRule struct {
	Id         *int    `json:"id"`
	Priority   *int    `json:"priority"`
	Protocol   *string `json:"protocol"`
	Direction  *string `json:"direction"`
	SourceIp   *string `json:"source_ip"`
	SourcePort *int    `json:"source_port"`
	DestIp     *string `json:"dest_ip"`
	DestPort   *int    `json:"dest_port"`
	Action     *string `json:"action"`
	IpFamily   *string `json:"ip_family"`
	Network    *string `json:"network"`
}
