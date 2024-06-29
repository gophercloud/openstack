package monitors

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath     = "lbaas"
	resourcePath = "healthmonitors"
)

func rootURL(c gophercloud.Client) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c gophercloud.Client, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
