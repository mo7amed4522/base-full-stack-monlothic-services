package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type OrderPayoutSearch struct {
	request.PageInfo
	request.RangeTime
	MUserNo             string  `json:"mUserNo" form:"mUserNo"`
	MUserName           string  `json:"mUserName" form:"mUserName"`
	StartAmount         float64 `json:"startAmount" form:"startAmount"`
	EndAmount           float64 `json:"endAmount" form:"endAmount"`
	StartPayTime        string  `json:"startPayTime" form:"startPayTime"`
	EndPayTime          string  `json:"endPayTime" form:"endPayTime"`
	CPayPoolAccountName string  `json:"cPayPoolAccountName" form:"cPayPoolAccountName"`
	CPayPoolName        string  `json:"cPayPoolName" form:"cPayPoolName"`
	FailMessageId       string  `json:"failMessageId" form:"failMessageId"`
}
