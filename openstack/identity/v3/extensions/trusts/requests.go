package trusts

import (
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/gophercloud/gophercloud/pagination"
)

// AuthOptsExt extends the base Identity v3 tokens AuthOpts with a TrustID.
type AuthOptsExt struct {
	tokens.AuthOptionsBuilder

	// TrustID is the ID of the trust.
	TrustID string `json:"id"`
}

// ToTokenV3CreateMap builds a create request body from the AuthOpts.
func (opts AuthOptsExt) ToTokenV3CreateMap(scope map[string]interface{}) (map[string]interface{}, error) {
	return opts.AuthOptionsBuilder.ToTokenV3CreateMap(scope)
}

// ToTokenV3ScopeMap builds a scope from AuthOpts.
func (opts AuthOptsExt) ToTokenV3ScopeMap() (map[string]interface{}, error) {
	b, err := opts.AuthOptionsBuilder.ToTokenV3ScopeMap()
	if err != nil {
		return nil, err
	}

	if opts.TrustID != "" {
		if b == nil {
			b = make(map[string]interface{})
		}
		b["OS-TRUST:trust"] = map[string]interface{}{
			"id": opts.TrustID,
		}
	}

	return b, nil
}

func (opts AuthOptsExt) CanReauth() bool {
	return opts.AuthOptionsBuilder.CanReauth()
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToTrustCreateMap() (map[string]interface{}, error)
}

// CreateOpts provides options used to create a new trust.
type CreateOpts struct {
	// Impersonation allows the trustee to impersonate the trustor.
	Impersonation bool `json:"impersonation"`

	// TrusteeUserID is a user who is capable of consuming the trust.
	TrusteeUserID string `json:"trustee_user_id" required:"true"`

	// TrustorUserID is a user who created the trust.
	TrustorUserID string `json:"trustor_user_id" required:"true"`

	// AllowRedelegation enables redelegation of a trust.
	AllowRedelegation bool `json:"allow_redelegation,omitempty"`

	// ExpiresAt sets expiration time on trust.
	ExpiresAt *time.Time `json:"-"`

	// ProjectID identifies the project.
	ProjectID string `json:"project_id,omitempty"`

	// RedelegationCount specifies a depth of the redelegation chain.
	RedelegationCount int `json:"redelegation_count,omitempty"`

	// RemainingUses specifies how many times a trust can be used to get a token.
	RemainingUses int `json:"remaining_uses,omitempty"`

	// Roles specifies roles that need to be granted to trustee.
	Roles []Role `json:"roles,omitempty"`
}

// ToTrustCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToTrustCreateMap() (map[string]interface{}, error) {
	parent := "trust"

	b, err := gophercloud.BuildRequestBody(opts, parent)
	if err != nil {
		return nil, err
	}

	if opts.ExpiresAt != nil {
		if v, ok := b[parent].(map[string]interface{}); ok {
			v["expires_at"] = opts.ExpiresAt.Format(gophercloud.RFC3339Milli)
		}
	}

	return b, nil
}

type ListOptsBuilder interface {
	ToTrustListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// TrustorUserId filters the response by a trustor user Id.
	TrustorUserId string `q:"trustee_user_id"`

	// TrusteeUserId filters the response by a trustee user Id.
	TrusteeUserId string `q:"trustee_user_id"`
}

// ToTrustListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTrustListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// Create creates a new Trust.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToTrustCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

// Delete deletes a trust.
func Delete(client *gophercloud.ServiceClient, trustID string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, trustID), nil)
	return
}

// List enumerates the Trust to which the current token has access.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToTrustListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return TrustPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details on a single trust, by ID.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}
