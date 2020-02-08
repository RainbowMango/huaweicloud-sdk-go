package bootfromvolume

import "github.com/RainbowMango/huaweicloud-sdk-go"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-volumes_boot")
}
