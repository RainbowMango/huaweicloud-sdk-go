package flavors

import "github.com/RainbowMango/huaweicloud-sdk-go"

func listURL(sc *gophercloud.ServiceClient, databasename string) string {
	return sc.ServiceURL("flavors", databasename)
}
