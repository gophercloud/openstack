package resourceproviders

import "github.com/gophercloud/gophercloud"

const (
	apiName = "resource_providers"
)

func resourceProvidersListURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiName)
}

func getResourceProviderUsagesURL(client *gophercloud.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID, "usages")
}
