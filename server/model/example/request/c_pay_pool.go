package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CPayPoolSearch struct {
	request.PageInfo
	request.RangeTime
	CPayPoolName string `json:"cPayPoolName" form:"cPayPoolName"`
	CPayPoolCode string `json:"cPayPoolCode" form:"cPayPoolCode"`
}
