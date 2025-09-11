package manage

import (
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bxapi"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type CBxAdvertisingService struct{}

// GetAdvertisingList 获取广告列表
func (cBxAdvertisingService *CBxAdvertisingService) GetAdvertisingList(req request.BxAdvertisingSearchRequest) (advertisingResponse []response.BxAdvertisingResponse, total int64, err error) {
	if global.GVA_DB == nil {
		return nil, 0, fmt.Errorf("db not init")
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 构建查询条件
	query := global.GVA_DB.Model(&bxapi.BxAdvertising{}).
		Where("deleted_at IS NULL")

	// 广告ID筛选
	if req.AdId != "" {
		query = query.Where("ad_id = ?", req.AdId)
	}

	// 标题搜索
	if req.Title != "" {
		query = query.Where("title LIKE ?", "%"+req.Title+"%")
	}

	// 广告商搜索
	if req.Advertiser != "" {
		query = query.Where("advertiser LIKE ?", "%"+req.Advertiser+"%")
	}

	// 站点筛选
	if req.SiteId != "" {
		query = query.Where("site_id = ?", req.SiteId)
	}

	// 状态筛选
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	// 时间范围筛选
	if req.StartTime != "" {
		query = query.Where("created_at >= ?", req.StartTime)
	}
	if req.EndTime != "" {
		query = query.Where("created_at <= ?", req.EndTime)
	}

	// 预加载位置信息，避免 N+1 查询
	query = query.Preload("Positions")

	// 页面类型筛选 - 通过关联表筛选
	if req.PageType != "" {
		query = query.Joins("JOIN bx_advertising_position ON bx_advertising.ad_id = bx_advertising_position.ad_id").
			Where("bx_advertising_position.page_type = ?", req.PageType).
			Group("bx_advertising.ad_id")
	}

	// 使用分页工具进行查询
	var advertisements []bxapi.BxAdvertising
	paginateResult, err := utils.PaginateWithCustomQuery(
		query,                                // count查询
		query.Order("sort DESC, ad_id DESC"), // 数据查询，按排序和ID倒序
		req.Page, req.PageSize, &advertisements,
	)
	if err != nil {
		return advertisingResponse, 0, err
	}

	advertisingResponse = cBxAdvertisingService.buildAdvertisingResponse(advertisements)

	return advertisingResponse, paginateResult.Total, nil
}

// buildAdvertisingResponseFromPreloaded 广告响应数据
func (cBxAdvertisingService *CBxAdvertisingService) buildAdvertisingResponse(advertisements []bxapi.BxAdvertising) []response.BxAdvertisingResponse {
	var advertisingResponse []response.BxAdvertisingResponse

	for _, ad := range advertisements {
		// 从预加载的关联数据中提取位置数组
		var positionIds []int
		for _, pos := range ad.Positions {
			positionIds = append(positionIds, int(pos.Position))
		}

		adResponse := response.BxAdvertisingResponse{
			AdId:          ad.AdId,
			SiteId:        ad.SiteId,
			Title:         ad.Title,
			Advertiser:    ad.Advertiser,
			AndroidUrl:    ad.AndroidUrl,
			IosUrl:        ad.IosUrl,
			H5Url:         ad.H5Url,
			ImgUrl:        ad.ImgUrl,
			CornerMarkUrl: ad.CornerMarkUrl,
			Description:   ad.Description,
			JumpType:      ad.JumpType,
			ClickCount:    ad.ClickCount,
			Sort:          ad.Sort,
			Status:        ad.Status,
			CreatedAt:     ad.CreatedAt,
			UpdatedAt:     ad.UpdatedAt,
			Positions:     positionIds,
			PageType:      ad.PageType,
		}
		advertisingResponse = append(advertisingResponse, adResponse)
	}

	return advertisingResponse
}

// CreateAdvertising 创建广告
func (cBxAdvertisingService *CBxAdvertisingService) CreateAdvertising(req request.BxAdvertisingCreateRequest) error {
	if global.GVA_DB == nil {
		return fmt.Errorf("db not init")
	}

	// 开始数据库事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建广告记录
	advertising := bxapi.BxAdvertising{
		SiteId:        req.SiteId,
		Title:         req.Title,
		Advertiser:    req.Advertiser,
		AndroidUrl:    req.AndroidUrl,
		IosUrl:        req.IosUrl,
		H5Url:         req.H5Url,
		ImgUrl:        req.ImgUrl,
		CornerMarkUrl: req.CornerMarkUrl,
		Description:   req.Description,
		JumpType:      req.JumpType,
		Sort:          req.Sort,
		Status:        req.Status,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := tx.Create(&advertising).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("创建广告失败: %v", err)
	}

	// 创建位置关联记录
	for _, position := range req.Position {
		positionRecord := bxapi.BxAdvertisingPosition{
			AdId:     advertising.AdId,
			Position: position,
			PageType: req.PageType,
		}
		if err := tx.Create(&positionRecord).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("创建位置关联失败: %v", err)
		}
	}

	return tx.Commit().Error
}

// UpdateAdvertising 更新广告
func (cBxAdvertisingService *CBxAdvertisingService) UpdateAdvertising(req request.BxAdvertisingUpdateRequest) error {
	if global.GVA_DB == nil {
		return fmt.Errorf("db not init")
	}

	// 开始数据库事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 检查广告是否存在
	var existing bxapi.BxAdvertising
	if err := tx.Where("ad_id = ? AND deleted_at IS NULL", req.AdId).First(&existing).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("广告不存在")
		}
		return fmt.Errorf("查询广告失败: %v", err)
	}

	// 更新广告信息
	updateData := map[string]interface{}{
		"site_id":         req.SiteId,
		"title":           req.Title,
		"advertiser":      req.Advertiser,
		"android_url":     req.AndroidUrl,
		"ios_url":         req.IosUrl,
		"h5_url":          req.H5Url,
		"img_url":         req.ImgUrl,
		"corner_mark_url": req.CornerMarkUrl,
		"description":     req.Description,
		"jump_type":       req.JumpType,
		"sort":            req.Sort,
		"status":          req.Status,
		"updated_at":      time.Now(),
	}

	if err := tx.Model(&bxapi.BxAdvertising{}).Where("ad_id = ?", req.AdId).Updates(updateData).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新广告失败: %v", err)
	}

	// 删除旧的位置关联
	if err := tx.Where("ad_id = ?", req.AdId).Delete(&bxapi.BxAdvertisingPosition{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除旧位置关联失败: %v", err)
	}

	// 创建新的位置关联记录
	for _, position := range req.Position {
		positionRecord := bxapi.BxAdvertisingPosition{
			AdId:     req.AdId,
			Position: position,
			PageType: req.PageType,
		}
		if err := tx.Create(&positionRecord).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("创建位置关联失败: %v", err)
		}
	}

	return tx.Commit().Error
}

