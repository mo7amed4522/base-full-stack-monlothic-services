// 自动生成模板MUserWallet
package example

import (
	"time"
)

// mUserWallet表 结构体  MUserWallet
type MUserWallet struct {
	Id           *int      `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                       //id字段
	MUserId      int       `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:m_user主键id;size:19;"`                       //m_user主键id
	CPayConfigId *int      `json:"cPayConfigId" form:"cPayConfigId" gorm:"column:c_pay_config_id;comment:c_pay_config主键id;size:19;"` //c_pay_config主键id
	Status       *int      `json:"status" form:"status" gorm:"column:status;comment:账户状态 0正常 1冻结;"`                                  //账户状态 0正常 1冻结
	CreateTime   time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`                                 //createTime字段
	UpdateTime   time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`                                 //updateTime字段
	DelFlag      *int      `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）;"`                      //逻辑删除状态0（未删除）和 1（已删除）
	Balance      *float64  `json:"balance" form:"balance" gorm:"column:balance;comment:商户渠道账户余额;size:10;"`                           //商户渠道账户余额
	CreateBy     *int      `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`                            //创建人
	UpdateBy     *int      `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`                            //修改人
}

// TableName mUserWallet表 MUserWallet自定义表名 m_user_wallet
func (MUserWallet) TableName() string {
	return "m_user_wallet"
}
