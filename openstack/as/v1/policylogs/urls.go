package policylogs

import (
	"github.com/RainbowMango/huaweicloud-sdk-go"
)

func ListURL(c *gophercloud.ServiceClient, scalingPolicyId string) string {
	return c.ServiceURL("scaling_policy_execute_log", scalingPolicyId)
}
