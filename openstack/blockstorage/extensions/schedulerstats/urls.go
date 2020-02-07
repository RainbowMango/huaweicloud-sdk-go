package schedulerstats

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func storagePoolsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
