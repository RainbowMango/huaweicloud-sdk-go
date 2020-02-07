package flavors

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func listURL(sc *gophercloud.ServiceClient, databasename string) string {
	return sc.ServiceURL("flavors", databasename)
}
