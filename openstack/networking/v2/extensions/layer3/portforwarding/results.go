package portforwarding

import (
	"github.com/gophercloud/gophercloud"
)

type PortForwarding struct {
	// The ID of the floating IP port forwarding
	ID string `json:"id"`

	// The ID of the Neutron port associated to the floating IP port forwarding.
	InternalPortID string `json:"internal_port_id"`

	// The TCP/UDP/other protocol port number of the port forwarding’s floating IP address.
	ExternalPort int `json:"external_port"`

	// The IP protocol used in the floating IP port forwarding.
	Protocol string `json:"protocol"`

	// The TCP/UDP/other protocol port number of the Neutron port fixed
	// IP address associated to the floating ip port forwarding.
	InternalPort int `json:"internal_port"`

	// The fixed IPv4 address of the Neutron port associated
	// to the floating IP port forwarding.
	InternalIPAddress string `json:"internal_ip_address"`

	// The project ID
	ProjectID string `json:"project_id"`
}

type commonResult struct {
	gophercloud.Result
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a PortForwarding.
type CreateResult struct {
	commonResult
}

// Extract will extract a Port Forwarding resource from a result.
func (r commonResult) Extract() (*PortForwarding, error) {
	var s PortForwarding
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "port_forwarding")
}
