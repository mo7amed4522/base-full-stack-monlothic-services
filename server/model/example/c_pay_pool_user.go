// 自动生成模板CPayPoolUser
package example

import (
	"time"
)

// cPayPoolUser表 结构体  CPayPoolUser
type CPayPoolUser struct {
	Id         *int      `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键;size:19;"`                             //主键
	MUserId    *int      `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:m_user主键id;size:19;"`               //m_user主键id
	CPayPoolId *int      `json:"cPayPoolId" form:"cPayPoolId" gorm:"column:c_pay_pool_id;comment:c_pay_pool主键id;size:19;"` //c_pay_pool主键id
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                     //创建时间
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`                     //更新时间
	IsPrivate  *int      `json:"isPrivate" form:"isPrivate" gorm:"column:is_private;comment:是否专用，0为通用，1为专用;"`              //是否专用，0为通用，1为专用
	DelFlag    *int      `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`           //逻辑删除状态0（未删除）和 1（已删除）默认0
	CreateBy   *int      `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`                    //创建人
	UpdateBy   *int      `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`                    //修改人
	Remark     *string   `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                           //备注

	// 商户名称
	MUserName string `json:"mUserName" form:"mUserName"` // 商户名称

	// 支付池名称
	CPayPoolName string `json:"cPayPoolName" form:"cPayPoolName"` // 支付池名称
}

// TableName cPayPoolUser表 CPayPoolUser自定义表名 c_pay_pool_user
func (CPayPoolUser) TableName() string {
	return "c_pay_pool_user"
}
