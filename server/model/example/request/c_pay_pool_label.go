package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CPayPoolLabelSearch struct {
	request.PageInfo
	request.RangeTime
	CPayPoolName string `json:"cPayPoolName" form:"cPayPoolName"`
	LabelName    string `json:"labelName" form:"labelName"`
}
