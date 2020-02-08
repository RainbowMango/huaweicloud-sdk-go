package notifications

import (
	"github.com/RainbowMango/huaweicloud-sdk-go"
)

func DeleteURL(c *gophercloud.ServiceClient, scalingGroupId string, topicUrn string) string {
	return c.ServiceURL("scaling_notification", scalingGroupId, topicUrn)
}

func EnableURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_notification", scalingGroupId)
}

func ListURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_notification", scalingGroupId)
}
