// 自动生成模板OrderPay
package example

import (
	"time"
)

// orderPay表 结构体  OrderPay
type OrderPay struct {
	Id            int     `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                           //id字段
	MUserId       int     `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:m_user主键id;size:19;"`                           //m_user主键id
	MUserNo       string  `json:"mUserNo" form:"mUserNo" gorm:"column:m_user_no;comment:商户号;size:30;"`                                  //商户号
	MUserName     *string `json:"mUserName" form:"mUserName" gorm:"column:m_user_name;comment:商户名称;size:100;"`                          //商户名称
	MUserWalletId int     `json:"mUserWalletId" form:"mUserWalletId" gorm:"column:m_user_wallet_id;comment:m_user_wallet主键id;size:19;"` //m_user_wallet主键id
	MerSubUser    string  `json:"merSubUser" form:"merSubUser" gorm:"column:mer_sub_user;comment:商户下的用户号;size:100;"`                    //商户下的用户号
	OrderNo       string  `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:交易号;size:100;"`                                  //交易号
	// 唯一键
	ThirdOrderNo      string     `json:"thirdOrderNo" form:"thirdOrderNo" gorm:"column:third_order_no;uniqueIndex:third_order_no;comment:三方下单订单号;size:100;"`     //三方下单订单号
	ThirdInnerOrderNo *string    `json:"thirdInnerOrderNo" form:"thirdInnerOrderNo" gorm:"column:third_inner_order_no;comment:第三方内部订单号;size:100;"`               //第三方内部订单号
	CType             *int       `json:"cType" form:"cType" gorm:"column:c_type;comment:渠道类型，1自研，2三方;"`                                                          //渠道类型，1自研，2三方
	CPayConfigId      *int       `json:"cPayConfigId" form:"cPayConfigId" gorm:"column:c_pay_config_id;comment:c_pay_config主键id;size:19;"`                       //c_pay_config主键id
	CPayPoolId        int        `json:"cPayPoolId" form:"cPayPoolId" gorm:"column:c_pay_pool_id;comment:c_pay_pool主键id;size:19;"`                               //c_pay_pool主键id
	CPayPoolAccountId *int       `json:"cPayPoolAccountId" form:"cPayPoolAccountId" gorm:"column:c_pay_pool_account_id;comment:c_pay_pool_account主键id;size:19;"` //c_pay_pool_account主键id
	Currency          *string    `json:"currency" form:"currency" gorm:"column:currency;comment:币种:(CNY:人民币 USD美元);size:10;"`                                    //币种:(CNY:人民币 USD美元)
	Amount            float64    `json:"amount" form:"amount" gorm:"column:amount;comment:交易金额;size:22;"`                                                        //交易金额
	ActualAmount      *float64   `json:"actualAmount" form:"actualAmount" gorm:"column:actual_amount;comment:实际交易金额;size:22;"`                                   //实际交易金额
	OrderBody         string     `json:"orderBody" form:"orderBody" gorm:"column:order_body;comment:商品信息;size:1000;"`                                            //商品信息
	OrderDesc         *string    `json:"orderDesc" form:"orderDesc" gorm:"column:order_desc;comment:订单描述;size:100;"`                                             //订单描述
	Remark            string     `json:"remark" form:"remark" gorm:"column:remark;comment:订单备注;size:100;"`                                                       //订单备注
	CallbackUrl       string     `json:"callbackUrl" form:"callbackUrl" gorm:"column:callback_url;comment:后台回调Url;size:255;"`                                    //后台回调Url
	ReturnUrl         string     `json:"returnUrl" form:"returnUrl" gorm:"column:return_url;comment:前端返回地址;size:255;"`                                           //前端返回地址
	InnerStatus       *int       `json:"innerStatus" form:"innerStatus" gorm:"column:inner_status;comment:内部对账状态 0未开始 1开始对账 2 对账成功 -1对账失败;"`                     //内部对账状态 0未开始 1开始对账 2 对账成功 -1对账失败
	InnnerCheckTime   *time.Time `json:"innnerCheckTime" form:"innnerCheckTime" gorm:"column:innner_check_time;comment:内部对账时间;"`                                 //内部对账时间
	OuterStatus       *int       `json:"outerStatus" form:"outerStatus" gorm:"column:outer_status;comment:外部部对账状态 0未开始 1开始对账 2 对账成功 -1对账失败;"`                    //外部部对账状态 0未开始 1开始对账 2 对账成功 -1对账失败
	OuterCheckTime    *time.Time `json:"outerCheckTime" form:"outerCheckTime" gorm:"column:outer_check_time;comment:外部对账时间;"`                                    //外部对账时间
	PayIp             *string    `json:"payIp" form:"payIp" gorm:"column:pay_ip;comment:交易IP;size:255;"`
	// 交易类型，qrcode就是二维码，，bank和imps就是银行卡，这个对应收款方式那两个类型
	// 收款方式，有2种，二维码和银行卡
	// 对应关系就是payType 的 就是 1和3是二维码， bank就是银行卡
	PayType     string     `json:"payType" form:"payType" gorm:"column:pay_type;comment:交易类型;size:255;"`             //                                                  //交易IP
	PayRemark   *string    `json:"payRemark" form:"payRemark" gorm:"column:pay_remark;comment:支付备注信息;size:255;"`     //支付备注信息
	CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`             //创建时间
	UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`             //修改时间
	PlatUserId  *int       `json:"platUserId" form:"platUserId" gorm:"column:plat_user_id;comment:运营人员的id;size:19;"` //运营人员的id
	OrderSource *int       `json:"orderSource" form:"orderSource" gorm:"column:order_source;comment:支付来源;"`          //支付来源
	UtrNo       string     `json:"utrNo" form:"utrNo" gorm:"column:utr_no;comment:paytm账号;size:100;"`                //paytm账号
	DelFlag     *int       `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`   //逻辑删除状态0（未删除）和 1（已删除）默认0
	CreateBy    *int       `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`            //创建人
	UpdateBy    *int       `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`            //修改人
	CreatedBy   *int       `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;"`                 //创建者
	UpdatedBy   *int       `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;"`                 //更新者
	DeletedBy   *int       `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;"`                 //删除者
	Status      int        `json:"status" form:"status" gorm:"column:status;comment:支付状态;"`                          //删除者
	CreatedAt   time.Time  `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`                    //createdAt字段
	UpdatedAt   time.Time  `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`                    //updatedAt字段
	DeletedAt   time.Time  `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`                    //deletedAt字段
	PayTime     time.Time  `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:支付时间;"`                      //支付时间

	// 支持池名称
	CPayPoolName string `json:"cPayPoolName" form:"cPayPoolName"` // 支持池名称
	// 渠道名称
	CPayConfigName string `json:"cPayConfigName" form:"cPayConfigName"` // 渠道名称
	// 收款方式名称
	CPayPoolAccountName string `json:"cPayPoolAccountName" form:"cPayPoolAccountName"`                                  // 收款方式名称
	LabelCode           string `json:"labelCode" form:"labelCode" gorm:"column:label_code;comment:labelCode;size:100;"` // 标签code
	CardNumber          string `json:"cardNumber" form:"cardNumber" gorm:"-"`
	// 还需要个 UTR NO 字段
	UTRTime      int64 `json:"UTRTime" form:"UTRTime" gorm:"column:utr_time"`
	NotifyStatus int   `json:"notifyStatus" form:"notifyStatus" gorm:"notify_status"` // 回调状态 1回调中，-1回调失败，2回调成功
}

type PayorderStatics struct {
	AvgTime           string `json:"avgTime"`
	TotalAmount       string `json:"totalAmount"`
	SuccessOrderCount string `json:"successOrderCount"`
}

// TableName orderPay表 OrderPay自定义表名 order_pay
func (OrderPay) TableName() string {
	return "order_pay"
}
