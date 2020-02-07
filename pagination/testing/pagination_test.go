package testing

import (
	"github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"
	"github.com/gophercloud/gophercloud/testhelper"
)

func createClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
