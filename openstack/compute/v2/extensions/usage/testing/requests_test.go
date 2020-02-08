package testing

import (
	"testing"

	"github.com/RainbowMango/huaweicloud-sdk-go/openstack/compute/v2/extensions/usage"
	th "github.com/RainbowMango/huaweicloud-sdk-go/testhelper"
	"github.com/RainbowMango/huaweicloud-sdk-go/testhelper/client"
)

func TestGetTenant(t *testing.T) {
	var getOpts usage.SingleTenantOpts
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSingleTenantSuccessfully(t)
	page, err := usage.SingleTenant(client.ServiceClient(), FirstTenantID, getOpts).AllPages()
	th.AssertNoErr(t, err)
	actual, err := usage.ExtractSingleTenant(page)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &SingleTenantUsageResults, actual)
}
