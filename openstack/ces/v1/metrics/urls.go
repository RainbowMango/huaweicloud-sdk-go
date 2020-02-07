package metrics

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func getMetricsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("metrics")
}
