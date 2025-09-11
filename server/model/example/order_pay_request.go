// 自动生成模板OrderPayRequest
package example

import (
	"time"
)

// orderPayRequest表 结构体  OrderPayRequest
type OrderPayRequest struct {
	Id      int    `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`           //id字段
	MUserId string `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:商户id;size:30;"` //商户id

	CPayConfigId  string    `json:"cPayConfigId" form:"cPayConfigId" gorm:"column:c_pay_config_id;comment:c_pay_config主键id;size:10;"` //c_pay_config主键id
	RequestData   string    `json:"requestData" form:"requestData" gorm:"column:request_data;comment:请求参数;"`                          //请求参数
	Status        int       `json:"status" form:"status" gorm:"column:status;comment:状态：0为异常，1为成功;"`                                  //状态：0为异常，1为成功
	ReponseData   string    `json:"reponseData" form:"reponseData" gorm:"column:reponse_data;comment:响应信息;"`                          //响应信息
	ReponseTime   int       `json:"reponseTime" form:"reponseTime" gorm:"column:reponse_time;comment:响应时间（ms）;size:19;"`              //响应时间（ms）
	CreateTime    time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                             //创建时间
	ExceptionType int       `json:"exceptionType" form:"exceptionType" gorm:"column:exception_type;comment:异常类型 0为无异常;"`              //异常类型 0为无异常
	DelFlag       int       `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`                   //逻辑删除状态0（未删除）和 1（已删除）默认0
	UpdateTime    time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`                             //修改时间
	CreateBy      int       `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`                            //创建人
	UpdateBy      int       `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`                            //修改人
	Remark        string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
	ThirdOrderNo  string    `json:"thirdOrderNo" form:"thirdOrderNo" gorm:"column:third_order_no;comment:三方订单号;size:255;"` //三方订单号
}

// TableName orderPayRequest表 OrderPayRequest自定义表名 order_pay_request
func (OrderPayRequest) TableName() string {
	return "order_pay_request"
}
