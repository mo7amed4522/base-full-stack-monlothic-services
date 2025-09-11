package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CSubDominManageSearch struct {
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

type CSubDominManageAdd struct {
	CDominId   *int     `json:"cDominId" form:"cDominId" gorm:"column:c_domin_id;comment:主域名表ID;size:19;" binding:"required"` //主域名表ID
	SubDomains []string `json:"subDomains" form:"subDomains" binding:"required"`                                              //域名地址
}
