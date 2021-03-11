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

type DomainService Service

type Domain struct {
	Id          *int    `json:"id"`
	Cidr        *string `json:"cidr"`
	Allocated   *string `json:"allocated"`
	Type        *string `json:"type"`
	Name        *string `json:"name"`
	Status      *string `json:"status"`
	LastNetwork *string `json:"last_network"`
}

type DomainEpRatelimit struct {
	Default *Ratelimit `json:"default"`
	Max     *Ratelimit `json:"max"`
}
