// 自动生成模板PayBank
package example

import (
	"time"
)

// payBank表 结构体  PayBank
type PayBank struct {
	Id         *int      `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`              //id字段
	BankName   *string   `json:"bankName" form:"bankName" gorm:"column:bank_name;comment:银行名;size:100;"`  //银行名
	BankCode   *string   `json:"bankCode" form:"bankCode" gorm:"column:bank_code;comment:银行编码;size:100;"` //银行编码
	DelFlag    *int      `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:是否删除   0未删除  1删除;"` //是否删除   0未删除  1删除
	CreateBy   *int      `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`   //创建人
	UpdateBy   *int      `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`   //修改人
	Remark     *string   `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`          //备注
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:添加时间;"`    //添加时间
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`    //修改时间
}

// TableName payBank表 PayBank自定义表名 pay_bank
func (PayBank) TableName() string {
	return "pay_bank"
}
