package manage

import (
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage"
	manageReq "github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
)

type CSubDominManageService struct{}

// CreateCSubDominManage 创建cSubDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cSubDominManageService *CSubDominManageService) CreateCSubDominManage(cSubDominManage *manageReq.CSubDominManageAdd) (err error) {
	if len(cSubDominManage.SubDomains) > 3 {
		return errors.New("The number of subdomains exceeds the limit")
	}

	var (
		list []*manage.CSubDominManage
		db   = global.GVA_DB.Begin()
		n    = time.Now()
	)

	err = db.Model(&manage.CSubDominManage{}).Where("c_domin_id = ?", cSubDominManage.CDominId).Delete(&list).Error
	if err != nil {
		db.Rollback()
		return err
	}

	for _, domain := range cSubDominManage.SubDomains {
		list = append(list, &manage.CSubDominManage{
			CDominId:  *cSubDominManage.CDominId,
			SubDomain: domain,
			CreatedAt: &n,
		})
	}
	err = db.Create(&list).Error
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

// DeleteCSubDominManage 删除cSubDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cSubDominManageService *CSubDominManageService) DeleteCSubDominManage(id string) (err error) {
	err = global.GVA_DB.Delete(&manage.CSubDominManage{}, "id = ?", id).Error
	return err
}

// DeleteCSubDominManageByIds 批量删除cSubDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cSubDominManageService *CSubDominManageService) DeleteCSubDominManageByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]manage.CSubDominManage{}, "id in ?", ids).Error
	return err
}

// UpdateCSubDominManage 更新cSubDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cSubDominManageService *CSubDominManageService) UpdateCSubDominManage(cSubDominManage manage.CSubDominManage) (err error) {
	err = global.GVA_DB.Model(&manage.CSubDominManage{}).Where("id = ?", cSubDominManage.Id).Updates(&cSubDominManage).Error
	return err
}

// GetCSubDominManage 根据id获取cSubDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cSubDominManageService *CSubDominManageService) GetCSubDominManage(id string) (cSubDominManage manage.CSubDominManage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cSubDominManage).Error
	return
}

// GetCSubDominManageInfoList 分页获取cSubDominManage表记录
// Author [yourname](https://github.com/yourname)
func (cSubDominManageService *CSubDominManageService) GetCSubDominManageInfoList(info manageReq.CSubDominManageSearch) (list []manage.CSubDominManage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.CSubDominManage{})
	var cSubDominManages []manage.CSubDominManage
	// 如果有条件搜索 下方会自动创建搜索语句
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

	err = db.Find(&cSubDominManages).Error
	return cSubDominManages, total, err
}
func (cSubDominManageService *CSubDominManageService) GetCSubDominManagePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
