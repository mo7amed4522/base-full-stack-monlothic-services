package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MPayChannelConfigSearch struct {
	request.PageInfo
	request.RangeTime
	CPayConfigName string `json:"cPayConfigName" form:"cPayConfigName"`
	MUserName      string `json:"mUserName" form:"mUserName"`
}
