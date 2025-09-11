// 自动生成模板CSubDominManage
package manage

import (
	"time"
)

// cSubDominManage表 结构体  CSubDominManage
type CSubDominManage struct {
	Id        *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键ID;size:20;"`                                //主键ID
	CDominId  int        `json:"cDominId" form:"cDominId" gorm:"column:c_domin_id;comment:主域名表ID;size:19;" binding:"required"`  //主域名表ID
	SubDomain string     `json:"subDomain" form:"subDomain" gorm:"column:sub_domain;comment:域名地址;size:255;" binding:"required"` //域名地址
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;"`                             //创建时间
}

// TableName cSubDominManage表 CSubDominManage自定义表名 c_sub_domin_manage
func (CSubDominManage) TableName() string {
	return "c_sub_domin_manage"
}
