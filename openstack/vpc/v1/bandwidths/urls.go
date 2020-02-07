package bandwidths

import (
	"github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"
)

func GetURL(c *gophercloud.ServiceClient, bandwidthId string) string {
	return c.ServiceURL("bandwidths", bandwidthId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("bandwidths")
}

func UpdateURL(c *gophercloud.ServiceClient, bandwidthId string) string {
	return c.ServiceURL("bandwidths", bandwidthId)
}
