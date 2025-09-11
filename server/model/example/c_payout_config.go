// 自动生成模板CPayoutConfig
package example

import (
	"time"
)

// cPayoutConfig表 结构体  CPayoutConfig
type CPayoutConfig struct {
	Id          *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                      //id字段
	ThreeCode   *string    `json:"threeCode" form:"threeCode" gorm:"column:three_code;comment:第三方编码;size:200;"`     //第三方编码
	Name        *string    `json:"name" form:"name" gorm:"column:name;comment:配置名称;size:200;"`                      //配置名称
	Value       *string    `json:"value" form:"value" gorm:"column:value;comment:配置值;size:100;"`                    //配置值
	Valid       *int       `json:"valid" form:"valid" gorm:"column:valid;comment:是否启用，1代表启用，0代表停用;"`                //是否启用，1代表启用，0代表停用
	CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`                //createTime字段
	UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`                //updateTime字段
	PayType     *int       `json:"payType" form:"payType" gorm:"column:pay_type;comment:交易类型 1支付 2代付;"`             //交易类型 1支付 2代付
	ChannelType *int       `json:"channelType" form:"channelType" gorm:"column:channel_type;comment:通道类型 1自研 2三方;"` //通道类型 1自研 2三方
	CreateBy    *int       `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`           //创建人
	UpdateBy    *int       `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`           //修改人
	Remark      *string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                  //备注
	DelFlag     *int       `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`  //逻辑删除状态0（未删除）和 1（已删除）默认0
}

// TableName cPayoutConfig表 CPayoutConfig自定义表名 c_payout_config
func (CPayoutConfig) TableName() string {
	return "c_payout_config"
}
