package example

import "time"

// Transactions 账变记录
type Transactions struct {
	ID                  int       `json:"id"`
	Type                int       `json:"type"` // 账变类型 1收款 2出款
	CPayPoolAccountID   int       `json:"cPayPoolAccountId" gorm:"column:c_pay_pool_account_id"`
	BeforeAmount        float64   `json:"beforeAmount" gorm:"column:before_amount"`
	Amount              float64   `json:"amount"`
	AfterAmount         float64   `json:"afterAmount" gorm:"column:after_amount"`
	Desc                string    `json:"desc"`
	CreateTime          time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"` //创建时间
	UpdateTime          time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"` //更新时间
	CPayPoolAccountName string    `json:"cPayPoolAccountName" gorm:"-"`
	OperatorID          int       `json:"operatorID" gorm:"column:operator_id"` // 操作人
	OperatorName        string    `json:"operatorName" gorm:"-"`
}
