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

type NetworkService Service

type Network struct {
	Id          *int    `json:"id"`
	Description *string `json:"description"`
	DomainId    *int    `json:"domain_id"`
	Domain      *string `json:"domain"`
	Name        *string `json:"name"`
	Status      *string `json:"status"`
	Cidr        *string `json:"cidr"`
	Tags        *string `json:"tags"`
}

type NetworkEpRatelimit struct {
	Default *Ratelimit `json:"default"`
	Max     *Ratelimit `json:"max"`
}
