package configurations

import "github.com/RainbowMango/huaweicloud-sdk-go"

func listURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("configurations")
}

func createURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("configurations")
}
