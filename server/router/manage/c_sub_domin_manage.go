package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CSubDominManageRouter struct{}

// InitCSubDominManageRouter 初始化 cSubDominManage表 路由信息
func (s *CSubDominManageRouter) InitCSubDominManageRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cSubDominManageRouter := Router.Group("cSubDominManage").Use(middleware.OperationRecord())
	cSubDominManageRouterWithoutRecord := Router.Group("cSubDominManage")
	cSubDominManageRouterWithoutAuth := PublicRouter.Group("cSubDominManage")
	{
		cSubDominManageRouter.POST("createCSubDominManage", cSubDominManageApi.CreateCSubDominManage)             // 新建cSubDominManage表
		cSubDominManageRouter.DELETE("deleteCSubDominManage", cSubDominManageApi.DeleteCSubDominManage)           // 删除cSubDominManage表
		cSubDominManageRouter.DELETE("deleteCSubDominManageByIds", cSubDominManageApi.DeleteCSubDominManageByIds) // 批量删除cSubDominManage表
		cSubDominManageRouter.PUT("updateCSubDominManage", cSubDominManageApi.UpdateCSubDominManage)              // 更新cSubDominManage表
	}
	{
		cSubDominManageRouterWithoutRecord.GET("findCSubDominManage", cSubDominManageApi.FindCSubDominManage)       // 根据ID获取cSubDominManage表
		cSubDominManageRouterWithoutRecord.GET("getCSubDominManageList", cSubDominManageApi.GetCSubDominManageList) // 获取cSubDominManage表列表
	}
	{
		cSubDominManageRouterWithoutAuth.GET("getCSubDominManagePublic", cSubDominManageApi.GetCSubDominManagePublic) // cSubDominManage表开放接口
	}
}
