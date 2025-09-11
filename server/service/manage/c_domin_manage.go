package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage"
	manageReq "github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
	"time"
)

type CDominManageService struct{}

// CreateCDominManage 创建cDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cDominManageService *CDominManageService) CreateCDominManage(cDominManage *manage.CDominManage) (err error) {
	err = global.GVA_DB.Create(cDominManage).Error
	return err
}

// DeleteCDominManage 删除cDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cDominManageService *CDominManageService) DeleteCDominManage(id string) (err error) {
	err = global.GVA_DB.Delete(&manage.CDominManage{}, "id = ?", id).Error
	return err
}

// DeleteCDominManageByIds 批量删除cDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cDominManageService *CDominManageService) DeleteCDominManageByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]manage.CDominManage{}, "id in ?", ids).Error
	return err
}

// UpdateCDominManage 更新cDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cDominManageService *CDominManageService) UpdateCDominManage(cDominManage manage.CDominManage) (err error) {
	err = global.GVA_DB.Model(&manage.CDominManage{}).Where("id = ?", cDominManage.Id).Updates(&cDominManage).Error
	return err
}

// GetCDominManage 根据id获取cDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cDominManageService *CDominManageService) GetCDominManage(id string) (cDominManage manage.CDominManage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cDominManage).Error
	return
}

// GetCDominManageInfoList 分页获取cDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cDominManageService *CDominManageService) GetCDominManageInfoList(info manageReq.CDominManageSearch) (list []manage.CDominManage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.CDominManage{})
	var cDominManages []manage.CDominManage
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	if info.Domain != "" {
		db = db.Where("domain = ?", info.Domain)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["created_at"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cDominManages).Error
	return cDominManages, total, err
}
func (cDominManageService *CDominManageService) GetCDominManagePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// VisibleCDominManage 更新cDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cDominManageService *CDominManageService) VisibleCDominManage(uid uint, cDominManage manage.CDominManage) (err error) {
	return global.GVA_DB.Model(&manage.CDominManage{}).
		Where("id = ?", cDominManage.Id).
		Updates(map[string]any{
			"visible":    cDominManage.Visible,
			"updated_at": time.Now(),
			"updated_by": uid,
		}).Error
}
