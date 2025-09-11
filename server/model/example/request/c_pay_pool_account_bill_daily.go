package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CPayPoolAccountBillDailySearch struct {
	SearchMap []SearchMap `json:"searchMap"`
	request.PageInfo
	request.RangeTime
	Name         string `json:"name" form:"name"`                 //渠道名称
	PoolName     string `json:"poolName" form:"poolName"`         //支付池名称
	TimeType     int    `json:"timeType" form:"timeType"`         // 1半小时内 2 1小时内 3 当天
	CPayConfigId int    `json:"cPayConfigId" form:"cPayConfigId"` // 渠道id
}

type CPayPoolAccountBillDailyStatistics struct {
	TotalTimes        int64   `json:"totalTimes" gorm:"column:total_times"`
	TotalMoney        float64 `json:"totalMoney" gorm:"column:total_money"`
	TotalRatio        string  `json:"totalRatio" gorm:"column:total_ratio"`
	TotalSuccessCount int64   `json:"totalSuccessCount" gorm:"column:total_success_count"`
}
