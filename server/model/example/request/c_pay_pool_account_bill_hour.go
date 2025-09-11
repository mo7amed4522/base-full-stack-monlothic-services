package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CPayPoolAccountBillHourSearch struct {
	SearchMap []SearchMap `json:"searchMap"`
	request.PageInfo
	request.RangeTime
	Name         string `json:"name" form:"name"`         //渠道名称
	PoolName     string `json:"poolName" form:"poolName"` //支付池名称
	TimeType     int    `json:"timeType" form:"timeType"` // 1半小时内 2 1小时内
	StartHour    *int   `json:"startHour" form:"startHour"`
	EndHour      *int   `json:"endHour" form:"endHour"`
	CPayConfigId int    `json:"cPayConfigId" form:"cPayConfigId"` // 渠道id
}
