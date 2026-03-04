package proxmox

import (
	"context"
	"encoding/json"
	"log"
)

type NodeFirewallRule struct {
	Proto  string `json:"proto,omitempty"`
	Type   string `json:"type"`
	Pos    int    `json:"pos,omitempty"`
	Action string `json:"action"`
	Sport  string `json:"sport,omitempty"`
	Digest string `json:"digest,omitempty"`
	Log    string `json:"log,omitempty"`
	Dport  string `json:"dport,omitempty"`
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
	retrievedFirewallRule, err := client.ReadNodeFirewallRule(ctx, node, pos)
	if err != nil {
		return nil, err
	}
	return mapToNodeFirewallRule(retrievedFirewallRule), nil
}

func ReadAllFirewallRulesFromAPICall(ctx context.Context, node string, client *Client) (firewallRules []NodeFirewallRule, err error) {
	retrievedFirewallRules, err := client.ReadNodeFirewallRules(ctx, node)
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

func (nodeFirewallRule *NodeFirewallRule) DeleteFirewallRule(ctx context.Context, node string, client *Client) error {
	return client.DeleteNodeFirewallRule(ctx, node, nodeFirewallRule.Pos)
}

func (nodeFirewallRule *NodeFirewallRule) CreateFirewallRule(ctx context.Context, node string, client *Client) error {
	return client.CreateNodeFirewallRule(ctx, node, nodeFirewallRuleToMap(*nodeFirewallRule))
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
	sport := ""
	if _, isSet := input["sport"]; isSet {
		sport = input["sport"].(string)
	}
	digest := ""
	if _, isSet := input["digest"]; isSet {
		digest = input["digest"].(string)
	}
	log := ""
	if _, isSet := input["log"]; isSet {
		log = input["log"].(string)
	}
	dport := ""
	if _, isSet := input["dport"]; isSet {
		dport = input["dport"].(string)
	}
	enable := 0
	if _, isSet := input["enable"]; isSet {
		enable = input["enable"].(int)
	}
	return &NodeFirewallRule{
		Proto:  proto,
		Type:   ruleType,
		Pos:    pos,
		Action: action,
		Sport:  sport,
		Digest: digest,
		Log:    log,
		Dport:  dport,
		Enable: enable,
	}
}

func nodeFirewallRuleToMap(rule NodeFirewallRule) map[string]interface{} {
	result := make(map[string]interface{})
	result["proto"] = rule.Proto
	result["type"] = rule.Type
	result["pos"] = rule.Pos
	result["action"] = rule.Action
	if rule.Sport != "" {
		result["sport"] = rule.Sport
	}
	if rule.Digest != "" {
		result["digest"] = rule.Digest
	}
	if rule.Log != "" {
		result["log"] = rule.Log
	}
	if rule.Dport != "" {
		result["dport"] = rule.Dport
	}
	result["enable"] = rule.Enable
	return result
}
