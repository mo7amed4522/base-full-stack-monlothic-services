// 自动生成模板MPayChannelConfig
package example

import (
	"time"
)

// mPayChannelConfig表 结构体  MPayChannelConfig
type MPayChannelConfig struct {
	Id           int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                       //id字段
	MUserId      int       `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:m_user主键id;size:19;"`                       //m_user主键id
	CPayConfigId int       `json:"cPayConfigId" form:"cPayConfigId" gorm:"column:c_pay_config_id;comment:c_pay_config主键id;size:19;"` //c_pay_config主键id
	MerAccountId int       `json:"merAccountId" form:"merAccountId" gorm:"column:mer_account_id;comment:关联商户账户ID;size:20;"`          //关联商户账户ID
	PayMinMoney  float64   `json:"payMinMoney" form:"payMinMoney" gorm:"column:pay_min_money;comment:最小支付金额;size:17;"`               //最小支付金额
	PayMaxMoney  float64   `json:"payMaxMoney" form:"payMaxMoney" gorm:"column:pay_max_money;comment:最大支付金额;size:17;"`               //最大支付金额
	Valid        *int      `json:"valid" form:"valid" gorm:"column:valid;comment:是否启用，1代表启用，0代表失效;"`                                 //是否启用，1代表启用，0代表失效
	Remark       string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:100;"`                                   //备注
	CreateTime   time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                             //创建时间
	UpdateTime   time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`                             //更新时间
	DelFlag      *int      `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`                   //逻辑删除状态0（未删除）和 1（已删除）默认0
	CreateBy     *int      `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`                            //创建人
	UpdateBy     *int      `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`                            //修改人

	// 商户名称
	MUserName string `json:"mUserName" form:"mUserName"` // 商户名称
	// 渠道名称
	CPayConfigName   string `json:"cPayConfigName" form:"cPayConfigName"`
	UpdateByNickname string `json:"updateByNickname" form:"updateByNickname" gorm:"-"`
}

// TableName mPayChannelConfig表 MPayChannelConfig自定义表名 m_pay_channel_config
func (MPayChannelConfig) TableName() string {
	return "m_pay_channel_config"
}
