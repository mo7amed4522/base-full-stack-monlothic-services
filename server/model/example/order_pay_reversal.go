// 自动生成模板OrderPayReversal
package example

import (
	"time"
)

// orderPayReversal表 结构体  OrderPayReversal
type OrderPayReversal struct {
	Id           *int    `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                       //id字段
	CPayConfigId *int    `json:"cPayConfigId" form:"cPayConfigId" gorm:"column:c_pay_config_id;comment:c_pay_config主键id;size:19;"` //c_pay_config主键id
	CPayPoolId   *int    `json:"cPayPoolId" form:"cPayPoolId" gorm:"column:c_pay_pool_id;comment:c_pay_pool主键id;size:19;"`         //c_pay_pool主键id
	Utr          *string `json:"utr" form:"utr" gorm:"column:utr;comment:utr编码;size:100;"`                                         //utr编码

	CardNumber        *string    `json:"cardNumber" form:"cardNumber" gorm:"column:card_number;comment:银行卡号;size:100;"`                                          //银行卡号
	Status            int        `json:"status" form:"status" gorm:"column:status;comment:状态 1处理中 2成功  3超时  -1失败 -2已关闭;"`                                        //状态 0处理中 1成功  2超时  -1失败 -2已关闭
	Amount            *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:金额;size:17;"`                                                          //金额
	CPayPoolAccountId *int       `json:"cPayPoolAccountId" form:"cPayPoolAccountId" gorm:"column:c_pay_pool_account_id;comment:c_pay_pool_account主键ID;size:19;"` //c_pay_pool_account主键ID
	PayOrderId        *int       `json:"payOrderId" form:"payOrderId" gorm:"column:pay_order_id;comment:pay_order表主键ID;size:19;"`                                //pay_order表主键ID
	OrderNo           *string    `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:三方订单号;size:100;"`                                                  //三方订单号
	SourceType        *int       `json:"sourceType" form:"sourceType" gorm:"column:source_type;comment:支付来源  1银行转账  2upi转账;"`                                    //支付来源  1银行转账  2upi转账
	OprType           *int       `json:"oprType" form:"oprType" gorm:"column:opr_type;comment:操作类型：1API操作 2支付后台操作;"`                                             //操作类型：1API操作 2支付后台操作
	CreateTime        *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:添加时间;"`                                                   //添加时间
	CreateBy          *int       `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;"`                                                          //创建人
	UpdateTime        *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`                                                   //修改时间
	DelFlag           *int       `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:是否删除   0未删除  1删除;"`                                                //是否删除   0未删除  1删除
	Remark            string     `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                                                         //备注
	UpdateBy          *int       `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`                                                  //修改人
	CreateByNickname  string     `json:"createByNickname" form:"createByNickname" gorm:"-"`
	UpdateByNickname  string     `json:"updateByNickname" form:"updateByNickname" gorm:"-"`
	// 收款方式名称
	CPayPoolAccountName string `json:"cPayPoolAccountName" form:"cPayPoolAccountName" gorm:"column:cPayPoolAccountName"`
}

// TableName orderPayReversal表 OrderPayReversal自定义表名 order_pay_reversal
func (OrderPayReversal) TableName() string {
	return "order_pay_reversal"
}
