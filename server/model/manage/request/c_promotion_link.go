
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type CPromotionLinkSearch struct{
    Name  *string `json:"name" form:"name" `
    AppId  *int `json:"appId" form:"appId" `
    TemplateId  *int `json:"templateId" form:"templateId" `
    IsAllowEditUrl  *int `json:"isAllowEditUrl" form:"isAllowEditUrl" `
    DomainId  *int `json:"domainId" form:"domainId" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
