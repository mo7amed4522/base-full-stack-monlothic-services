package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CPayPoolAccountSearch struct {
	request.PageInfo
	request.RangeTime
	SearchMap    []SearchMap `json:"searchMap"`
	Type         string      `json:"type" form:"type"`
	Name         string      `json:"name" form:"name"` //名称
	StartAmount  *float64    `json:"startAmount" form:"startAmount"`
	EndAmount    float64     `json:"endAmount" form:"endAmount"`
	CPayPoolName string      `json:"cPayPoolName" form:"cPayPoolName"`
	CardNumber   string      `json:"cardNumber" form:"cardNumber"`
}

type CPayPoolAccountTransactions struct {
	ID     int     `json:"id" validate:"required"`     // 支付池账户id
	Amount float64 `json:"amount" validate:"required"` // 账变金额
	Type   int     `json:"type" validate:"oneof=1 2"`  // 类型 1收款 2出款
	Desc   string  `json:"desc"`                       // 备注
}

type LockReq struct {
	ID   int `json:"id" validate:"required"`
	Type int `json:"type" validate:"oneof=1 2"`
}

type GetCPayPoolAccountListByTypeReq struct {
	request.PageInfo
	Type string `json:"type" form:"type"` // 类型 1 收款账户 2出款账户 3备付金 4 代理 5即是收款也是出款
}

type TransactionsSearch struct {
	request.PageInfo
	request.RangeTime
	SearchMap         []SearchMap `json:"searchMap"`
	Type              int         `json:"type" form:"type"`
	CPayPoolAccountID int         `json:"cPayPoolAccountId" form:"cPayPoolAccountId"`
}

type TransactionsCardSearch struct {
	request.PageInfo
	request.RangeTime
	SearchMap      []SearchMap `json:"searchMap"`
	Type           int         `json:"type" form:"type"`
	CPayPoolCardID int         `json:"cPayPoolCardId" form:"cPayPoolCardId"`
}
