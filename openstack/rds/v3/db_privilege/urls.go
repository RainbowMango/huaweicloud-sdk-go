package db_privilege

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

func createURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "db_privilege")
}

func deleteURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "db_privilege")
}
