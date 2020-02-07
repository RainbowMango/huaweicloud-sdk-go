package publicips

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func CreateURL(c *gophercloud.ServiceClient)string{
	return c.ServiceURL("publicips")
}