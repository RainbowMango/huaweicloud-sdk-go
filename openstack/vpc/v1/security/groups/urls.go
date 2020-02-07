package groups

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

const rootPath = "security-groups"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath)
}

