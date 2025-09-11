package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CPayPoolSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Code           *string    `json:"code" form:"code" `
	Name           *string    `json:"name" form:"name" `
	MUserID        *int64     `json:"mUserId" form:"mUserId" `
	CPayPoolID     *int64     `json:"cPayPoolId" form:"cPayPoolId" `
	Qrcode         *string    `json:"qrcode" form:"qrcode" `

	request.PageInfo
}
