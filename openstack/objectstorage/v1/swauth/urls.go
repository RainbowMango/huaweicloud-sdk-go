package swauth

import "github.com/RainbowMango/huaweicloud-sdk-go"

func getURL(c *gophercloud.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}
