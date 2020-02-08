package quotas

import (
	"github.com/RainbowMango/huaweicloud-sdk-go"
)

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}
