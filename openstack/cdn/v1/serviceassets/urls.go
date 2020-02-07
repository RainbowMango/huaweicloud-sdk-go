package serviceassets

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
