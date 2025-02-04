package listeners

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath       = "lbaas"
	resourcePath   = "listeners"
	statisticsPath = "stats"
)

func rootURL(c gophercloud.Client) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c gophercloud.Client, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

func statisticsRootURL(c gophercloud.Client, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, statisticsPath)
}
