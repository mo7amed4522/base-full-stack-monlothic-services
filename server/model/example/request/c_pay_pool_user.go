package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CPayPoolUserSearch struct {
	request.PageInfo
	SearchMap []SearchMap `json:"searchMap"`
}
