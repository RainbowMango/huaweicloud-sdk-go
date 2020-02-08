package db_privilege

import (
	"github.com/RainbowMango/huaweicloud-sdk-go"
)


type DbprivilegeResult struct {
	gophercloud.Result
}

type Dbprivilege struct {
	Resp string `json:"resp"`
}

func (r DbprivilegeResult) Extract() (*Dbprivilege, error) {
	var response Dbprivilege
	err := r.ExtractInto(&response)
	return &response, err
}

