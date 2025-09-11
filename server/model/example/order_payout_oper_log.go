// 自动生成模板OrderPayoutOperLog
package example

import (
	"time"
)

// orderPayoutOperLog表 结构体  OrderPayoutOperLog
type OrderPayoutOperLog struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                  //id字段
	OrderNo    *string    `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;size:255;"`         //订单号
	MUserNo    *string    `json:"mUserNo" form:"mUserNo" gorm:"column:m_user_no;comment:商户号;size:255;"`        //商户号
	MUserName  *string    `json:"mUserName" form:"mUserName" gorm:"column:m_user_name;comment:商户名称;size:255;"` //商户名称
	OpType     *int       `json:"opType" form:"opType" gorm:"column:op_type;comment:订单类型：1重制;"`                //订单类型：1重制
	SmId       *int       `json:"smId" form:"smId" gorm:"column:sm_id;comment:系统操作人员id;size:19;"`              //系统操作人员id
	Remark     *string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`              //备注
	Msg1       *string    `json:"msg1" form:"msg1" gorm:"column:msg1;comment:操作说明：该订单的说明;"`                    //操作说明：该订单的说明
	Msg2       *string    `json:"msg2" form:"msg2" gorm:"column:msg2;comment:该订单的补充说明;"`                       //该订单的补充说明
	LogLevel   *int       `json:"logLevel" form:"logLevel" gorm:"column:log_level;comment:日志级别：1正常2错误3有潜在错误;"` //日志级别：1正常2错误3有潜在错误
	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`        //创建时间
	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`        //修改时间
	CreateBy   *int       `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`       //创建人
	UpdateBy   *int       `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`       //修改人
}

// TableName orderPayoutOperLog表 OrderPayoutOperLog自定义表名 order_payout_oper_log
func (OrderPayoutOperLog) TableName() string {
	return "order_payout_oper_log"
}
