package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MUserLabelSearch struct {
	request.PageInfo
	request.RangeTime
	SearchMap []SearchMap `json:"searchMap"`
	MUserName string      `json:"mUserName" form:"mUserName"`
	LabelName string      `json:"labelName" form:"labelName"`
}
