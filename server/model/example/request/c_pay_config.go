package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CPayConfigSearch struct {
	SearchMap []SearchMap `json:"searchMap"`
	request.PageInfo
	request.RangeTime
	CPayConfigName string `json:"cPayConfigName" form:"cPayConfigName"`
	CPayConfigCode string `json:"cPayConfigCode" form:"cPayConfigCode"`
	ChannelType    int    `json:"channelType" form:"channelType"` //来源类型 1自研 2三方
	Type           int    `json:"type" form:"type"`               //渠道类型 1支付 2出款
	PayType        string `json:"payType" form:"payType"`         //支付方式：1二维码 2银行卡
}
