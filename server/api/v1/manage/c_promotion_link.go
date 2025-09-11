package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage"
	manageReq "github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CPromotionLinkApi struct{}

// CreateCPromotionLink 创建cPromotionLink表
// @Tags CPromotionLink
// @Summary 创建cPromotionLink表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CPromotionLink true "创建cPromotionLink表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cPromotionLink/createCPromotionLink [post]
func (cPromotionLinkApi *CPromotionLinkApi) CreateCPromotionLink(c *gin.Context) {
	var cPromotionLink manage.CPromotionLink
	err := c.ShouldBindJSON(&cPromotionLink)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cPromotionLinkService.CreateCPromotionLink(&cPromotionLink)
	if err != nil {
		global.GVA_LOG.Error("Create failed!", zap.Error(err))
		response.FailWithMessage("Create failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Create successful", c)
}

// DeleteCPromotionLink 删除cPromotionLink表
// @Tags CPromotionLink
// @Summary 删除cPromotionLink表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CPromotionLink true "删除cPromotionLink表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cPromotionLink/deleteCPromotionLink [delete]
func (cPromotionLinkApi *CPromotionLinkApi) DeleteCPromotionLink(c *gin.Context) {
	id := c.Query("id")
	err := cPromotionLinkService.DeleteCPromotionLink(id)
	if err != nil {
		global.GVA_LOG.Error("Delete failed!", zap.Error(err))
		response.FailWithMessage("Delete failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Delete successful", c)
}

// DeleteCPromotionLinkByIds 批量删除cPromotionLink表
// @Tags CPromotionLink
// @Summary 批量删除cPromotionLink表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cPromotionLink/deleteCPromotionLinkByIds [delete]
func (cPromotionLinkApi *CPromotionLinkApi) DeleteCPromotionLinkByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := cPromotionLinkService.DeleteCPromotionLinkByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("Batch delete failed!", zap.Error(err))
		response.FailWithMessage("Batch delete failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Batch delete successful", c)
}

// UpdateCPromotionLink 更新cPromotionLink表
// @Tags CPromotionLink
// @Summary 更新cPromotionLink表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CPromotionLink true "更新cPromotionLink表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cPromotionLink/updateCPromotionLink [put]
func (cPromotionLinkApi *CPromotionLinkApi) UpdateCPromotionLink(c *gin.Context) {
	var cPromotionLink manage.CPromotionLink
	err := c.ShouldBindJSON(&cPromotionLink)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cPromotionLinkService.UpdateCPromotionLink(cPromotionLink)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// FindCPromotionLink 用id查询cPromotionLink表
// @Tags CPromotionLink
// @Summary 用id查询cPromotionLink表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query int true "用id查询cPromotionLink表"
// @Success 200 {object} response.Response{data=manage.CPromotionLink,msg=string} "查询成功"
// @Router /cPromotionLink/findCPromotionLink [get]
func (cPromotionLinkApi *CPromotionLinkApi) FindCPromotionLink(c *gin.Context) {
	id := c.Query("id")
	recPromotionLink, err := cPromotionLinkService.GetCPromotionLink(id)
	if err != nil {
		global.GVA_LOG.Error("Find failed!", zap.Error(err))
		response.FailWithMessage("Find failed:"+err.Error(), c)
		return
	}
	response.OkWithData(recPromotionLink, c)
}

// GetCPromotionLinkList 分页获取cPromotionLink表列表
// @Tags CPromotionLink
// @Summary 分页获取cPromotionLink表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query manageReq.CPromotionLinkSearch true "分页获取cPromotionLink表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cPromotionLink/getCPromotionLinkList [get]
func (cPromotionLinkApi *CPromotionLinkApi) GetCPromotionLinkList(c *gin.Context) {
	var pageInfo manageReq.CPromotionLinkSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cPromotionLinkService.GetCPromotionLinkInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Find failed!", zap.Error(err))
		response.FailWithMessage("Find failed:"+err.Error(), c)
		return
	}

	var db = global.GVA_DB
	for k, link := range list {
		if link.DomainId != 0 {
			var domainURL string
			db.Model(&manage.CDominManage{}).Where("id = ?", link.DomainId).Pluck("domain", &domainURL)
			list[k].DomainURL = domainURL
		}
		if link.AppId != 0 {
			var appURL string
			db.Model(&manage.CAppManage{}).Where("id = ?", link.AppId).Pluck("app_name", &appURL)
			list[k].AppName = appURL
		}
		if link.TemplateId != 0 {
			var templateName string
			db.Model(&manage.CPageTemplate{}).Where("id = ?", link.TemplateId).Pluck("name", &templateName)
			list[k].TemplateName = templateName
		}
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Find successful", c)
}

// GetCPromotionLinkPublic 不需要鉴权的cPromotionLink表接口
// @Tags CPromotionLink
// @Summary 不需要鉴权的cPromotionLink表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cPromotionLink/getCPromotionLinkPublic [get]
func (cPromotionLinkApi *CPromotionLinkApi) GetCPromotionLinkPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	cPromotionLinkService.GetCPromotionLinkPublic()
	response.OkWithDetailed(gin.H{
		"info": "No authentication required",
	}, "Find successful", c)
}
