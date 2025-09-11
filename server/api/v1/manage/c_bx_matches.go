package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CBxMatchesApi struct{}

// GetMatchesList 获取比赛列表
func (cBxMatchesApi *CBxMatchesApi) GetMatchesList(c *gin.Context) {
	var pageInfo request.BxMatchesSearchRequest
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage("Parameter error: "+err.Error(), c)
		return
	}

	list, total, err := cBxMatchesService.GetBxMatchesList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Get matches list failed!", zap.Error(err))
		response.FailWithMessage("Get failed: "+err.Error(), c)
		return
	}

	// 使用分页工具构建响应
	responseData := utils.BuildPaginateResponse(list, total, pageInfo.Page, pageInfo.PageSize)

	response.OkWithDetailed(responseData, "Get matches list successful", c)
}

// BatchSetHot 批量设置热门
func (cBxMatchesApi *CBxMatchesApi) BatchSetHot(c *gin.Context) {
	var batchReq request.BatchMatchesRequest
	err := c.ShouldBindJSON(&batchReq)
	if err != nil {
		response.FailWithMessage("Parameter error: "+err.Error(), c)
		return
	}

	err = cBxMatchesService.BatchSetHot(batchReq)
	if err != nil {
		global.GVA_LOG.Error("Batch set hot failed!", zap.Error(err))
		response.FailWithMessage("Set failed: "+err.Error(), c)
		return
	}

	response.OkWithMessage("Batch set hot successful", c)
}

// BatchSetTop 批量置顶
func (cBxMatchesApi *CBxMatchesApi) BatchSetTop(c *gin.Context) {
	var batchReq request.BatchMatchesRequest
	err := c.ShouldBindJSON(&batchReq)
	if err != nil {
		response.FailWithMessage("Parameter error: "+err.Error(), c)
		return
	}

	err = cBxMatchesService.BatchSetTop(batchReq)
	if err != nil {
		global.GVA_LOG.Error("Batch set top failed!", zap.Error(err))
		response.FailWithMessage("Set failed: "+err.Error(), c)
		return
	}

	response.OkWithMessage("Batch set top successful", c)
}
