package testing

import (
	"context"
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	fake "github.com/gophercloud/gophercloud/v2/openstack/networking/v2/common"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/lbaas_v2/loadbalancers"
	"github.com/gophercloud/gophercloud/v2/pagination"
	th "github.com/gophercloud/gophercloud/v2/testhelper"
)

func TestListLoadbalancers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLoadbalancerListSuccessfully(t)

	pages := 0
	err := loadbalancers.List(fake.ServiceClient(), loadbalancers.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := loadbalancers.ExtractLoadBalancers(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 loadbalancers, got %d", len(actual))
		}
		th.CheckDeepEquals(t, LoadbalancerWeb, actual[0])
		th.CheckDeepEquals(t, LoadbalancerDb, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllLoadbalancers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLoadbalancerListSuccessfully(t)

	allPages, err := loadbalancers.List(fake.ServiceClient(), loadbalancers.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := loadbalancers.ExtractLoadBalancers(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, LoadbalancerWeb, actual[0])
	th.CheckDeepEquals(t, LoadbalancerDb, actual[1])
}

func TestCreateLoadbalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLoadbalancerCreationSuccessfully(t, SingleLoadbalancerBody)

	actual, err := loadbalancers.Create(context.TODO(), fake.ServiceClient(), loadbalancers.CreateOpts{
		Name:         "db_lb",
		AdminStateUp: gophercloud.Enabled,
		VipSubnetID:  "9cedb85d-0759-4898-8a4b-fa5a5ea10086",
		VipAddress:   "10.30.176.48",
		FlavorID:     "bba40eb2-ee8c-11e9-81b4-2a2ae2dbcce4",
		Provider:     "haproxy",
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, LoadbalancerDb, *actual)
}

func TestRequiredCreateOpts(t *testing.T) {
	res := loadbalancers.Create(context.TODO(), fake.ServiceClient(), loadbalancers.CreateOpts{})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
	res = loadbalancers.Create(context.TODO(), fake.ServiceClient(), loadbalancers.CreateOpts{Name: "foo"})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
	res = loadbalancers.Create(context.TODO(), fake.ServiceClient(), loadbalancers.CreateOpts{Name: "foo", Description: "bar"})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
	res = loadbalancers.Create(context.TODO(), fake.ServiceClient(), loadbalancers.CreateOpts{Name: "foo", Description: "bar", VipAddress: "bar"})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
}

func TestGetLoadbalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLoadbalancerGetSuccessfully(t)

	client := fake.ServiceClient()
	actual, err := loadbalancers.Get(context.TODO(), client, "36e08a3e-a78f-4b40-a229-1e7e23eee1ab").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, LoadbalancerDb, *actual)
}

func TestGetLoadbalancerStatusesTree(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLoadbalancerGetStatusesTree(t)

	client := fake.ServiceClient()
	actual, err := loadbalancers.GetStatuses(context.TODO(), client, "36e08a3e-a78f-4b40-a229-1e7e23eee1ab").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, LoadbalancerStatusesTree, *actual)
}

func TestDeleteLoadbalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLoadbalancerDeletionSuccessfully(t)

	res := loadbalancers.Delete(context.TODO(), fake.ServiceClient(), "36e08a3e-a78f-4b40-a229-1e7e23eee1ab")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateLoadbalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLoadbalancerUpdateSuccessfully(t)

	client := fake.ServiceClient()
	name := "NewLoadbalancerName"
	actual, err := loadbalancers.Update(context.TODO(), client, "36e08a3e-a78f-4b40-a229-1e7e23eee1ab", loadbalancers.UpdateOpts{
		Name: &name,
	}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}

	th.CheckDeepEquals(t, LoadbalancerUpdated, *actual)
}

func TestCascadingDeleteLoadbalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLoadbalancerDeletionSuccessfully(t)

	sc := fake.ServiceClient()
	sc.Type = "network"
	err := loadbalancers.CascadingDelete(context.TODO(), sc, "36e08a3e-a78f-4b40-a229-1e7e23eee1ab").ExtractErr()
	if err == nil {
		t.Fatalf("expected error running CascadingDelete with Neutron service client but didn't get one")
	}

	sc.Type = "load-balancer"
	err = loadbalancers.CascadingDelete(context.TODO(), sc, "36e08a3e-a78f-4b40-a229-1e7e23eee1ab").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestGetLoadbalancerStatsTree(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLoadbalancerGetStatsTree(t)

	client := fake.ServiceClient()
	actual, err := loadbalancers.GetStats(context.TODO(), client, "36e08a3e-a78f-4b40-a229-1e7e23eee1ab").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, LoadbalancerStatsTree, *actual)
}
