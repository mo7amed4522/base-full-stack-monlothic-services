
// 自动生成模板CPageTemplate
package manage
import (
	"time"
)

// cPageTemplate表 结构体  CPageTemplate
type CPageTemplate struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`  //id字段
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:模版名称;size:64;" binding:"required"`  //模版名称
    Code  *string `json:"code" form:"code" gorm:"column:code;comment:模版code;size:32;" binding:"required"`  //模版code
    CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`  //createdAt字段
    UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`  //updatedAt字段
}


// TableName cPageTemplate表 CPageTemplate自定义表名 c_page_template
func (CPageTemplate) TableName() string {
    return "c_page_template"
}





