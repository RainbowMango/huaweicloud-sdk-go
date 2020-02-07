package accounts

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}

func updateURL(c *gophercloud.ServiceClient) string {
	return getURL(c)
}
