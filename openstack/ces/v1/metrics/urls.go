package metrics

import "github.com/RainbowMango/huaweicloud-sdk-go"

func getMetricsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("metrics")
}
