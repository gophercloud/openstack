package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/shares"
	"github.com/gophercloud/gophercloud/pagination"
)

func TestShareCreate(t *testing.T) {
	client, err := clients.NewSharedFileSystemV2Client()
	if err != nil {
		t.Fatalf("Unable to create a sharedfs client: %v", err)
	}

	share, err := CreateShare(t, client, "my test share")
	if err != nil {
		t.Fatalf("Unable to create a share: %v", err)
	}

	defer DeleteShare(t, client, share)

	created, err := shares.Get(client, share.ID).Extract()
	if err != nil {
		t.Errorf("Unable to retrieve share: %v", err)
	}
	PrintShare(t, created)
}

func TestShareListShort(t *testing.T) {
	client, err := clients.NewSharedFileSystemV2Client()
	if err != nil {
		t.Fatalf("Unable to create a sharedfs client: %v", err)
	}

	names := []string{"share_one", "share_two", "share_three"}

	created, err := CreateShares(t, client, names)
	if err != nil {
		t.Fatalf("Unable to create a share: %v", err)
	}

	defer DeleteShares(t, client, created)

	pages, err := shares.List(client, &shares.ListOpts{}, false).AllPages()
	if err != nil {
		t.Fatalf("Failed to list shares: %v", err)
	}

	shares, err := shares.ExtractShares(pages)
	if err != nil {
		t.Fatalf("Unable to extract shares: %v", err)
	}

	for _, share := range shares {
		PrintShare(t, &share)
	}
}

func TestShareListDetail(t *testing.T) {
	client, err := clients.NewSharedFileSystemV2Client()
	if err != nil {
		t.Fatalf("Unable to create a sharedfs client: %v", err)
	}

	names := []string{"share_one", "share_two", "share_three"}

	created, err := CreateShares(t, client, names)
	if err != nil {
		t.Fatalf("Unable to create a share: %v", err)
	}

	defer DeleteShares(t, client, created)

	pages, err := shares.List(client, &shares.ListOpts{}, true).AllPages()
	if err != nil {
		t.Fatalf("Failed to list shares: %v", err)
	}

	shares, err := shares.ExtractShares(pages)
	if err != nil {
		t.Fatalf("Unable to extract shares: %v", err)
	}

	for _, share := range shares {
		PrintShare(t, &share)
	}
}

func TestShareListDetailPaginate(t *testing.T) {
	client, err := clients.NewSharedFileSystemV2Client()
	if err != nil {
		t.Fatalf("Unable to create a sharedfs client: %v", err)
	}

	names := []string{"share_one", "share_two", "share_three"}

	created, err := CreateShares(t, client, names)
	if err != nil {
		t.Fatalf("Unable to create a share: %v", err)
	}

	defer DeleteShares(t, client, created)

	err = shares.List(client, &shares.ListOpts{}, true).EachPage(func(page pagination.Page) (bool, error) {
		l, _ := shares.ExtractShares(page)
		t.Logf("Got some shares: %+v\n", l)
		return true, nil
	})

}
