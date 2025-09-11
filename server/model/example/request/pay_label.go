package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PayLabelSearch struct {
	request.PageInfo
	request.RangeTime
	LabelName string `json:"labelName" form:"labelName"`
	LabelCode string `json:"labelCode" form:"labelCode"`
}
