package serviceassets

import "github.com/RainbowMango/huaweicloud-sdk-go"

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
