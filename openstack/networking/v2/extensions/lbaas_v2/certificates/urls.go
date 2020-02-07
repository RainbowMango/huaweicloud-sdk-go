package certificates

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

const (
	rootPath     = "lbaas"
	resourcePath = "certificates"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
