package job

import "github.com/RainbowMango/huaweicloud-sdk-go"

// Querying task statuses URL
func jobURL(sc *gophercloud.ServiceClient, jobId string) string {
	return sc.ServiceURL("jobs", jobId)
}