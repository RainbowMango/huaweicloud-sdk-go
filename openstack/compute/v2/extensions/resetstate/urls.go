package resetstate

import (
	"github.com/RainbowMango/huaweicloud-sdk-go"
)

func actionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
