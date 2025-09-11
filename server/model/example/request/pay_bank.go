package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PayBankSearch struct {
	request.PageInfo
	request.RangeTime
	BankName string `json:"bankName" form:"bankName"`
	BankCode string `json:"bankCode" form:"bankCode"`
}
