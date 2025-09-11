// 自动生成模板MUser
package example

import (
	"time"
)

// TableName mUser表 MUser自定义表名 m_user
func (MUser) TableName() string {
	return "m_user"
}

// mUser表 结构体  MUser
type MUser struct {
	Id         int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键;size:19;"`                      //主键
	Name       *string   `json:"name" form:"name" gorm:"column:name;comment:用户名;size:100;"`                         //用户名
	LoginPwd   *string   `json:"loginPwd" form:"loginPwd" gorm:"column:login_pwd;comment:登录密码;size:200;"`           //登录密码
	PayPwd     *string   `json:"payPwd" form:"payPwd" gorm:"column:pay_pwd;comment:支付密码;size:100;"`                 //支付密码
	No         string    `json:"no" form:"no" gorm:"column:no;comment:商户号;size:50;"`                                //商户号
	Phone      *string   `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:191;"`                      //手机号
	Email      *string   `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:200;"`                       //邮箱
	PublicKey  *string   `json:"publicKey" form:"publicKey" gorm:"column:publicKey;comment:公钥;size:1000;"`          //公钥
	PrivateKey *string   `json:"privateKey" form:"privateKey" gorm:"column:privateKey;comment:私钥;size:1000;"`       //私钥
	IsLimitIp  *int      `json:"isLimitIp" form:"isLimitIp" gorm:"column:is_limit_ip;comment:是否限制ip，0不限制，1限制，默认0;"` //是否限制ip，0不限制，1限制，默认0
	Valid      *int      `json:"valid" form:"valid" gorm:"column:valid;comment:是否支付代付生效 0 无效 1有效 默认0;"`             //是否支付代付生效 0 无效 1有效 默认0
	Status     *int      `json:"status" form:"status" gorm:"column:status;comment:用户使用状态，0 停用 1 启用 默认0;size:19;"`   //用户使用状态，0 停用 1 启用 默认0
	Remark     *string   `json:"remark" form:"remark" gorm:"column:remark;comment:备注信息;size:200;"`                  //备注信息
	Language   *string   `json:"language" form:"language" gorm:"column:language;comment:语言设置;size:50;"`             //语言设置
	CreateBy   *string   `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:255;"`            //创建人
	UpdateBy   *int      `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;"`                     //修改人
	CreatedAt  time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`                     //createdAt字段
	UpdatedAt  time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`                     //updatedAt字段
	DeletedAt  time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`                     //deletedAt字段
	DelFlag    *int      `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`    //逻辑删除状态0（未删除）和 1（已删除）默认0

	// 赠送配置
	Config           string `json:"config" form:"config" gorm:"column:config;type:text;comment:赠送配置;"` // 赠送配置
	UpdateByNickname string `json:"updateByNickname" form:"updateByNickname" gorm:"-"`
}
