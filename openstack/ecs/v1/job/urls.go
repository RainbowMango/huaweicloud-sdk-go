package job

import "github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"

// Querying task statuses URL
func jobURL(sc *gophercloud.ServiceClient, jobId string) string {
	return sc.ServiceURL("jobs", jobId)
}