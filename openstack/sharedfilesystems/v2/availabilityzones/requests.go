package availabilityzones

import (
	"github.com/RainbowMango/huaweicloud-sdk-go"
	"github.com/RainbowMango/huaweicloud-sdk-go/pagination"
)

// List will return the existing availability zones.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return AvailabilityZonePage{pagination.SinglePageBase(r)}
	})
}
