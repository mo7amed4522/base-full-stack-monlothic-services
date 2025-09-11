// 自动生成模板CPromotionLink
package manage

import (
	"time"
)

// cPromotionLink表 结构体  CPromotionLink
type CPromotionLink struct {
	Id             *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                   //id字段
	Name           *string    `json:"name" form:"name" gorm:"column:name;comment:链接名称;size:64;"`                    //链接名称
	AppId          int        `json:"appId" form:"appId" gorm:"column:app_id;comment:app id;size:10;"`              //app id
	AppName        string     `json:"appName" form:"appName" gorm:"-"`                                              //app名称
	TemplateId     int        `json:"templateId" form:"templateId" gorm:"column:template_id;comment:模版id;size:10;"` //模版id
	TemplateName   string     `json:"templateName" form:"templateName" gorm:"-"`                                    //模版名称
	AccessUrl      string     `json:"accessUrl" form:"accessUrl" gorm:"column:access_url"`
	Url            *string    `json:"url" form:"url" gorm:"column:url;comment:包网下载地址;size:255;"`                                           //包网下载地址
	IsAllowEditUrl *int       `json:"isAllowEditUrl" form:"isAllowEditUrl" gorm:"column:is_allow_edit_url;comment:是否允许修改url 1 允许 -1 不允许;"` //是否允许修改url 1 允许 -1 不允许
	IosRedirectUrl *string    `json:"iosRedirectUrl" form:"iosRedirectUrl" gorm:"column:ios_redirect_url;comment:ios跳转地址;size:255;"`       //ios跳转地址
	DomainId       int        `json:"domainId" form:"domainId" gorm:"column:domain_id;comment:投放域名id;size:10;"`                            //投放域名id
	DomainURL      string     `json:"domainURL" form:"domainURL" gorm:"-"`                                                                 //投放域名url
	Remark         *string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                                      //备注
	CreatedAt      *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`                                       //createdAt字段
	UpdatedAt      *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`                                       //updatedAt字段
}

// TableName cPromotionLink表 CPromotionLink自定义表名 c_promotion_link
func (CPromotionLink) TableName() string {
	return "c_promotion_link"
}
