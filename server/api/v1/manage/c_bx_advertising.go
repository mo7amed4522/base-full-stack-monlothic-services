package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CBxAdvertisingApi struct{}

// GetAdvertisingList 获取广告列表
func (cBxAdvertisingApi *CBxAdvertisingApi) GetAdvertisingList(c *gin.Context) {
	var pageInfo request.BxAdvertisingSearchRequest
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage("Parameter error: "+err.Error(), c)
		return
	}

	list, total, err := cBxAdvertisingService.GetAdvertisingList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Query advertising list failed!", zap.Error(err))
		response.FailWithMessage("Query failed: "+err.Error(), c)
		return
	}

	// 使用分页工具构建响应
	responseData := utils.BuildPaginateResponse(list, total, pageInfo.Page, pageInfo.PageSize)

	response.OkWithDetailed(responseData, "Query advertising list successful", c)
}

// CreateAdvertising 创建广告
func (cBxAdvertisingApi *CBxAdvertisingApi) CreateAdvertising(c *gin.Context) {
	var req request.BxAdvertisingCreateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("Parameter error: "+err.Error(), c)
		return
	}

	err = cBxAdvertisingService.CreateAdvertising(req)
	if err != nil {
		global.GVA_LOG.Error("Create advertising failed!", zap.Error(err))
		response.FailWithMessage("Create failed: "+err.Error(), c)
		return
	}

	response.OkWithMessage("Create advertising successful", c)
}

// UpdateAdvertising 更新广告
func (cBxAdvertisingApi *CBxAdvertisingApi) UpdateAdvertising(c *gin.Context) {
	var req request.BxAdvertisingUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("Parameter error: "+err.Error(), c)
		return
	}

	err = cBxAdvertisingService.UpdateAdvertising(req)
	if err != nil {
		global.GVA_LOG.Error("Update advertising failed!", zap.Error(err))
		response.FailWithMessage("Update failed: "+err.Error(), c)
		return
	}

	response.OkWithMessage("Update advertising successful", c)
}

// DeleteAdvertising 删除广告
func (cBxAdvertisingApi *CBxAdvertisingApi) DeleteAdvertising(c *gin.Context) {
	var req request.BxAdvertisingDeleteRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage("Parameter error: "+err.Error(), c)
		return
	}

	err = cBxAdvertisingService.DeleteAdvertising(req.AdId)
	if err != nil {
		global.GVA_LOG.Error("Delete advertising failed!", zap.Error(err))
		response.FailWithMessage("Delete failed: "+err.Error(), c)
		return
	}

	response.OkWithMessage("Delete advertising successful", c)
}

// UpdateDictValue 更新字典值
func (cBxAdvertisingApi *CBxAdvertisingApi) UpdateDictValue(c *gin.Context) {
	var req request.BxDictUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("Parameter error: "+err.Error(), c)
		return
	}

	err = cBxAdvertisingService.UpdateDictValue(req)
	if err != nil {
		global.GVA_LOG.Error("Update dict value failed!", zap.Error(err))
		response.FailWithMessage("Update failed: "+err.Error(), c)
		return
	}

	response.OkWithMessage("Update dict value successful", c)
}
