package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MBillPayoutDailySearch struct {
	request.PageInfo
	SearchMap []SearchMap `json:"searchMap"`
}
