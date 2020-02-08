package publicips

import "github.com/RainbowMango/huaweicloud-sdk-go"

func CreateURL(c *gophercloud.ServiceClient)string{
	return c.ServiceURL("publicips")
}