package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type OrderPayReversalSearch struct {
	request.PageInfo
	request.RangeTime
	UTR                 string  `json:"utr" form:"utr"`
	CardNumber          string  `json:"cardNumber" form:"cardNumber"`
	StartAmount         float64 `json:"startAmount" form:"startAmount"`
	EndAmount           float64 `json:"endAmount" form:"endAmount"`
	SourceType          int     `json:"sourceType" form:"sourceType"`
	CPayPoolAccountName string  `json:"cPayPoolAccountName" form:"cPayPoolAccountName" gorm:"-"`
	OperatorID          int     `json:"operatorID" form:"operatorID"`
	OprType             int     `json:"oprType" form:"oprType"`
}

type LogListReq struct {
	request.PageInfo
	request.RangeTime
	OperateID int `json:"operateID" form:"operateID"`
	Classes   int `json:"classes" form:"classes"` // 班次 1白班 2夜班
}
