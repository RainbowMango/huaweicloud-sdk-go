package availabilityzones

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

// listURL generates URL for list avaliabilityzones
func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
