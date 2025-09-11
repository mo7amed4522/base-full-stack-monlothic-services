package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type OrderPayUserSearch struct {
	request.PageInfo
	request.RangeTime
	MUserNo    string `json:"mUserNo" form:"mUserNo"`
	SubUserId  string `json:"subUserId" form:"subUserId"`
	SubUserMsg string `json:"subUserMsg" form:"subUserMsg"`
}
