package common

import (
	"github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc
}
