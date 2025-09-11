
// 自动生成模板MBillPayDaily
package system
import (
	"time"
)

// mBillPayDaily表 结构体  MBillPayDaily
type MBillPayDaily struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键ID;size:19;"`  //主键ID
    MUserId  *int `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:m_user主键id;"`  //m_user主键id
    MUserName  *string `json:"mUserName" form:"mUserName" gorm:"column:m_user_name;comment:商户名称;size:32;"`  //商户名称
    BillDate  *time.Time `json:"billDate" form:"billDate" gorm:"column:bill_date;comment:记账日期;"`  //记账日期
    SubmitOrderAmount  *float64 `json:"submitOrderAmount" form:"submitOrderAmount" gorm:"column:submit_order_amount;comment:提交订单金额;"`  //提交订单金额
    SuccessOrderAmount  *float64 `json:"successOrderAmount" form:"successOrderAmount" gorm:"column:success_order_amount;comment:成功订单金额;"`  //成功订单金额
    SubmitOrderNum  *int `json:"submitOrderNum" form:"submitOrderNum" gorm:"column:submit_order_num;comment:订单提交笔数;"`  //订单提交笔数
    SuccessOrderNum  *int `json:"successOrderNum" form:"successOrderNum" gorm:"column:success_order_num;comment:订单成功笔数;"`  //订单成功笔数
    ReviewOrderNum  *int `json:"reviewOrderNum" form:"reviewOrderNum" gorm:"column:review_order_num;comment:补单笔数;"`  //补单笔数
    CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`  //创建时间
    UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`  //更新时间
    DelFlag  *int `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;size:19;"`  //逻辑删除状态0（未删除）和 1（已删除）默认0
    CreateBy  *int `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;"`  //创建人
    UpdateBy  *int `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;"`  //修改人
    Remark  *string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`  //备注
}


// TableName mBillPayDaily表 MBillPayDaily自定义表名 m_bill_pay_daily
func (MBillPayDaily) TableName() string {
    return "m_bill_pay_daily"
}





