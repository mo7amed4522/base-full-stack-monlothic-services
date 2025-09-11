// 自动生成模板MUserLabel
package example

import (
	"time"
)

// mUserLabel表 结构体  MUserLabel
type MUserLabel struct {
	Id         int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                     //id字段
	MUserId    int       `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:m_user主键id;size:19;"`     //m_user主键id
	LabelId    int       `json:"labelId" form:"labelId" gorm:"column:label_id;comment:sys_lable主键ID;size:19;"`   //sys_lable主键ID
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`           //创建时间
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`           //更新时间
	Remark     *string   `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                 //备注
	DelFalg    *int      `json:"delFalg" form:"delFalg" gorm:"column:del_falg;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"` //逻辑删除状态0（未删除）和 1（已删除）默认0
	CreateBy   int       `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`          //创建人
	UpdateBy   int       `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`          //修改人

	// 商户名称
	MUserName string `json:"mUserName" form:"mUserName" gorm:"-"` // 商户名称
	// 标签名称
	LabelName        string `json:"labelName" form:"labelName" gorm:"column:label_name"` // 标签名称
	UpdateByNickname string `json:"updateByNickname" form:"updateByNickname" gorm:"-"`
}

// TableName mUserLabel表 MUserLabel自定义表名 m_user_label
func (MUserLabel) TableName() string {
	return "m_user_label"
}
