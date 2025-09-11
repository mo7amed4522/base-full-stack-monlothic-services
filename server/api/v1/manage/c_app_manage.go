package manage

import (
	"encoding/json"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CAppManager struct{}

// CreateCPayConfig 创建cPayConfig表
func (cAppManager *CAppManager) CreateCApp(c *gin.Context) {
	var appCreateReq manage.CAppManage
	err := c.ShouldBindJSON(&appCreateReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 验证language_json是否为有效的JSON字符串
	var languageJsonObj interface{}
	if err := json.Unmarshal(appCreateReq.LanguageJson, &languageJsonObj); err != nil {
		response.FailWithMessage("language_json format is incorrect: "+err.Error(), c)
		return
	}

	// 直接转换为数据库模型
	appManage := manage.CAppManage{
		AppName:      appCreateReq.AppName,
		AppRemark:    appCreateReq.AppRemark,
		AppIcon:      appCreateReq.AppIcon,
		LanguageJson: appCreateReq.LanguageJson, // 将字符串转换为json.RawMessage
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		UpdatedBy:    &[]uint{uint(utils.GetUserID(c))}[0],
	}

	err = cAppManageService.CreateCApp(&appManage)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

// DeleteCPayConfig 删除cPayConfig表
func (cAppManager *CAppManager) DeleteApp(c *gin.Context) {
	id := c.Query("id")
	err := cAppManageService.DeleteCApp(id)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// UpdateCPayConfig 更新cPayConfig表
func (cAppManager *CAppManager) UpdateCPayConfig(c *gin.Context) {
	var cPayConfig manage.CAppManage
	err := c.ShouldBindJSON(&cPayConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cPayConfig.UpdatedAt = time.Now()
	err = cAppManageService.UpdateCApp(cPayConfig)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// FindCPayConfig 用id查询cPayConfig表
func (cAppManager *CAppManager) FindCPayConfig(c *gin.Context) {
	id := c.Query("id")
	recPayConfig, err := cAppManageService.GetCApp(id)
	if err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Error(err))
		response.FailWithMessage("Query failed:"+err.Error(), c)
		return
	}
	response.OkWithData(recPayConfig, c)
}

// GetCAppList 分页获取表列表
func (cAppManager *CAppManager) GetCAppList(c *gin.Context) {
	var pageInfo request.AppManageSearchRequest
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cAppManageService.GetCAppInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Error(err))
		response.FailWithMessage("Query failed:"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Query successful", c)
}

// GetCPayConfigPublic 不需要鉴权的cPayConfig表接口
func (cAppManager *CAppManager) GetCPayConfigPublic(c *gin.Context) {
	id := c.Query("id")
	var info manage.CAppManage
	global.GVA_DB.
		Preload("PromotionLinkInfo").
		Model(&manage.CAppManage{}).
		Where("id = ?", id).
		First(&info)

	response.OkWithDetailed(info, "Query successful", c)
}
