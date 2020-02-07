package apiversions

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}
