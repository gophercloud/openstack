package policies

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
)

// PortCreateOptsExt adds QoS options to the base ports.CreateOpts.
type PortCreateOptsExt struct {
	ports.CreateOptsBuilder

	// QoSPolicyID represents an associated QoS policy.
	QoSPolicyID string `json:"qos_policy_id,omitempty"`
}

// ToPortCreateMap casts a CreateOpts struct to a map.
func (opts PortCreateOptsExt) ToPortCreateMap() (map[string]interface{}, error) {
	base, err := opts.CreateOptsBuilder.ToPortCreateMap()
	if err != nil {
		return nil, err
	}

	port := base["port"].(map[string]interface{})

	if opts.QoSPolicyID != "" {
		port["qos_policy_id"] = opts.QoSPolicyID
	}

	return base, nil
}

// PortUpdateOptsExt adds QoS options to the base ports.UpdateOpts.
type PortUpdateOptsExt struct {
	ports.UpdateOptsBuilder

	// QoSPolicyID represents an associated QoS policy.
	// Setting it to a pointer of an empty string will remove associated QoS policy from port.
	QoSPolicyID *string `json:"qos_policy_id,omitempty"`
}

// ToPortUpdateMap casts a UpdateOpts struct to a map.
func (opts PortUpdateOptsExt) ToPortUpdateMap() (map[string]interface{}, error) {
	base, err := opts.UpdateOptsBuilder.ToPortUpdateMap()
	if err != nil {
		return nil, err
	}

	port := base["port"].(map[string]interface{})

	if opts.QoSPolicyID != nil {
		qosPolicyID := *opts.QoSPolicyID
		if qosPolicyID != "" {
			port["qos_policy_id"] = qosPolicyID
		} else {
			port["qos_policy_id"] = nil
		}
	}

	return base, nil
}
