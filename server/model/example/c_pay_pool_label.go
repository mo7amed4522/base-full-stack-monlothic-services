// 自动生成模板CPayPoolLabel
package example

import (
	"time"
)

// cPayPoolLabel表 结构体  CPayPoolLabel
type CPayPoolLabel struct {
	Id         *int      `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                               //id字段
	CPayPoolId int       `json:"cPayPoolId" form:"cPayPoolId" gorm:"column:c_pay_pool_id;comment:c_pay_pool主键id;size:10;"` //c_pay_pool主键id
	MUserId    *int      `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:m_user主键id;size:19;"`               //m_user主键id
	LabelId    *int      `json:"labelId" form:"labelId" gorm:"column:label_id;comment:sys_label主键id;size:19;"`             //sys_label主键id
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                     //创建时间
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`                     //更新时间
	Remark     *string   `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                           //备注
	DelFalg    *int      `json:"delFalg" form:"delFalg" gorm:"column:del_falg;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`           //逻辑删除状态0（未删除）和 1（已删除）默认0
	CreateBy   *int      `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`                    //创建人
	UpdateBy   *int      `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`                    //修改人

	// 标签名称
	LabelName string `json:"labelName" form:"labelName" gorm:"-"` // 标签名称
	// 支付池名称
	CPayPoolName string `json:"cPayPoolName" form:"cPayPoolName" gorm:"-"` // 支付池名称
}

// TableName cPayPoolLabel表 CPayPoolLabel自定义表名 c_pay_pool_label
func (CPayPoolLabel) TableName() string {
	return "c_pay_pool_label"
}
