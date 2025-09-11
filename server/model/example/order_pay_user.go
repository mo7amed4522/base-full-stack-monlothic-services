// 自动生成模板OrderPayUser
package example

import (
	"time"
)

// orderPayUser表 结构体  OrderPayUser
type OrderPayUser struct {
	Id         int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                                    //id字段
	MUserNo    string    `json:"mUserNo" form:"mUserNo" gorm:"column:m_user_no;comment:商户号;size:100;"`                                          //商户号
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`                                              //createTime字段
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`                                              //updateTime字段
	SubUserId  string    `json:"subUserId" form:"subUserId" gorm:"column:sub_user_id;comment:商户下用户id;size:100;"`                                //商户下用户id
	Valid      int       `json:"valid" form:"valid" gorm:"column:valid;comment:是否启用，1代表启用，0代表停用;size:10;"`                                      //是否启用，1代表启用，0代表停用
	HisAccount string    `json:"hisAccount" form:"hisAccount" gorm:"column:his_account;comment:使用过的历史账号信息，对应sm_tc_qrcode表主键ID,最多10条;size:500;"` //使用过的历史账号信息，对应sm_tc_qrcode表主键ID,最多10条
	SubUserMsg string    `json:"subUserMsg" form:"subUserMsg" gorm:"column:sub_user_msg;comment:用户标识名称/描述;size:500;"`                           //用户标识名称/描述
	DelFlag    int       `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`                                //逻辑删除状态0（未删除）和 1（已删除）默认0
}

// TableName orderPayUser表 OrderPayUser自定义表名 order_pay_user
func (OrderPayUser) TableName() string {
	return "order_pay_user"
}
