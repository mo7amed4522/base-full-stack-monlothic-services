// 自动生成模板OrderPayout
package example

import (
	"time"
)

// orderPayout表 结构体  OrderPayout
type OrderPayout struct {
	Id           int        `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                                //id字段
	MUserId      int        `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:商户id;size:19;"`                                      //商户id
	MUserNo      string     `json:"mUserNo" form:"mUserNo" gorm:"column:m_user_no;comment:商户号;size:30;"`                                       //商户号
	MerSubUser   string     `json:"merSubUser" form:"merSubUser" gorm:"column:mer_sub_user;comment:商户下的用户号;size:100;"`                         //商户下的用户号
	OrderNo      string     `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:交易号;size:100;"`                                       //交易号
	ThirdOrderNo string     `json:"thirdOrderNo" form:"thirdOrderNo" gorm:"column:third_order_no;comment:三方下单订单号;size:100;"`                   //三方下单订单号
	CType        int        `json:"cType" form:"cType" gorm:"column:c_type;comment:渠道类型，1自研，2三方;"`                                             //渠道类型，1自研，2三方
	CPayConfigId int        `json:"cPayConfigId" form:"cPayConfigId" gorm:"column:c_pay_config_id;comment:关联tc_payout_config支付渠道表ID;size:19;"` //关联tc_payout_config支付渠道表ID
	Type         *int       `json:"type" form:"type" gorm:"column:type;comment:0:系统代付，1:商户代付;"`                                                //0:系统代付，1:商户代付
	Currency     string     `json:"currency" form:"currency" gorm:"column:currency;comment:币种:(CNY:人民币 USD美元);size:10;"`                       //币种:(CNY:人民币 USD美元)
	Amount       float64    `json:"amount" form:"amount" gorm:"column:amount;comment:交易金额;size:17;"`                                           //交易金额
	ActualAmount float64    `json:"actualAmount" form:"actualAmount" gorm:"column:actual_amount;comment:实际交易金额;size:17;"`                      //实际交易金额
	OrderBody    string     `json:"orderBody" form:"orderBody" gorm:"column:order_body;comment:订单请求参数体;size:1000;"`                            //订单请求参数体
	OrderDesc    *string    `json:"orderDesc" form:"orderDesc" gorm:"column:order_desc;comment:订单描述;size:100;"`                                //订单描述
	Remark       string     `json:"remark" form:"remark" gorm:"column:remark;comment:订单备注;size:100;"`                                          //订单备注
	CallbackUrl  string     `json:"callbackUrl" form:"callbackUrl" gorm:"column:callback_url;comment:后台回调Url;size:255;"`                       //后台回调Url
	CardName     string     `json:"cardName" form:"cardName" gorm:"column:card_name;comment:收款人姓名;size:100;"`                                  //收款人姓名
	CardNumber   string     `json:"cardNumber" form:"cardNumber" gorm:"column:card_number;comment:收款人卡号;size:100;"`                            //收款人卡号
	CardBank     string     `json:"cardBank" form:"cardBank" gorm:"column:card_bank;comment:银行名称;size:255;"`
	Ifsc         string     `json:"ifsc" form:"ifsc" gorm:"column:ifsc;comment:IFSC;size:100;"`
	PayIp        *string    `json:"payIp" form:"payIp" gorm:"column:pay_ip;comment:交易IP;size:255;"`                           //交易IP
	PayRemark    *string    `json:"payRemark" form:"payRemark" gorm:"column:pay_remark;comment:支付备注信息;size:255;"`             //支付备注信息
	CreateTime   time.Time  `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                     //创建时间
	UpdateTime   *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`                     //修改时间
	UtrNo        *string    `json:"utrNo" form:"utrNo" gorm:"column:utr_no;comment:paytm账号;size:100;"`                        //paytm账号
	DelFlag      int        `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`           //逻辑删除状态0（未删除）和 1（已删除）默认0
	Status       int        `json:"status" form:"status" gorm:"column:status;comment:出款订单状态 1出款申请 2支付中 3出款成功 -1出款失败 4出款待确认;"` //逻辑删除状态0（未删除）和 1（已删除）默认0
	PayTime      *time.Time `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:支付时间;"`                              //支付时间

	CPayPoolAccountName string `json:"cPayPoolAccountName" form:"cPayPoolAccountName" gorm:"-"`                                     // 支持池名称
	CPayConfigName      string `json:"cPayConfigName" form:"cPayConfigName" gorm:"-"`                                               // 渠道名称
	LockStatus          int    `json:"lockStatus" gorm:"column:lock_status"`                                                        // 锁定状态 1锁定 2未锁定
	LockAdminID         int    `json:"lockAdminID" gorm:"column:lock_admin_id"`                                                     // 锁定人
	CPayPoolAccountID   int    `json:"cPayPoolAccountID" gorm:"column:c_pay_pool_account_id"`                                       // 三方出款id
	OperatorID          int    `json:"operatorID" gorm:"column:operator_id"`                                                        // 操作人
	FailMessage         string `json:"failMessage" form:"failMessage" gorm:"column:fail_message;comment:订单备注;size:255;"`            //失败原因
	FailMessageType     int    `json:"failMessageType" form:"failMessageType" gorm:"column:fail_message_type;comment:订单备注;size:2;"` //失败原因type
	UpdateByNickname    string `json:"updateByNickname" form:"updateByNickname" gorm:"-"`
	Message             string `json:"message" form:"message" gorm:"column:message;comment:订单备注;size:255;"` //
	MUserName           string `json:"mUserName" form:"mUserName" gorm:"-"`
	DiffTime            int64  `json:"diffTime" form:"diffTime" gorm:"column:diff_time"`
	LabelCode           string `json:"labelCode" form:"labelCode" gorm:"column:label_code;comment:标签;size:100;"` //标签
	EditTimes           int64  `json:"editTimes" form:"editTimes" gorm:"column:edit_times"`                      // 补单次数
	NotifyStatus        int    `json:"notifyStatus" form:"notifyStatus" gorm:"notify_status"`                    // 回调状态 1回调中，-1回调失败，2回调成功
}

type PayoutOrderStatics struct {
	AvgTime           string `json:"avgTime"`
	TotalAmount       string `json:"totalAmount"`
	SuccessOrderCount string `json:"successOrderCount"`
}

// TableName orderPayout表 OrderPayout自定义表名 order_payout
func (OrderPayout) TableName() string {
	return "order_payout"
}
