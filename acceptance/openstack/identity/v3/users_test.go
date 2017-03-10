// +build acceptance

package v3

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
)

func TestUsersList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	var iTrue bool = true
	listOpts := users.ListOpts{
		Enabled: &iTrue,
	}

	allPages, err := users.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to list users: %v", err)
	}

	allUsers, err := users.ExtractUsers(allPages)
	if err != nil {
		t.Fatalf("Unable to extract users: %v", err)
	}

	for _, user := range allUsers {
		tools.PrintResource(t, user)
	}
}

func TestUsersGet(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	allPages, err := users.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list users: %v", err)
	}

	allUsers, err := users.ExtractUsers(allPages)
	if err != nil {
		t.Fatalf("Unable to extract users: %v", err)
	}

	user := allUsers[0]
	p, err := users.Get(client, user.ID, nil).Extract()
	if err != nil {
		t.Fatalf("Unable to get user: %v", err)
	}

	tools.PrintResource(t, p)
}

func TestUserCRUD(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	project, err := CreateProject(t, client, nil)
	if err != nil {
		t.Fatalf("Unable to create project: %v", err)
	}
	defer DeleteProject(t, client, project.ID)

	tools.PrintResource(t, project)

	createOpts := users.CreateOpts{
		DefaultProjectID: project.ID,
		Password:         "foobar",
		DomainID:         "default",
	}

	user, err := CreateUser(t, client, &createOpts)
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}

	tools.PrintResource(t, user)
}
