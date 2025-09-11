package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CBxAdvertisingRouter struct{}

// InitBxAdvertisingRouter 初始化广告管理路由
func (s *CAppManageRouter) InitBxAdvertisingRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cBxAdvertisingRouter := Router.Group("advertising").Use(middleware.OperationRecord())
	{
		cBxAdvertisingRouter.GET("getAdvertisingList", cBxAdvertisingApi.GetAdvertisingList)
		cBxAdvertisingRouter.POST("createAdvertising", cBxAdvertisingApi.CreateAdvertising)
		cBxAdvertisingRouter.POST("updateAdvertising", cBxAdvertisingApi.UpdateAdvertising)
		cBxAdvertisingRouter.POST("deleteAdvertising", cBxAdvertisingApi.DeleteAdvertising)
		cBxAdvertisingRouter.POST("updateDictValue", cBxAdvertisingApi.UpdateDictValue)
	}
}
