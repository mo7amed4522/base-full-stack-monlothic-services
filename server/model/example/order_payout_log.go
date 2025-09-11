// 自动生成模板OrderPayout
package example

import (
	"time"
)

// OrderPayoutLog 结构体  OrderPayoutLog
type OrderPayoutLog struct {
	Id                   int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"` //id字段
	CardBank             string    `json:"cardBank" form:"cardBank" gorm:"column:card_bank;comment:银行名称;size:255;"`
	OrderID              int       `json:"orderID" form:"orderID" gorm:"column:order_id"`                       // 订单id
	OrderNo              string    `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:交易号;size:100;"` //交易号
	OldCPayPoolAccountID int       `json:"oldCPayPoolAccountID" form:"oldCPayPoolAccountID" gorm:"column:old_c_pay_pool_account_id"`
	CPayPoolAccountID    int       `json:"cPayPoolAccountID" form:"cPayPoolAccountID" gorm:"column:c_pay_pool_account_id"`
	CType                int       `json:"cType" form:"cType" gorm:"column:c_type;comment:渠道类型，1自研，2三方;"` //渠道类型，1自研，2三方
	Ifsc                 string    `json:"ifsc" form:"ifsc" gorm:"column:ifsc;comment:IFSC;size:100;"`
	ActualAmount         float64   `json:"actualAmount" form:"actualAmount" gorm:"column:actual_amount;comment:实际交易金额;size:17;"`     //实际交易金额
	CardName             string    `json:"cardName" form:"cardName" gorm:"column:card_name;comment:收款人姓名;size:100;"`                 //收款人姓名
	CardNumber           string    `json:"cardNumber" form:"cardNumber" gorm:"column:card_number;comment:收款人卡号;size:100;"`           //收款人卡号
	Status               int       `json:"status" form:"status" gorm:"column:status;comment:出款订单状态 1出款申请 2支付中 3出款成功 -1出款失败 4出款待确认;"` //逻辑删除状态0（未删除）和 1（已删除）默认0
	Remark               string    `json:"remark" form:"remark" gorm:"column:remark;comment:订单备注;size:100;"`                         //订单备注
	OperatorID           int       `json:"operatorID" gorm:"column:operator_id"`                                                     // 操作人
	CreateTime           time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                     //创建时间
	UpdateTime           time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`                     //修改时间
	OperatorName         string    `json:"operatorName" form:"operatorName" gorm:"-"`

	OldCPayPoolAccountName string `json:"oldCPayPoolAccountName" form:"oldCPayPoolAccountName" gorm:"-"`
	CPayPoolAccountName    string `json:"cPayPoolAccountName" form:"cPayPoolAccountName" gorm:"-"`
}

// TableName orderPayout表 OrderPayout自定义表名 order_payout_log
func (OrderPayoutLog) TableName() string {
	return "order_payout_log"
}
