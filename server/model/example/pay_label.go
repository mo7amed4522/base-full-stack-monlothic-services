// 自动生成模板PayLabel
package example

import (
	"time"
)

// payLabel表 结构体  PayLabel
type PayLabel struct {
	Id         *int      `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                     //id字段
	Name       *string   `json:"name" form:"name" gorm:"column:name;comment:标记名称;size:200;"`                     //标记名称
	Code       *string   `json:"code" form:"code" gorm:"column:code;comment:标记值,需唯一;size:100;"`                  //标记值,需唯一
	Valid      *int      `json:"valid" form:"valid" gorm:"column:valid;comment:是否启用，1代表启用，0代表停用;"`               //是否启用，1代表启用，0代表停用
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`               //createTime字段
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`               //updateTime字段
	Remark     *string   `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                 //备注
	DelFlag    *int      `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"` //逻辑删除状态0（未删除）和 1（已删除）默认0
	CreateBy   *int      `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`          //创建人
	UpdateBy   *int      `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`          //修改人
}

// TableName payLabel表 PayLabel自定义表名 pay_label
func (PayLabel) TableName() string {
	return "pay_label"
}
