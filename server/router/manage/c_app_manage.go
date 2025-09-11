package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CAppManageRouter struct{}

// InitAppManageRouter 初始化 InitAppManageRouter 路由信息
func (s *CAppManageRouter) InitAppManageRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cAppManage := Router.Group("appManage").Use(middleware.OperationRecord())
	{
		cAppManage.POST("createApp", CAppManageApi.CreateCApp)
		cAppManage.POST("getAppPage", CAppManageApi.GetCAppList)
		cAppManage.DELETE("deleteAppById", CAppManageApi.DeleteApp)
		cAppManage.PUT("updateApp", CAppManageApi.UpdateCPayConfig)
	}
	pAppManage := PublicRouter.Group("appManage")
	{
		pAppManage.GET("appInfo", CAppManageApi.GetCPayConfigPublic)
	}
}
