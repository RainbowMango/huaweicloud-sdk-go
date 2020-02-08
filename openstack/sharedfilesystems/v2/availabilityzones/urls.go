package availabilityzones

import "github.com/RainbowMango/huaweicloud-sdk-go"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
