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

// CreateCPageTemplate Create the cPageTemplate table
// @Tags CPageTemplate
// @Summary Create the cPageTemplate table
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CPageTemplate true "Create the cPageTemplate table"
// @Success 200 {object} response.Response{msg=string} "Created successfully"
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

// DeleteCPageTemplate Delete the cPageTemplate table
// @Tags CPageTemplate
// @Summary Delete the cPageTemplate table
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CPageTemplate true "Delete the cPageTemplate table"
// @Success 200 {object} response.Response{msg=string} "Deleted successfully"
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

// DeleteCPageTemplateByIds Batch Delete the cPageTemplate table
// @Tags CPageTemplate
// @Summary Batch Delete the cPageTemplate table
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "Batch deletion successful"
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

// UpdateCPageTemplate Update the cPageTemplate table
// @Tags CPageTemplate
// @Summary Update the cPageTemplate table
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CPageTemplate true "Update the cPageTemplate table"
// @Success 200 {object} response.Response{msg=string} "Update successful"
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

// FindCPageTemplate Query the cPageTemplate table using id
// @Tags CPageTemplate
// @Summary Query the cPageTemplate table using id
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query int true "Query the cPageTemplate table using id"
// @Success 200 {object} response.Response{data=manage.CPageTemplate,msg=string} "Query successful"
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

// GetCPageTemplateList Get cPageTemplate table list by paging
// @Tags CPageTemplate
// @Summary Get cPageTemplate table list by paging
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query manageReq.CPageTemplateSearch true "Get cPageTemplate table list by paging"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Get Success"
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

// GetCPageTemplatePublic cPageTemplate table interface that does not require authentication
// @Tags CPageTemplate
// @Summary cPageTemplate table interface that does not require authentication
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "Get Success"
// @Router /cPageTemplate/getCPageTemplatePublic [get]
func (cPageTemplateApi *CPageTemplateApi) GetCPageTemplatePublic(c *gin.Context) {
	// This interface does not require authentication
	// The example returns a fixed message interface. Generally, this interface is used for C-side services and requires implementing business logic by yourself.
	cPageTemplateService.GetCPageTemplatePublic()
	response.OkWithDetailed(gin.H{
		"info": "No authentication required",
	}, "Get successful", c)
}
