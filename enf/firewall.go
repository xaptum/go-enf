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
)

type FirewallService Service

type FirewallRuleRequest struct {
	Priority   *int    `json:"priority"`
	Protocol   *string `json:"protocol"`
	Direction  *string `json:"direction"`
	SourceIp   *string `json:"source_ip"`
	SourcePort *int    `json:"source_port"`
	DestIp     *string `json:"dest_ip"`
	DestPort   *int    `json:"dest_port"`
	Action     *string `json:"action"`
	IpFamily   *string `json:"ip_family"`
}

type FirewallRule struct {
	*FirewallRuleRequest
	Id      *string `json:"id"`
	Network *string `json:"network"`
}

type firewallRuleResponse struct {
	Data []*FirewallRule        `json:"data"`
	Page map[string]interface{} `json:"page"`
}

func (svc *FirewallService) AddRule(ctx context.Context, newRule *FirewallRuleRequest, network ...string) (*FirewallRule, error) {
	// create reulst struct
	result := &firewallRuleResponse{}

	// create api path
	path := addRulePath(network...)

	// call the api
	err := svc.client.Post(ctx, path, newRule, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	// return the new rule
	return result.Data[0], nil
}

func (svc *FirewallService) ListRules(ctx context.Context, network string) ([]*FirewallRule, error) {
	// create reulst struct
	result := &firewallRuleResponse{}

	// create api path
	path := FirewallApiPath(fmt.Sprintf("/%s/rule", network))

	// call the api
	err := svc.client.Get(ctx, path, result)

	// check if request failed
	if nil != err {
		return nil, err
	}

	// return the new rule
	return result.Data, nil
}

func (svc *FirewallService) DeleteRule(ctx context.Context, network, id string) error {
	// create path
	path := FirewallApiPath(fmt.Sprintf("/%s/rule/%s", network, id))

	// call the api
	err := svc.client.Delete(ctx, path, nil)

	// check if request failed
	if nil != err {
		return err
	}

	return nil
}

func addRulePath(network ...string) string {
	if 0 == len(network) {
		return FirewallApiPath("/rule")
	} else {
		return FirewallApiPath(fmt.Sprintf("/%s/rule", network[0]))
	}
}
