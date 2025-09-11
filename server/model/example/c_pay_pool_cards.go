package example

import (
	"time"
)

// CPayPoolCards 结构体
type CPayPoolCards struct {
	Id               int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                             //id字段
	Name             string    `json:"name" form:"name"`                                                                       // 账户名称
	Valid            int       `json:"valid" form:"valid" gorm:"column:valid;comment:是否启用，1代表启用，0代表停用;"`                       //是否启用，1代表启用，0代表停用
	Remark           string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                         //备注
	CreateTime       time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                   //创建时间
	UpdateTime       time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`                   //更新时间
	AmountActual     float64   `json:"amountActual" form:"amountActual" gorm:"column:amount_actual;comment:二维码的实际金额;size:17;"` //二维码的实际金额
	CardName         string    `json:"cardName" form:"cardName" gorm:"column:card_name;comment:银行卡人名称;size:200;"`              //银行卡人名称
	CardNumber       string    `json:"cardNumber" form:"cardNumber" gorm:"column:card_number;comment:银行卡人卡号;size:200;"`        //银行卡人卡号
	BankMark         string    `json:"bankMark" form:"bankMark" gorm:"column:bank_mark;comment:银行编码;size:200;"`                //银行编码
	BankName         string    `json:"bankName" form:"bankName" gorm:"column:bank_name;comment:银行名称;size:200;"`                //银行名称
	BankId           int       `json:"bankId" form:"bankId" gorm:"column:bank_id;comment:银行卡id;size:19;"`                      //银行卡id
	CardPhone        string    `json:"cardPhone" form:"cardPhone" gorm:"column:card_phone;comment:银行卡预留电话;size:200;"`          //银行卡预留电话
	DelFlag          int       `json:"delFlag" form:"delFlag" gorm:"column:del_flag;"`                                         //逻辑删除状态1（未删除）和 -1（已删除）默认1
	CreateBy         int       `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`                  //创建人
	UpdateBy         int       `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:创建人;size:19;"`                  //创建人
	UpdateByNickname string    `json:"updateByNickname" form:"updateByNickname" gorm:"-"`
	BindAccountNum   int64     `json:"bindAccountNum" form:"bindAccountNum" gorm:"-"`
	MaxAmount        float64   `json:"maxAmount" form:"maxAmount" gorm:"column:max_amount;comment:最大金额;size:19;"`      //最大金额
	TotalAmount      float64   `json:"totalAmount" form:"totalAmount" gorm:"column:total_amount;comment:总金额;size:19;"` //累计金额
}

// TableName cPayPoolAccount表 CPayPoolCards自定义表名 c_pay_pool_account
func (CPayPoolCards) TableName() string {
	return "c_pay_pool_cards"
}
