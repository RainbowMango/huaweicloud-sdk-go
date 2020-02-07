package quotas

import (
	"github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"
)

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}

func ListWithInstancesURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("quotas", scalingGroupId)
}
