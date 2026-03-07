package proxmox

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type NodeFirewallRule struct {
	Proto  string `json:"proto,omitempty"`
	Type   string `json:"type"`
	Pos    int    `json:"pos,omitempty"`
	Action string `json:"action"`

	Comment string `json:"comment,omitempty"`
	Dest    string `json:"dest,omitempty"`
	Dport   string `json:"dport,omitempty"`
	Source  string `json:"source,omitempty"`
	Sport   string `json:"sport,omitempty"`

	IcmpType  string `json:"icmp-type,omitempty"`
	Iface     string `json:"iface,omitempty"`
	Ipversion int    `json:"ipversion,omitempty"`

	Log   string `json:"log,omitempty"`
	Macro string `json:"macro,omitempty"`

	Digest string `json:"digest,omitempty"`
	Enable int    `json:"enable,omitempty"`
}

func NewNodeFirewallRule() NodeFirewallRule {
	return NodeFirewallRule{
		Type:   "in",
		Action: "REJECT",
	}
}

func NewNodeFirewallRuleFromJson(input []byte) (firewallRule NodeFirewallRule, err error) {
	firewallRule = NewNodeFirewallRule()
	err = json.Unmarshal([]byte(input), &firewallRule)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func NewNodeFirewallRuleFromAPICall(ctx context.Context, node string, pos int, client *Client) (firewallRule *NodeFirewallRule, err error) {
	retrievedFirewallRule, err := readNodeFirewallRule(ctx, node, pos, client)
	if err != nil {
		return nil, err
	}
	return mapToNodeFirewallRule(retrievedFirewallRule), nil
}

func ReadAllFirewallRulesFromAPICall(ctx context.Context, node string, client *Client) (firewallRules []NodeFirewallRule, err error) {
	retrievedFirewallRules, err := readNodeFirewallRules(ctx, node, client)
	if err != nil {
		return nil, err
	}
	var rules []NodeFirewallRule
	for _, item := range retrievedFirewallRules {
		if ruleMap, ok := item.(map[string]any); ok {
			rule := mapToNodeFirewallRule(ruleMap)
			rules = append(rules, *rule)
		}
	}
	return rules, nil
}

func (nodeFirewallRule NodeFirewallRule) DeleteFirewallRule(ctx context.Context, node string, client *Client) error {
	if client == nil {
		return errors.New(Client_Error_Nil)
	}
	err := client.Delete(ctx, fmt.Sprintf("/nodes/%s/firewall/rules/%d", node, nodeFirewallRule.Pos))
	if err != nil {
		return err
	}
	return nil
}

func (nodeFirewallRule NodeFirewallRule) CreateFirewallRule(ctx context.Context, node string, client *Client) error {
	if client == nil {
		return errors.New(Client_Error_Nil)
	}
	err := client.Post(ctx, nodeFirewallRule.mapToApiValues(), fmt.Sprintf("/nodes/%s/firewall/rules", node))
	if err != nil {
		return err
	}
	return nil
}

func (nodeFirewallRule NodeFirewallRule) UpdateFirewallRule(ctx context.Context, node string, client *Client) error {
	if client == nil {
		return errors.New(Client_Error_Nil)
	}
	err := client.Put(ctx, nodeFirewallRule.mapToApiValues(), fmt.Sprintf("/nodes/%s/firewall/rules/%d", node, nodeFirewallRule.Pos))
	if err != nil {
		return err
	}
	return nil
}

func readNodeFirewallRules(ctx context.Context, node string, client *Client) ([]interface{}, error) {
	if client == nil {
		return nil, errors.New(Client_Error_Nil)
	}
	rules, err := client.GetItemConfigInterfaceArray(ctx,
		fmt.Sprintf("/nodes/%s/firewall/rules", node),
		"node firewall",
		"rules")
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func readNodeFirewallRule(ctx context.Context, node string, pos int, client *Client) (map[string]interface{}, error) {
	if client == nil {
		return nil, errors.New(Client_Error_Nil)
	}
	rules, err := client.GetItemConfigMapStringInterface(ctx,
		fmt.Sprintf("/nodes/%s/firewall/rules/%d", node, pos),
		"node firewall",
		fmt.Sprintf("rule on position: %d", pos))
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func mapToNodeFirewallRule(input map[string]interface{}) *NodeFirewallRule {
	proto := ""
	if _, isSet := input["proto"]; isSet {
		proto = input["proto"].(string)
	}

	ruleType := ""
	if _, isSet := input["type"]; isSet {
		ruleType = input["type"].(string)
	}

	pos := 0
	if _, isSet := input["pos"]; isSet {
		pos = input["pos"].(int)
	}

	action := ""
	if _, isSet := input["action"]; isSet {
		action = input["action"].(string)
	}

	comment := ""
	if _, isSet := input["comment"]; isSet {
		comment = input["comment"].(string)
	}

	dest := ""
	if _, isSet := input["dest"]; isSet {
		dest = input["dest"].(string)
	}

	source := ""
	if _, isSet := input["source"]; isSet {
		source = input["source"].(string)
	}

	sport := ""
	if _, isSet := input["sport"]; isSet {
		sport = input["sport"].(string)
	}

	dport := ""
	if _, isSet := input["dport"]; isSet {
		dport = input["dport"].(string)
	}

	icmpType := ""
	if _, isSet := input["icmp-type"]; isSet {
		icmpType = input["icmp-type"].(string)
	}

	iface := ""
	if _, isSet := input["iface"]; isSet {
		iface = input["iface"].(string)
	}

	ipversion := 0
	if _, isSet := input["ipversion"]; isSet {
		ipversion = input["ipversion"].(int)
	}

	log := ""
	if _, isSet := input["log"]; isSet {
		log = input["log"].(string)
	}

	macro := ""
	if _, isSet := input["macro"]; isSet {
		macro = input["macro"].(string)
	}

	digest := ""
	if _, isSet := input["digest"]; isSet {
		digest = input["digest"].(string)
	}

	enable := 0
	if _, isSet := input["enable"]; isSet {
		enable = input["enable"].(int)
	}

	return &NodeFirewallRule{
		Proto:     proto,
		Type:      ruleType,
		Pos:       pos,
		Action:    action,
		Comment:   comment,
		Dest:      dest,
		Source:    source,
		Sport:     sport,
		Dport:     dport,
		IcmpType:  icmpType,
		Iface:     iface,
		Ipversion: ipversion,
		Log:       log,
		Macro:     macro,
		Digest:    digest,
		Enable:    enable,
	}
}

func (nodeFirewallRule NodeFirewallRule) mapToApiValues() map[string]interface{} {
	data, _ := json.Marshal(&nodeFirewallRule)
	var result map[string]interface{}
	_ = json.Unmarshal(data, &result)
	return result
}
