// 自动生成模板CPayPoolBillDaily
package example

import (
	"time"
)

// cPayPoolBillDaily表 结构体  CPayPoolBillDaily
type CPayPoolBillDaily struct {
	Id               *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                           //id字段
	CPayConfigId     *int       `json:"cPayConfigId" form:"cPayConfigId" gorm:"column:c_pay_config_id;comment:c_pay_config主键id;size:19;"`     //c_pay_config主键id
	CPayPoolId       *int       `json:"cPayPoolId" form:"cPayPoolId" gorm:"column:c_pay_pool_id;comment:c_pay_pool主键id;size:19;"`             //c_pay_pool主键id
	Date             *string    `json:"date" form:"date" gorm:"column:date;comment:具体的某一天;size:10;"`                                          //具体的某一天
	SuccessCountDay  *int       `json:"successCountDay" form:"successCountDay" gorm:"column:success_count_day;comment:每天完成次数;size:10;"`       //每天完成次数
	SuccessAmountDay *float64   `json:"successAmountDay" form:"successAmountDay" gorm:"column:success_amount_day;comment:每天已经完成的金额;size:17;"` //每天已经完成的金额
	CreateTime       *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                                 //创建时间
	UpdateTime       *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`                                 //更新时间
	CreateBy         *int       `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`                                //创建人
	UpdateBy         *int       `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`                                //修改人
	Remark           *string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                                       //备注
	DelFlag          *int       `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`                       //逻辑删除状态0（未删除）和 1（已删除）默认0
}

// TableName cPayPoolBillDaily表 CPayPoolBillDaily自定义表名 c_pay_pool_bill_daily
func (CPayPoolBillDaily) TableName() string {
	return "c_pay_pool_bill_daily"
}
