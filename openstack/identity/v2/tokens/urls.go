package tokens

import "github.com/gophercloud/gophercloud/v2"

// CreateURL generates the URL used to create new Tokens.
func CreateURL(client gophercloud.Client) string {
	return client.ServiceURL("tokens")
}

// GetURL generates the URL used to Validate Tokens.
func GetURL(client gophercloud.Client, token string) string {
	return client.ServiceURL("tokens", token)
}
