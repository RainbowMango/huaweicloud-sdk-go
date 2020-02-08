package logs

import (
	"github.com/RainbowMango/huaweicloud-sdk-go"
)

func ListURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_activity_log", scalingGroupId)
}
