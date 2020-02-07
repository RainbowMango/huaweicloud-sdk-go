package buildinfo

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
