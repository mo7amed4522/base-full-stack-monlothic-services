package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CDominManageSearch struct {
	request.PageInfo
	Sort   string `json:"sort" form:"sort"`
	Order  string `json:"order" form:"order"`
	ID     int    `json:"id" form:"id"`
	Domain string `json:"domain" form:"domain"`
}
