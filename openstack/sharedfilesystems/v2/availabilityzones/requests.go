package availabilityzones

import (
	"github.com/huaweicloud/huaweicloud-sdk-go/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// List will return the existing availability zones.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return AvailabilityZonePage{pagination.SinglePageBase(r)}
	})
}
