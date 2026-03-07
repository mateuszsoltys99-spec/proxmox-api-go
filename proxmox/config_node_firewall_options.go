package proxmox

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type NodeFirewallOptions struct {
	Enable                           bool   `json:"enable,omitempty"`
	LogLevelForward                  string `json:"log_level_forward,omitempty"`
	LogLevelIn                       string `json:"log_level_in,omitempty"`
	LogLevelOut                      string `json:"log_level_out,omitempty"`
	LogNfConntrack                   bool   `json:"log_nf_conntrack,omitempty"`
	Ndp                              bool   `json:"ndp,omitempty"`
	NfConntrackAllowInvalid          bool   `json:"nf_conntrack_allow_invalid,omitempty"`
	NfConntrackHelpers               string `json:"nf_conntrack_helpers,omitempty"`
	NfConntrackMax                   int    `json:"nf_conntrack_max,omitempty"`
	NfConntrackTcpTimeoutEstablished int    `json:"nf_conntrack_tcp_timeout_established,omitempty"`
	NfConntrackTcpTimeoutSynRecv     int    `json:"nf_conntrack_tcp_timeout_syn_recv,omitempty"`
	Nftables                         bool   `json:"nftables,omitempty"`
	Nosmurfs                         bool   `json:"nosmurfs,omitempty"`
	ProtectionSynflood               bool   `json:"protection_synflood,omitempty"`
	ProtectionSynfloodBurst          int    `json:"protection_synflood_burst,omitempty"`
	ProtectionSynfloodRate           int    `json:"protection_synflood_rate,omitempty"`
	SmurfLogLevel                    string `json:"smurf_log_level,omitempty"`
	TcpFlagsLogLevel                 string `json:"tcp_flags_log_level,omitempty"`
	Tcpflags                         bool   `json:"tcpflags,omitempty"`
}

func NewNodeFirewallOptions() NodeFirewallOptions {
	return NodeFirewallOptions{}
}

func NewNodeFirewallOptionsFromJson(input []byte) (firewallOptions NodeFirewallOptions, err error) {
	firewallOptions = NewNodeFirewallOptions()
	err = json.Unmarshal([]byte(input), &firewallOptions)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func NewNodeFirewallOptionsFromAPICall(ctx context.Context, node string, client *Client) (firewallOptions *NodeFirewallOptions, err error) {
	if client == nil {
		return nil, errors.New(Client_Error_Nil)
	}
	rules, err := client.GetItemConfigMapStringInterface(ctx,
		fmt.Sprintf("/nodes/%s/firewall/options", node),
		"node firewall",
		fmt.Sprintf("options for node: %s", node))
	if err != nil {
		return nil, err
	}
	return mapToNodeFirewallOptions(rules), nil
}

func (nodeFirewallOptions NodeFirewallOptions) UpdateNodeFirewallOptions(ctx context.Context, node string, client *Client) error {
	if client == nil {
		return errors.New(Client_Error_Nil)
	}
	err := client.Put(ctx, nodeFirewallOptions.mapToApiValues(), fmt.Sprintf("/nodes/%s/firewall/options", node))
	if err != nil {
		return err
	}
	return nil
}

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func getBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

func getInt(m map[string]interface{}, key string) int {
	if v, ok := m[key]; ok {
		switch val := v.(type) {
		case int:
			return val
		case float64: // common when coming from json.Unmarshal
			return int(val)
		}
	}
	return 0
}

func mapToNodeFirewallOptions(input map[string]interface{}) *NodeFirewallOptions {
	return &NodeFirewallOptions{
		Enable:                           getBool(input, "enable"),
		LogLevelForward:                  getString(input, "log_level_forward"),
		LogLevelIn:                       getString(input, "log_level_in"),
		LogLevelOut:                      getString(input, "log_level_out"),
		LogNfConntrack:                   getBool(input, "log_nf_conntrack"),
		Ndp:                              getBool(input, "ndp"),
		NfConntrackAllowInvalid:          getBool(input, "nf_conntrack_allow_invalid"),
		NfConntrackHelpers:               getString(input, "nf_conntrack_helpers"),
		NfConntrackMax:                   getInt(input, "nf_conntrack_max"),
		NfConntrackTcpTimeoutEstablished: getInt(input, "nf_conntrack_tcp_timeout_established"),
		NfConntrackTcpTimeoutSynRecv:     getInt(input, "nf_conntrack_tcp_timeout_syn_recv"),
		Nftables:                         getBool(input, "nftables"),
		Nosmurfs:                         getBool(input, "nosmurfs"),
		ProtectionSynflood:               getBool(input, "protection_synflood"),
		ProtectionSynfloodBurst:          getInt(input, "protection_synflood_burst"),
		ProtectionSynfloodRate:           getInt(input, "protection_synflood_rate"),
		SmurfLogLevel:                    getString(input, "smurf_log_level"),
		TcpFlagsLogLevel:                 getString(input, "tcp_flags_log_level"),
		Tcpflags:                         getBool(input, "tcpflags"),
	}
}

func (nodeFirewallOptions NodeFirewallOptions) mapToApiValues() map[string]interface{} {
	data, _ := json.Marshal(&nodeFirewallOptions)
	var result map[string]interface{}
	_ = json.Unmarshal(data, &result)
	return result
}
