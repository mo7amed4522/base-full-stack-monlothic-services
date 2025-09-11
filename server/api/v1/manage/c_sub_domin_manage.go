package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage"
	manageReq "github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CSubDominManageApi struct{}

// CreateCSubDominManage 创建cSubDominManage表
// @Tags CSubDominManage
// @Summary 创建cSubDominManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CSubDominManage true "创建cSubDominManage表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cSubDominManage/createCSubDominManage [post]
func (cSubDominManageApi *CSubDominManageApi) CreateCSubDominManage(c *gin.Context) {
	var cSubDominManage manageReq.CSubDominManageAdd
	err := c.ShouldBindJSON(&cSubDominManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cSubDominManageService.CreateCSubDominManage(&cSubDominManage)
	if err != nil {
		global.GVA_LOG.Error("Create failed!", zap.Error(err))
		response.FailWithMessage("Create failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Create successful", c)
}

// DeleteCSubDominManage 删除cSubDominManage表
// @Tags CSubDominManage
// @Summary 删除cSubDominManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CSubDominManage true "删除cSubDominManage表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cSubDominManage/deleteCSubDominManage [delete]
func (cSubDominManageApi *CSubDominManageApi) DeleteCSubDominManage(c *gin.Context) {
	id := c.Query("id")
	err := cSubDominManageService.DeleteCSubDominManage(id)
	if err != nil {
		global.GVA_LOG.Error("Delete failed!", zap.Error(err))
		response.FailWithMessage("Delete failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Delete successful", c)
}

// DeleteCSubDominManageByIds 批量删除cSubDominManage表
// @Tags CSubDominManage
// @Summary 批量删除cSubDominManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cSubDominManage/deleteCSubDominManageByIds [delete]
func (cSubDominManageApi *CSubDominManageApi) DeleteCSubDominManageByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := cSubDominManageService.DeleteCSubDominManageByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("Batch delete failed!", zap.Error(err))
		response.FailWithMessage("Batch delete failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Batch delete successful", c)
}

// UpdateCSubDominManage 更新cSubDominManage表
// @Tags CSubDominManage
// @Summary 更新cSubDominManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CSubDominManage true "更新cSubDominManage表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cSubDominManage/updateCSubDominManage [put]
func (cSubDominManageApi *CSubDominManageApi) UpdateCSubDominManage(c *gin.Context) {
	var cSubDominManage manage.CSubDominManage
	err := c.ShouldBindJSON(&cSubDominManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cSubDominManageService.UpdateCSubDominManage(cSubDominManage)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// FindCSubDominManage 用id查询cSubDominManage表
// @Tags CSubDominManage
// @Summary 用id查询cSubDominManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query int true "用id查询cSubDominManage表"
// @Success 200 {object} response.Response{data=manage.CSubDominManage,msg=string} "查询成功"
// @Router /cSubDominManage/findCSubDominManage [get]
func (cSubDominManageApi *CSubDominManageApi) FindCSubDominManage(c *gin.Context) {
	id := c.Query("id")
	recSubDominManage, err := cSubDominManageService.GetCSubDominManage(id)
	if err != nil {
		global.GVA_LOG.Error("Find failed!", zap.Error(err))
		response.FailWithMessage("Find failed:"+err.Error(), c)
		return
	}
	response.OkWithData(recSubDominManage, c)
}

// GetCSubDominManageList 分页获取cSubDominManage表列表
// @Tags CSubDominManage
// @Summary 分页获取cSubDominManage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query manageReq.CSubDominManageSearch true "分页获取cSubDominManage表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cSubDominManage/getCSubDominManageList [get]
func (cSubDominManageApi *CSubDominManageApi) GetCSubDominManageList(c *gin.Context) {
	var pageInfo manageReq.CSubDominManageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cSubDominManageService.GetCSubDominManageInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Find failed!", zap.Error(err))
		response.FailWithMessage("Find failed:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Find successful", c)
}

// GetCSubDominManagePublic 不需要鉴权的cSubDominManage表接口
// @Tags CSubDominManage
// @Summary 不需要鉴权的cSubDominManage表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cSubDominManage/getCSubDominManagePublic [get]
func (cSubDominManageApi *CSubDominManageApi) GetCSubDominManagePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	cSubDominManageService.GetCSubDominManagePublic()
	response.OkWithDetailed(gin.H{
		"info": "No authentication required",
	}, "Find successful", c)
}
