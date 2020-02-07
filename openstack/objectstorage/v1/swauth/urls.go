package swauth

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func getURL(c *gophercloud.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}
