package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage"
	manageReq "github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
)

type CAppManageService struct{}

// CreateCApp 创建应用管理记录
func (cAppManageService *CAppManageService) CreateCApp(appManage *manage.CAppManage) (err error) {
	err = global.GVA_DB.Create(appManage).Error
	return err
}

// DeleteCApp 删除应用管理记录
func (cAppManageService *CAppManageService) DeleteCApp(id string) (err error) {
	err = global.GVA_DB.Delete(&manage.CAppManage{}, "id = ?", id).Error
	return err
}

// UpdateCApp 更新应用管理记录
func (cAppManageService *CAppManageService) UpdateCApp(appManage manage.CAppManage) (err error) {
	err = global.GVA_DB.Model(&manage.CAppManage{}).Where("id = ?", appManage.ID).Updates(&appManage).Error
	return err
}

// GetCApp 根据id获取应用管理记录
func (cAppManageService *CAppManageService) GetCApp(id any) (appManage manage.CAppManage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appManage).Error
	return
}

// GetCAppInfoList 分页获取应用管理记录
func (cAppManageService *CAppManageService) GetCAppInfoList(req manageReq.AppManageSearchRequest) (list []manage.CAppManage, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	db := global.GVA_DB.Model(&manage.CAppManage{})
	var appManages []manage.CAppManage

	if req.AppName != "" {
		db = db.Where("app_name LIKE ?", "%"+req.AppName+"%")
	}

	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	db = db.Order("id desc")
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&appManages).Error
	return appManages, total, err
}

// GetRandomCPayPool 随机获取
func (cAppManageService *CAppManageService) GetRandomCPayPool() (appManage manage.CAppManage, err error) {
	err = global.GVA_DB.Order("RAND()").First(&appManage).Error
	return
}
