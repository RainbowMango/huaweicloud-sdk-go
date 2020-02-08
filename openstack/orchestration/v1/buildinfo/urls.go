package buildinfo

import "github.com/RainbowMango/huaweicloud-sdk-go"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
