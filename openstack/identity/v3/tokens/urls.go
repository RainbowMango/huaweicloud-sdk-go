package tokens

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
