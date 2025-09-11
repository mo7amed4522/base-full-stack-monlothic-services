package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage"
	manageReq "github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CPageTemplateApi struct{}

// CreateCPageTemplate 创建cPageTemplate表
// @Tags CPageTemplate
// @Summary 创建cPageTemplate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CPageTemplate true "创建cPageTemplate表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cPageTemplate/createCPageTemplate [post]
func (cPageTemplateApi *CPageTemplateApi) CreateCPageTemplate(c *gin.Context) {
	var cPageTemplate manage.CPageTemplate
	err := c.ShouldBindJSON(&cPageTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cPageTemplateService.CreateCPageTemplate(&cPageTemplate)
	if err != nil {
		global.GVA_LOG.Error("Create failed!", zap.Error(err))
		response.FailWithMessage("Create failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Create successful", c)
}

// DeleteCPageTemplate 删除cPageTemplate表
// @Tags CPageTemplate
// @Summary 删除cPageTemplate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CPageTemplate true "删除cPageTemplate表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cPageTemplate/deleteCPageTemplate [delete]
func (cPageTemplateApi *CPageTemplateApi) DeleteCPageTemplate(c *gin.Context) {
	id := c.Query("id")
	err := cPageTemplateService.DeleteCPageTemplate(id)
	if err != nil {
		global.GVA_LOG.Error("Delete failed!", zap.Error(err))
		response.FailWithMessage("Delete failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Delete successful", c)
}

// DeleteCPageTemplateByIds 批量删除cPageTemplate表
// @Tags CPageTemplate
// @Summary 批量删除cPageTemplate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cPageTemplate/deleteCPageTemplateByIds [delete]
func (cPageTemplateApi *CPageTemplateApi) DeleteCPageTemplateByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := cPageTemplateService.DeleteCPageTemplateByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("Delete failed!", zap.Error(err))
		response.FailWithMessage("Delete failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Delete successful", c)
}

// UpdateCPageTemplate 更新cPageTemplate表
// @Tags CPageTemplate
// @Summary 更新cPageTemplate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CPageTemplate true "更新cPageTemplate表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cPageTemplate/updateCPageTemplate [put]
func (cPageTemplateApi *CPageTemplateApi) UpdateCPageTemplate(c *gin.Context) {
	var cPageTemplate manage.CPageTemplate
	err := c.ShouldBindJSON(&cPageTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cPageTemplateService.UpdateCPageTemplate(cPageTemplate)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// FindCPageTemplate 用id查询cPageTemplate表
// @Tags CPageTemplate
// @Summary 用id查询cPageTemplate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query int true "用id查询cPageTemplate表"
// @Success 200 {object} response.Response{data=manage.CPageTemplate,msg=string} "查询成功"
// @Router /cPageTemplate/findCPageTemplate [get]
func (cPageTemplateApi *CPageTemplateApi) FindCPageTemplate(c *gin.Context) {
	id := c.Query("id")
	recPageTemplate, err := cPageTemplateService.GetCPageTemplate(id)
	if err != nil {
		global.GVA_LOG.Error("Get failed!", zap.Error(err))
		response.FailWithMessage("Get failed:"+err.Error(), c)
		return
	}
	response.OkWithData(recPageTemplate, c)
}

// GetCPageTemplateList 分页获取cPageTemplate表列表
// @Tags CPageTemplate
// @Summary 分页获取cPageTemplate表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query manageReq.CPageTemplateSearch true "分页获取cPageTemplate表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cPageTemplate/getCPageTemplateList [get]
func (cPageTemplateApi *CPageTemplateApi) GetCPageTemplateList(c *gin.Context) {
	var pageInfo manageReq.CPageTemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cPageTemplateService.GetCPageTemplateInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Get failed!", zap.Error(err))
		response.FailWithMessage("Get failed:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Get successful", c)
}

// GetCPageTemplatePublic 不需要鉴权的cPageTemplate表接口
// @Tags CPageTemplate
// @Summary 不需要鉴权的cPageTemplate表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cPageTemplate/getCPageTemplatePublic [get]
func (cPageTemplateApi *CPageTemplateApi) GetCPageTemplatePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	cPageTemplateService.GetCPageTemplatePublic()
	response.OkWithDetailed(gin.H{
		"info": "No authentication required",
	}, "Get successful", c)
}
