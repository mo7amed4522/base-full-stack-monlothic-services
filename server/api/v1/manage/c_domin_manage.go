package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage"
	manageReq "github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CDominManageApi struct{}

// CreateCDominManage 创建cDominManage表
// @Tags CDominManage
// @Summary 创建cDominManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CDominManage true "创建cDominManage表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cDominManage/createCDominManage [post]
func (cDominManageApi *CDominManageApi) CreateCDominManage(c *gin.Context) {
	var cDominManage manage.CDominManage
	err := c.ShouldBindJSON(&cDominManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := int(utils.GetUserID(c))
	cDominManage.CreatedBy = &uid
	err = cDominManageService.CreateCDominManage(&cDominManage)
	if err != nil {
		global.GVA_LOG.Error("Create failed!", zap.Error(err))
		response.FailWithMessage("Create failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Create successful", c)
}

// DeleteCDominManage 删除cDominManage表
// @Tags CDominManage
// @Summary 删除cDominManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CDominManage true "删除cDominManage表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cDominManage/deleteCDominManage [delete]
func (cDominManageApi *CDominManageApi) DeleteCDominManage(c *gin.Context) {
	id := c.Query("id")
	err := cDominManageService.DeleteCDominManage(id)
	if err != nil {
		global.GVA_LOG.Error("Delete failed!", zap.Error(err))
		response.FailWithMessage("Delete failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Delete successful", c)
}

// DeleteCDominManageByIds 批量删除cDominManage表
// @Tags CDominManage
// @Summary 批量删除cDominManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cDominManage/deleteCDominManageByIds [delete]
func (cDominManageApi *CDominManageApi) DeleteCDominManageByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := cDominManageService.DeleteCDominManageByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("Batch delete failed!", zap.Error(err))
		response.FailWithMessage("Batch delete failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Batch delete successful", c)
}

// UpdateCDominManage 更新cDominManage表
// @Tags CDominManage
// @Summary 更新cDominManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CDominManage true "更新cDominManage表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cDominManage/updateCDominManage [put]
func (cDominManageApi *CDominManageApi) UpdateCDominManage(c *gin.Context) {
	var cDominManage manage.CDominManage
	err := c.ShouldBindJSON(&cDominManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cDominManageService.UpdateCDominManage(cDominManage)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// FindCDominManage 用id查询cDominManage表
// @Tags CDominManage
// @Summary 用id查询cDominManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query int true "用id查询cDominManage表"
// @Success 200 {object} response.Response{data=manage.CDominManage,msg=string} "查询成功"
// @Router /cDominManage/findCDominManage [get]
func (cDominManageApi *CDominManageApi) FindCDominManage(c *gin.Context) {
	id := c.Query("id")
	recDominManage, err := cDominManageService.GetCDominManage(id)
	if err != nil {
		global.GVA_LOG.Error("Find failed!", zap.Error(err))
		response.FailWithMessage("Find failed:"+err.Error(), c)
		return
	}
	response.OkWithData(recDominManage, c)
}

// GetCDominManageList 分页获取cDominManage表列表
// @Tags CDominManage
// @Summary 分页获取cDominManage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query manageReq.CDominManageSearch true "分页获取cDominManage表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cDominManage/getCDominManageList [get]
func (cDominManageApi *CDominManageApi) GetCDominManageList(c *gin.Context) {
	var pageInfo manageReq.CDominManageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cDominManageService.GetCDominManageInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Get failed!", zap.Error(err))
		response.FailWithMessage("Get failed:"+err.Error(), c)
		return
	}
	for k, v := range list {
		var subList []manage.CSubDominManage
		global.GVA_DB.Model(&manage.CSubDominManage{}).
			Where("c_domin_id = ?", v.Id).
			Find(&subList)
		list[k].SubDomainList = subList
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Get successful", c)
}

// GetCDominManagePublic 不需要鉴权的cDominManage表接口
// @Tags CDominManage
// @Summary 不需要鉴权的cDominManage表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cDominManage/getCDominManagePublic [get]
func (cDominManageApi *CDominManageApi) GetCDominManagePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	cDominManageService.GetCDominManagePublic()
	response.OkWithDetailed(gin.H{
		"info": "No authentication required",
	}, "Get successful", c)
}

// VisibleCDominManage 隐藏cDominManage
// @Tags CDominManage
// @Summary 隐藏cDominManage
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CDominManage true "隐藏cDominManage"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cDominManage/visibleCDominManage [put]
func (cDominManageApi *CDominManageApi) VisibleCDominManage(c *gin.Context) {
	var cDominManage manage.CDominManage
	err := c.ShouldBindJSON(&cDominManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	err = cDominManageService.VisibleCDominManage(uid, cDominManage)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Update successful", c)
}
