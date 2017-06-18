package tenants

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Tenant is a grouping of users in the identity service.
type Tenant struct {
	// ID is a unique identifier for this tenant.
	ID string `json:"id"`

	// Name is a friendlier user-facing name for this tenant.
	Name string `json:"name"`

	// Description is a human-readable explanation of this Tenant's purpose.
	Description string `json:"description"`

	// Enabled indicates whether or not a tenant is active.
	Enabled bool `json:"enabled"`
}

// TenantPage is a single page of Tenant results.
type TenantPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Tenants contains any results.
func (r TenantPage) IsEmpty() (bool, error) {
	tenants, err := ExtractTenants(r)
	return len(tenants) == 0, err
}

// NextPageURL extracts the "next" link from the tenants_links section of the result.
func (r TenantPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"tenants_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// ExtractTenants returns a slice of Tenants contained in a single page of results.
func ExtractTenants(r pagination.Page) ([]Tenant, error) {
	var s struct {
		Tenants []Tenant `json:"tenants"`
	}
	err := (r.(TenantPage)).ExtractInto(&s)
	return s.Tenants, err
}

type tenantResult struct {
	gophercloud.Result
}

// Extract interprets any tenantResults as a tenant.
func (r tenantResult) Extract() (*Tenant, error) {
	var s struct {
		tenant *Tenant `json:"tenant"`
	}
	err := r.ExtractInto(&s)
	return s.tenant, err
}

// GetResult temporarily contains the response from the Get call.
type GetResult struct {
	tenantResult
}

// CreateResult temporarily contains the reponse from the Create call.
type CreateResult struct {
	tenantResult
}

// DeleteResult temporarily contains the response from the Delete call.
type DeleteResult struct {
	tenantResult
}

// UpdateResult temporarily contains the response from the Update call.
type UpdateResult struct {
	tenantResult
}
