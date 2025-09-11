// 自动生成模板CPayConfig
package example

import (
	"time"
)

// cPayConfig表 结构体  CPayConfig
type CPayConfig struct {
	Id              *int      `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                       //id字段
	ChannelCode     *string   `json:"channelCode" form:"channelCode" gorm:"column:channel_code;comment:第三方编码;size:200;"`                //第三方编码
	Name            *string   `json:"name" form:"name" gorm:"column:name;comment:配置名称;size:200;"`                                       //配置名称
	ChannelValue    *string   `json:"channelValue" form:"channelValue" gorm:"column:channel_value;comment:配置值;size:100;"`               //配置值
	Valid           *int      `json:"valid" form:"valid" gorm:"column:valid;comment:是否启用，1代表启用，0代表停用;"`                                 //是否启用，1代表启用，0代表停用
	ChannelType     *int      `json:"channelType" form:"channelType" gorm:"column:channel_type;comment:通道类型 1自研 2三方;"`                  //通道类型 1自研 2三方
	ThirdMerchantNo *string   `json:"thirdMerchantNo" form:"thirdMerchantNo" gorm:"column:third_merchant_no;comment:三方渠道商户号;size:100;"` //三方渠道商户号
	ThirdPublicKey  *string   `json:"thirdPublicKey" form:"thirdPublicKey" gorm:"column:third_public_key;comment:三方渠道公钥;size:500;"`     //三方渠道公钥
	ThirdPrivateKey *string   `json:"thirdPrivateKey" form:"thirdPrivateKey" gorm:"column:third_private_key;comment:三方渠道私钥;size:500;"`  //三方渠道私钥
	ThirdUrl        *string   `json:"thirdUrl" form:"thirdUrl" gorm:"column:third_url;comment:三方渠道请求地址;size:500;"`                      //三方渠道请求地址
	Remark          *string   `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                                   //备注
	CreatedAt       time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`                                    //createdAt字段
	UpdatedAt       time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`                                    //updatedAt字段
	DeletedAt       time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`                                    //deletedAt字段
	CreatedBy       *int      `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`                         //创建者
	UpdatedBy       *int      `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`                         //更新者
	DeletedBy       *int      `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`                         //删除者
	// 增加pay_type 字段, string 类型。 1 二维码； 2和3银行卡
	PayType string `json:"payType" form:"payType" gorm:"column:pay_type;comment:支付类型;size:255;"` // 支付类型
	Type    *int   `json:"type" form:"type" gorm:"column:type;comment:支付类型 1支付 2出款 ;"`           // 支付类型
}

// TableName cPayConfig表 CPayConfig自定义表名 c_pay_config
func (CPayConfig) TableName() string {
	return "c_pay_config"
}
