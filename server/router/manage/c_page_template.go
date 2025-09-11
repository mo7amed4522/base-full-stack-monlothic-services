package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CPageTemplateRouter struct {}

// InitCPageTemplateRouter 初始化 cPageTemplate表 路由信息
func (s *CPageTemplateRouter) InitCPageTemplateRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	cPageTemplateRouter := Router.Group("cPageTemplate").Use(middleware.OperationRecord())
	cPageTemplateRouterWithoutRecord := Router.Group("cPageTemplate")
	cPageTemplateRouterWithoutAuth := PublicRouter.Group("cPageTemplate")
	{
		cPageTemplateRouter.POST("createCPageTemplate", cPageTemplateApi.CreateCPageTemplate)   // 新建cPageTemplate表
		cPageTemplateRouter.DELETE("deleteCPageTemplate", cPageTemplateApi.DeleteCPageTemplate) // 删除cPageTemplate表
		cPageTemplateRouter.DELETE("deleteCPageTemplateByIds", cPageTemplateApi.DeleteCPageTemplateByIds) // 批量删除cPageTemplate表
		cPageTemplateRouter.PUT("updateCPageTemplate", cPageTemplateApi.UpdateCPageTemplate)    // 更新cPageTemplate表
	}
	{
		cPageTemplateRouterWithoutRecord.GET("findCPageTemplate", cPageTemplateApi.FindCPageTemplate)        // 根据ID获取cPageTemplate表
		cPageTemplateRouterWithoutRecord.GET("getCPageTemplateList", cPageTemplateApi.GetCPageTemplateList)  // 获取cPageTemplate表列表
	}
	{
	    cPageTemplateRouterWithoutAuth.GET("getCPageTemplatePublic", cPageTemplateApi.GetCPageTemplatePublic)  // cPageTemplate表开放接口
	}
}
