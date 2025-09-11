
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type CPageTemplateSearch struct{
    Id  *int `json:"id" form:"id" `
    Name  *string `json:"name" form:"name" `
    Code  *string `json:"code" form:"code" `
    request.PageInfo
}
