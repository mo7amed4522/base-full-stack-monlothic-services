package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MBillPayDailySearch struct {
	request.PageInfo
	request.RangeTime
	SearchMap []SearchMap `json:"searchMap"`
}
