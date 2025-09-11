package request

import "time"

// 订单请求日志体
type PayOrderLogsReq struct {
	OrderNo   string    `json:"orderNo" form:"orderNo"`
	Amount    float64   `json:"amount" form:"amount"`
	MUserId   int       `json:"mUserId" form:"mUserId"`
	ReqBody   string    `json:"reqBody" form:"reqBody"`
	ResBody   string    `json:"resBody" form:"resBody"`
	StartTime time.Time `json:"startTime" form:"startTime"`
	EndTime   time.Time `json:"EndTime" form:"EndTime"`
	Remark    string    `json:"remark" form:"remark"`
	Status    int       `json:"status" form:"status"` //1失败 2成功
}