// DeleteAdvertising 删除广告（软删除）
func (cBxAdvertisingService *CBxAdvertisingService) DeleteAdvertising(adId uint64) error {
	if global.GVA_DB == nil {
		return fmt.Errorf("db not init")
	}

	// 开始数据库事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 检查广告是否存在
	var existing bxapi.BxAdvertising
	if err := tx.Where("ad_id = ? AND deleted_at IS NULL", adId).First(&existing).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("广告不存在")
		}
		return fmt.Errorf("查询广告失败: %v", err)
	}

	// 软删除广告
	now := time.Now()
	if err := tx.Model(&bxapi.BxAdvertising{}).Where("ad_id = ?", adId).Update("deleted_at", now).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除广告失败: %v", err)
	}

	// 删除位置关联记录
	if err := tx.Where("ad_id = ?", adId).Delete(&bxapi.BxAdvertisingPosition{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除位置关联失败: %v", err)
	}

	return tx.Commit().Error
}

// UpdateDictValue 更新字典值
func (cBxAdvertisingService *CBxAdvertisingService) UpdateDictValue(req request.BxDictUpdateRequest) error {
	if global.GVA_DB == nil {
		return fmt.Errorf("db not init")
	}

	// 定义类型映射
	typeMapping := map[string]string{
		"android": constants.DICT_TYPE_ANDROID_AD_ENABLE,
		"ios":     constants.DICT_TYPE_IOS_AD_ENABLE,  
		"h5":      constants.DICT_TYPE_H5_AD_ENABLE,
		"ad":      constants.DICT_TYPE_AD_ENABLE,
	}

	// 获取对应的字典类型
	dictType, exists := typeMapping[req.Type]
	if !exists {
		return fmt.Errorf("不支持的类型: %s", req.Type)
	}

	// 更新字典值
	result := global.GVA_DB.Table("bx_dict").
		Where("type = ?", dictType).
		Update("value", req.Value)

	if result.Error != nil {
		return fmt.Errorf("更新字典值失败: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到对应的字典记录")
	}

	return nil
}
