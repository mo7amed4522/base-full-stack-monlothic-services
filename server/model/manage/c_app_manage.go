// 自动生成模板CPayPool
package manage

import (
	"encoding/json"
	"time"
)

// CAppManage 应用管理模型（简化单表设计）
type CAppManage struct {
	ID           uint            `gorm:"primaryKey;column:id" json:"id"`
	AppName      string          `gorm:"column:app_name;not null" json:"appName"`            // 应用名称
	AppRemark    *string         `gorm:"column:app_remark" json:"appRemark"`                 // 应用备注
	AppIcon      *string         `gorm:"column:app_icon" json:"appIcon"`                     // 应用图标
	LanguageJson json.RawMessage `gorm:"column:language_json;type:json" json:"languageJson"` // 多语言配置JSON
	Status       int             `gorm:"column:status;default:1" json:"status"`              // 状态 1:启用 0:禁用
	CreatedAt    time.Time       `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time       `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt    *time.Time      `gorm:"column:deleted_at" jsonp:"deletedAt"`
	CreatedBy    *uint           `gorm:"column:created_by" json:"createdBy"`     // 创建人ID
	UpdatedBy    *uint           `gorm:"column:updated_by" json:"updatedBy"`     // 更新人ID
	ProjectPath  string          `gorm:"column:project_path" json:"projectPath"` // 包路径
	AccessUrl    string          `gorm:"column:access_url;" json:"accessUrl"`    // 接入url
	DomainId     uint            `gorm:"column:domain_id" json:"domainId"`       //域名Id

	PromotionLinkInfo CPromotionLink `gorm:"foreignKey:ID;references:AppId" json:"promotionLinkInfo"`
}

// TableName 表名
func (CAppManage) TableName() string {
	return "c_app_manage"
}
