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

type Ratelimit struct {
	PacketsPerSecond *int `json:"packets_per_second"`
	PacketsBurstSize *int `json:"packets_burst_size"`
	BytesPerSecond   *int `json:"bytes_per_second"`
	BytesBurstSize   *int `json:"bytes_burst_size"`
	Inherit          bool `json:"inherit"`
}
