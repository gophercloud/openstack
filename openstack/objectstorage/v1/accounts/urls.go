package accounts

import "github.com/gophercloud/gophercloud/v2"

func getURL(c gophercloud.Client) string {
	return c.Endpoint
}

func updateURL(c gophercloud.Client) string {
	return getURL(c)
}
