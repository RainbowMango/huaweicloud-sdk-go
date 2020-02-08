package datastores

import "github.com/RainbowMango/huaweicloud-sdk-go"

func listURL(sc *gophercloud.ServiceClient, databasename string) string {
	return sc.ServiceURL("datastores", databasename)
}
