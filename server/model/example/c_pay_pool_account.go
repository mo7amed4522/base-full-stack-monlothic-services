// 自动生成模板CPayPoolAccount
package example

import (
	"time"
)

// cPayPoolAccount表 结构体  CPayPoolAccount
type CPayPoolAccount struct {
	Id                  *int      `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                       //id字段
	Name                string    `json:"name"`                                                                                             // 账户名称
	Type                int       `json:"type"`                                                                                             //账户类型 1 收款账户 2出款账户 3备付金 4 代理 6 三方                                                                   //
	MUserId             int       `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:关联m_user表ID;size:19;"`                      //关联m_user表ID
	CPayPoolId          int       `json:"cPayPoolId" form:"cPayPoolId" gorm:"column:c_pay_pool_id;comment:c_pay_pool主键id;size:19;"`         //c_pay_pool主键id
	Qrcode              string    `json:"qrcode" form:"qrcode" gorm:"column:qrcode;comment:扫码的地址;"`                                         //扫码的地址
	LimitAmountMin      *float64  `json:"limitAmountMin" form:"limitAmountMin" gorm:"column:limit_amount_min;comment:单笔交易的最小金额限制;size:17;"` //单笔交易的最小金额限制
	LimitAmountMax      *float64  `json:"limitAmountMax" form:"limitAmountMax" gorm:"column:limit_amount_max;comment:单笔交易的最大金额限制;size:17;"` //单笔交易的最大金额限制
	QrcodeMsg           string    `json:"qrcodeMsg" form:"qrcodeMsg" gorm:"column:qrcode_msg;comment:该二维码的备注信息;size:50;"`                   //该二维码的备注信息
	Valid               *int      `json:"valid" form:"valid" gorm:"column:valid;comment:是否启用，1代表启用，0代表停用;"`                                 //是否启用，1代表启用，0代表停用
	Remark              *string   `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                                   //备注
	CreateTime          time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                             //创建时间
	UpdateTime          time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`                             //更新时间
	QrType              *int      `json:"qrType" form:"qrType" gorm:"column:qr_type;comment:二维码类型，0为二维码，1为赞赏码;"`                            //二维码类型，0为二维码，1为赞赏码
	QrContent           *string   `json:"qrContent" form:"qrContent" gorm:"column:qr_content;comment:二维码内容，Base64加密;"`                      //二维码内容，Base64加密
	AmountActual        float64   `json:"amountActual" form:"amountActual" gorm:"column:amount_actual;comment:二维码的实际金额;size:17;"`           //二维码的实际金额
	CardName            string    `json:"cardName" form:"cardName" gorm:"column:card_name;comment:银行卡人名称;size:200;"`                        //银行卡人名称
	CardNumber          string    `json:"cardNumber" form:"cardNumber" gorm:"column:card_number;comment:银行卡人卡号;size:200;"`                  //银行卡人卡号
	BankMark            string    `json:"bankMark" form:"bankMark" gorm:"column:bank_mark;comment:银行编码;size:200;"`                          //银行编码
	BankName            string    `json:"bankName" form:"bankName" gorm:"column:bank_name;comment:银行名称;size:200;"`                          //银行名称
	BankId              *int      `json:"bankId" form:"bankId" gorm:"column:bank_id;comment:银行卡id;size:19;"`                                //银行卡id
	CardPhone           *string   `json:"cardPhone" form:"cardPhone" gorm:"column:card_phone;comment:银行卡预留电话;size:200;"`                    //银行卡预留电话
	DelFlag             int       `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`                   //逻辑删除状态0（未删除）和 1（已删除）默认0
	CreateBy            *int      `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`                            //创建人
	UpdateBy            *int      `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:创建人;size:19;"`                            //创建人
	UpiId               string    `json:"upiId" form:"upiId" gorm:"column:upi_id;comment:upiId;size:200;"`                                  // upiId
	LimitFlag           int       `json:"limitFlag" form:"limitFlag" gorm:"column:limit_flag;comment:临时取消限制 1关闭 2开启，默认关闭;size:10;"`
	SuccessCountPerDay  int       `json:"successCountPerDay" form:"successCountPerDay" gorm:"column:success_count_per_day;comment:每天成功次数，limit_flag=1生效;size:10;"`
	SuccessAmountPerDay float64   `json:"successAmountPerDay" form:"successAmountPerDay" gorm:"column:success_amount_per_day;comment:每天成功金额，limit_flag=1生效;size:19;"`
	UpdateByNickname    string    `json:"updateByNickname" form:"updateByNickname" gorm:"-"`
	CPayPoolCardsId     int       `json:"cPayPoolCardsId" form:"cPayPoolCardsId" gorm:"column:c_pay_pool_cards_id;comment:c_pay_pool_cards主键id;size:19;"` //c_pay_pool_cards_id主键id
	CPayPoolCardsName   string    `json:"cPayPoolCardsName" form:"cPayPoolCardsName" gorm:"-"`
}

// TableName cPayPoolAccount表 CPayPoolAccount自定义表名 c_pay_pool_account
func (CPayPoolAccount) TableName() string {
	return "c_pay_pool_account"
}
