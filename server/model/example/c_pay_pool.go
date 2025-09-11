// 自动生成模板CPayPool
package example

import (
	"time"
)

// cPayPool表 结构体  CPayPool
type CPayPool struct {
	Id                   int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                               //id字段
	MUserId              *int      `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:商户id;"`                                             //商户id
	CPayConfig           int       `json:"cPayConfig" form:"cPayConfig" gorm:"column:c_pay_config;comment:关联c_pay_config表ID，渠道id;size:19;"`          //关联c_pay_config表ID，渠道id
	Code                 string    `json:"code" form:"code" gorm:"column:code;comment:池编号;size:50;"`                                                 //池编号
	Name                 string    `json:"name" form:"name" gorm:"column:name;comment:池名称;size:100;"`                                                //池名称
	TotalAmount          float64   `json:"totalAmount" form:"totalAmount" gorm:"column:total_amount;comment:已经支付的总金额;"`                              //已经支付的总金额
	CPayConfigId         int       `json:"cPayConfigId" form:"cPayConfigId" gorm:"column:c_pay_config_id;comment:关联渠道id;"`                           //关联渠道id
	LimitCountRequestDay int       `json:"limitCountRequestDay" form:"limitCountRequestDay" gorm:"column:limit_count_request_day;comment:每天请求次数限制;"` //每天请求次数限制
	LimitCountUsedDay    int       `json:"limitCountUsedDay" form:"limitCountUsedDay" gorm:"column:limit_count_used_day;comment:每天实际完成次数限制;"`        //每天实际完成次数限制
	LimitAmountDay       float64   `json:"limitAmountDay" form:"limitAmountDay" gorm:"column:limit_amount_day;comment:限制每天交易金额;"`                    //限制每天交易金额
	LimitAmount          float64   `json:"limitAmount" form:"limitAmount" gorm:"column:limit_amount;comment:限制交易总金额;"`                               //限制交易总金额
	LimitAmountMin       float64   `json:"limitAmountMin" form:"limitAmountMin" gorm:"column:limit_amount_min;comment:每次支付的最小金额;"`                   //每次支付的最小金额
	LimitAmountMax       float64   `json:"limitAmountMax" form:"limitAmountMax" gorm:"column:limit_amount_max;comment:每次支付的最大金额;"`                   //每次支付的最大金额
	Valid                *int      `json:"valid" form:"valid" gorm:"column:valid;comment:是否启用，1代表启用，0代表停用;"`                                         //是否启用，1代表启用，0代表停用
	Remark               string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                                           //备注
	CreateBy             int       `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;"`                                            //创建人
	UpdateBy             int       `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;"`                                            //修改人
	CreatedAt            time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`                                            //createdAt字段
	UpdatedAt            time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`                                            //updatedAt字段
	DeletedAt            time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`                                            //deletedAt字段
	UpdateByNickname     string    `json:"updateByNickname" form:"updateByNickname" gorm:"-"`
}

// 列表响应结构体 在model基础上增加了一些字段
type CPayPoolResp struct {
	CPayPool
	AllAccountNum     int64 `json:"allAccountNum"`
	ValidAccountNum   int64 `json:"validAccountNum"`
	InvalidAccountNum int64 `json:"invalidAccountNum"`
}

// TableName cPayPool表 CPayPool自定义表名 c_pay_pool
func (CPayPool) TableName() string {
	return "c_pay_pool"
}
