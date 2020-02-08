package quotas

import (
	"github.com/RainbowMango/huaweicloud-sdk-go"
)

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}
