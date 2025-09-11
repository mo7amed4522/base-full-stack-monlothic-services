// 自动生成模板CPayPoolAccountBillHour
package example

import (
	"time"
)

// cPayPoolAccountBillHour表 结构体  CPayPoolAccountBillHour
type CPayPoolAccountBillHour struct {
	Id                *int      `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                                             //id字段
	CPayConfigId      int       `json:"cPayConfigId" form:"cPayConfigId" gorm:"column:c_pay_config_id;comment:c_pay_config主键id;size:19;"`                       //c_pay_config主键id
	CPayPoolId        *int      `json:"cPayPoolId" form:"cPayPoolId" gorm:"column:c_pay_pool_id;comment:c_pay_pool主键id;size:19;"`                               //c_pay_pool主键id
	CPayPoolAccountId int       `json:"cPayPoolAccountId" form:"cPayPoolAccountId" gorm:"column:c_pay_pool_account_id;comment:c_pay_pool_account主键ID;size:19;"` //c_pay_pool_account主键ID
	Date              string    `json:"date" form:"date" gorm:"column:date;comment:具体的某一天;size:10;"`                                                            //具体的某一天
	Hour              string    `json:"hour" form:"hour" gorm:"column:hour;comment:具体的某一小时0-23;size:10;"`                                                       //具体的某小时
	SuccessCountDay   int       `json:"successCountDay" form:"successCountDay" gorm:"column:success_count_day;comment:每天完成次数;size:10;"`                         //每天完成次数
	SuccessAmountDay  float64   `json:"successAmountDay" form:"successAmountDay" gorm:"column:success_amount_day;comment:每天已经完成的金额;size:17;"`                   //每天已经完成的金额
	RequestCountDay   int       `json:"requestCountDay" form:"requestCountDay" gorm:"column:request_count_day;comment:当天调用次数;size:10;"`                         //当天调用次数
	CreateTime        time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                                                   //创建时间
	UpdateTime        time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`                                                   //更新时间
	DelFlag           *int      `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`                                         //逻辑删除状态0（未删除）和 1（已删除）默认0
	CreateBy          *int      `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`                                                  //创建人
	UpdateBy          *int      `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:创建人;size:19;"`                                                  //创建人
	Remark            *string   `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                                                         //备注

	// 支付池名称
	CPayPoolName string `json:"cPayPoolName" form:"cPayPoolName" gorm:"-"` // 支付池名称
	// 渠道名称
	CPayConfigName string `json:"cPayConfigName" form:"cPayConfigName" gorm:"-"` // 渠道名称
	// 收款方式名称
	CPayPoolAccountName string `json:"cPayPoolAccountName" form:"cPayPoolAccountName" gorm:"-"`
	// 收款方式启用状态
	CPayPoolAccountValid int `json:"cPayPoolAccountValid" form:"cPayPoolAccountValid" gorm:"-"`
	// 成功比例   SuccessCountDay 除 RequestCountDay
	SuccessRatio string `json:"successRatio" form:"successRatio" gorm:"-"`
}

// TableName cPayPoolAccountBillHour表 CPayPoolAccountBillHour自定义表名 c_pay_pool_account_bill_hour
func (CPayPoolAccountBillHour) TableName() string {
	return "c_pay_pool_account_bill_hour"
}
