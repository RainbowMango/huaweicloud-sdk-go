package schedulerstats

import "github.com/RainbowMango/huaweicloud-sdk-go"

func storagePoolsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
