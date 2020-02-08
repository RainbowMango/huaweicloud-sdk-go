package events

import (
	"github.com/RainbowMango/huaweicloud-sdk-go"
)

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("events")
}
