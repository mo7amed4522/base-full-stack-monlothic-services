package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CPayPoolCardsSearch struct {
	request.PageInfo
	request.RangeTime
	Name        string   `json:"name" form:"name"` //名称
	StartAmount *float64 `json:"startAmount" form:"startAmount"`
	EndAmount   float64  `json:"endAmount" form:"endAmount"`
	CardNumber  string   `json:"cardNumber" form:"cardNumber"`
}
