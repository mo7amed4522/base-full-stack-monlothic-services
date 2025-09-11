package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MUserSearch struct {
	request.PageInfo
	request.RangeTime
	MUserName string `json:"mUserName" form:"mUserName"`
	MUserNo   string `json:"mUserNo" form:"mUserNo"`
}
