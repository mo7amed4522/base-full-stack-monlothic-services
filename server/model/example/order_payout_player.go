// 自动生成模板OrderPayoutPlayer
package example

import (
	"time"
)

// orderPayoutPlayer表 结构体  OrderPayoutPlayer
type OrderPayoutPlayer struct {
	Id      *int    `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`            //id字段
	MUserId *string `json:"mUserId" form:"mUserId" gorm:"column:m_user_id;comment:商户id;size:100;"` //商户id

	CreateTime     *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`                                                          //createTime字段
	UpdateTime     *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`                                                          //updateTime字段
	PlayerId       *string    `json:"playerId" form:"playerId" gorm:"column:player_id;comment:商户下用户id;size:100;"`                                                //商户下用户id
	Valid          *int       `json:"valid" form:"valid" gorm:"column:valid;comment:是否启用，1代表启用，0代表停用;size:10;"`                                                  //是否启用，1代表启用，0代表停用
	HistoryAccount *string    `json:"historyAccount" form:"historyAccount" gorm:"column:history_account;comment:使用过的历史账号信息，对应sm_tc_qrcode表主键ID,最多10条;size:500;"` //使用过的历史账号信息，对应sm_tc_qrcode表主键ID,最多10条
	PlayerMsg      *string    `json:"playerMsg" form:"playerMsg" gorm:"column:player_msg;comment:用户标识名称/描述;size:500;"`                                           //用户标识名称/描述
	DelFlag        *int       `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:逻辑删除状态0（未删除）和 1（已删除）默认0;"`                                            //逻辑删除状态0（未删除）和 1（已删除）默认0
}

// TableName orderPayoutPlayer表 OrderPayoutPlayer自定义表名 order_payout_player
func (OrderPayoutPlayer) TableName() string {
	return "order_payout_player"
}
