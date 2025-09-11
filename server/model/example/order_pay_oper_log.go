// 自动生成模板OrderPayOperLog
package example

import (
	"time"
)

// orderPayOperLog表 结构体  OrderPayOperLog
type OrderPayOperLog struct {
	Id      *int    `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"` //id字段
	OrderId *string `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id
;size:255;"` //订单id

	MUserId    *int       `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:m_user主键id;size:19;"`     //m_user主键id
	Type       *int       `json:"type" form:"type" gorm:"column:type;comment:订单类型：1重制;"`                          //订单类型：1重制
	CreateBy   *int       `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`          //创建人
	Remark     *string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                 //备注
	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`           //创建时间
	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`           //修改时间
	UpdateBy   *int       `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`          //修改人
	DelFlag    *int       `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"` //逻辑删除状态0（未删除）和 1（已删除）默认0
}

// TableName orderPayOperLog表 OrderPayOperLog自定义表名 order_pay_oper_log
func (OrderPayOperLog) TableName() string {
	return "order_pay_oper_log"
}
