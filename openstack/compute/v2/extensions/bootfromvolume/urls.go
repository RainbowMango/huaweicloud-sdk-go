package bootfromvolume

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-volumes_boot")
}
