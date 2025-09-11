// 自动生成模板CDominManage
package manage

import (
	"time"
)

// cDominManage表 结构体  CDominManage
type CDominManage struct {
	Id            *int              `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键ID;size:20;"`                         //主键ID
	Domain        *string           `json:"domain" form:"domain" gorm:"column:domain;comment:域名地址;size:255;"`                       //域名地址
	Status        *int              `json:"status" form:"status" gorm:"column:status;comment:状态（1匹配中，2正确匹配，3匹配失败）;"`                //状态（1匹配中，2正确匹配，3匹配失败）
	Visible       *int              `json:"visible" form:"visible" gorm:"column:visible;comment:是否显示（1: 显示，-1: 隐藏）;"`               //是否显示（1: 显示，-1: 隐藏）
	AccessType    *int              `json:"accessType" form:"accessType" gorm:"column:access_type;comment:接入方式（1域名托管接入）;"`          //接入方式（1域名托管接入）
	ResolveIp     *string           `json:"resolveIp" form:"resolveIp" gorm:"column:resolve_ip;comment:解析IP;size:255;"`             //解析IP
	DomainSource  *int              `json:"domainSource" form:"domainSource" gorm:"column:domain_source;comment:域名来源（如内部、外部、第三方）;"` //域名来源（如内部、外部、第三方）
	Remark        *string           `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:500;"`                         //备注
	CreatedAt     *time.Time        `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;"`                      //创建时间
	UpdatedAt     *time.Time        `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:更新时间;"`                      //更新时间
	DeletedAt     *time.Time        `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:删除时间;"`                      //删除时间
	CreatedBy     *int              `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建人ID;size:19;"`             //创建人ID
	UpdatedBy     *int              `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新人ID;size:19;"`             //更新人ID
	SubDomainList []CSubDominManage `gorm:"-" json:"subDomainList"`
}

// TableName cDominManage表 CDominManage自定义表名 c_domin_manage
func (CDominManage) TableName() string {
	return "c_domin_manage"
}
