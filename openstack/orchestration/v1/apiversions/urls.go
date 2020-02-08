package apiversions

import "github.com/RainbowMango/huaweicloud-sdk-go"

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}
