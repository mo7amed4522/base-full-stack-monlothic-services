// 自动生成模板OrderPayPlayer
package example

import (
	"time"
)

// orderPayPlayer表 结构体  OrderPayPlayer
type OrderPayPlayer struct {
	Id      *int    `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"` //id字段
	MUserId *string `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:商户id
;size:100;"` //商户id

	CreateTime     *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`                                                                //createTime字段
	UpdateTime     *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`                                                                //updateTime字段
	PlayerId       *string    `json:"playerId" form:"playerId" gorm:"column:player_id;comment:商户下用户id;size:100;"`                                                      //商户下用户id
	Valid          *int       `json:"valid" form:"valid" gorm:"column:valid;comment:是否启用，1代表启用，0代表停用;size:10;"`                                                        //是否启用，1代表启用，0代表停用
	HistoryAccount *string    `json:"historyAccount" form:"historyAccount" gorm:"column:history_account;comment:使用过的历史账号信息，对应c_pay_pool_account表主键ID,最多10条;size:500;"` //使用过的历史账号信息，对应c_pay_pool_account表主键ID,最多10条
	PlayerMsg      *string    `json:"playerMsg" form:"playerMsg" gorm:"column:player_msg;comment:用户标识名称/描述;size:500;"`                                                 //用户标识名称/描述
	DelFlag        *int       `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`                                                  //逻辑删除状态0（未删除）和 1（已删除）默认0
	CreateBy       *int       `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:19;"`                                                           //创建人
	UpdateBy       *int       `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:19;"`                                                           //修改人
	Remark         *string    `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                                                                  //备注
}

// TableName orderPayPlayer表 OrderPayPlayer自定义表名 order_pay_player
func (OrderPayPlayer) TableName() string {
	return "order_pay_player"
}
