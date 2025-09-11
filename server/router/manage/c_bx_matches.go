package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CBxMatchesRouter struct{}

// InitBxMatchesRouter 初始化 InitBxMatchesRouter 路由信息
func (s *CAppManageRouter) InitBxMatchesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cAppManage := Router.Group("bxMatches").Use(middleware.OperationRecord())
	{
		cAppManage.GET("getMatchesList", cBxMatchesApi.GetMatchesList)
		cAppManage.POST("batchSetHot", cBxMatchesApi.BatchSetHot)
		cAppManage.POST("batchSetTop", cBxMatchesApi.BatchSetTop)
	}
}
