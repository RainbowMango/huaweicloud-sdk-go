package resetstate

import (
	"github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"
)

func actionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
