package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type OrderPayRequestSearch struct {
	request.PageInfo
	request.RangeTime
	LogStatus int `json:"logStatus" form:"logStatus"` // 1异常 2成功
}
