package availabilityzones

import "github.com/RainbowMango/huaweicloud-sdk-go"

// listURL generates URL for list avaliabilityzones
func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
