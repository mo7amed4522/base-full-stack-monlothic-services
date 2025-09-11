package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type OrderPaySearch struct {
	request.PageInfo
	request.RangeTime
	SearchMap           []SearchMap `json:"searchMap"`
	UtrNo               string      `json:"utrNo" form:"utrNo"`
	PayType             string      `json:"payType" form:"payType"`
	MUserNo             string      `json:"mUserNo" form:"mUserNo"`
	MUserName           string      `json:"mUserName" form:"mUserName"`
	StartAmount         float64     `json:"startAmount" form:"startAmount"`
	EndAmount           float64     `json:"endAmount" form:"endAmount"`
	StartPayTime        string      `json:"startPayTime" form:"startPayTime"`
	EndPayTime          string      `json:"endPayTime" form:"endPayTime"`
	CPayPoolAccountName string      `json:"cPayPoolAccountName" form:"cPayPoolAccountName"`
	CPayPoolName        string      `json:"cPayPoolName" form:"cPayPoolName"`
	UtrStatus           int         `json:"utrStatus" form:"utrStatus"`
	CardNumber          string      `json:"cardNumber" form:"cardNumber"`
}
