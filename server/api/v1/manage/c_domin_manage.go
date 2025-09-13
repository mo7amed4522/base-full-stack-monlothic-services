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

// CreateCDominManage Create the cDominManage table
// @Tags CDominManage
// @Summary Create the cDominManage table
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CDominManage true "Create the cDominManage table"
// @Success 200 {object} response.Response{msg=string} "Created successfully"
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

// DeleteCDominManage Delete the cDominManage table
// @Tags CDominManage
// @Summary Delete the cDominManage table
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CDominManage true "Delete the cDominManage table"
// @Success 200 {object} response.Response{msg=string} "Deleted successfully"
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

// DeleteCDominManageByIds Delete the cDominManage table in batches
// @Tags CDominManage
// @Summary Delete the cDominManage table in batches
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "Batch deletion successful"
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

// UpdateCDominManage Update the cDominManage table
// @Tags CDominManage
// @Summary Update the cDominManage table
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CDominManage true "Update the cDominManage table"
// @Success 200 {object} response.Response{msg=string} "Update successful"
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

// FindCDominManage Query the cDominManage table using id
// @Tags CDominManage
// @Summary Query the cDominManage table using id
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query int true "Query the cDominManage table using id"
// @Success 200 {object} response.Response{data=manage.CDominManage,msg=string} "Query successful"
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

// GetCDominManageList Get the cDominManage table list by paging
// @Tags CDominManage
// @Summary Get the cDominManage table list by paging
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query manageReq.CDominManageSearch true "Get the cDominManage table list by paging"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Get Success"
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

// GetCDominManagePubliccDominManage table interface that does not require authentication
// @Tags CDominManage
// @SummarycDominManage table interface that does not require authentication
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "Get Success"
// @Router /cDominManage/getCDominManagePublic [get]
func (cDominManageApi *CDominManageApi) GetCDominManagePublic(c *gin.Context) {
	// This interface does not require authentication
	// The example returns a fixed message interface. Generally, this interface is used for C-side services and requires implementing business logic by yourself.
	cDominManageService.GetCDominManagePublic()
	response.OkWithDetailed(gin.H{
		"info": "No authentication required",
	}, "Get successful", c)
}

// VisibleCDominManage Hide cDominManage
// @Tags CDominManage
// @Summary Hide cDominManage
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manage.CDominManage true "Hide cDominManage"
// @Success 200 {object} response.Response{msg=string} "Update successful"
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
