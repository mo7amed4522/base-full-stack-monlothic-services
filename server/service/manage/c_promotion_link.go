package manage

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage"
	manageReq "github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
	"github.com/google/uuid"
)

type CPromotionLinkService struct{}

// CreateCPromotionLink 创建cPromotionLink表记录
// Author [yourname](https://github.com/yourname)
func (cPromotionLinkService *CPromotionLinkService) CreateCPromotionLink(cPromotionLink *manage.CPromotionLink) (err error) {
	prefix := "https://%s/%s/%s_download.html?link_id=%d"
	err = global.GVA_DB.Create(cPromotionLink).Error
	if err != nil {
		return err
	}
	// 查询domain 的url地址
	var domainInfo manage.CDominManage
	global.GVA_DB.Model(&manage.CDominManage{}).Where("id = ?", cPromotionLink.DomainId).Find(&domainInfo)
	uid := uuid.NewString()
	cPromotionLink.AccessUrl = fmt.Sprintf(prefix, *domainInfo.Domain, uid, uid, *cPromotionLink.Id)
	err = global.GVA_DB.Save(&cPromotionLink).Error
	return err
}

// DeleteCPromotionLink 删除cPromotionLink表记录
// Author [yourname](https://github.com/yourname)
func (cPromotionLinkService *CPromotionLinkService) DeleteCPromotionLink(id string) (err error) {
	err = global.GVA_DB.Delete(&manage.CPromotionLink{}, "id = ?", id).Error
	return err
}

// DeleteCPromotionLinkByIds 批量删除cPromotionLink表记录
// Author [yourname](https://github.com/yourname)
func (cPromotionLinkService *CPromotionLinkService) DeleteCPromotionLinkByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]manage.CPromotionLink{}, "id in ?", ids).Error
	return err
}

// UpdateCPromotionLink 更新cPromotionLink表记录
// Author [yourname](https://github.com/yourname)
func (cPromotionLinkService *CPromotionLinkService) UpdateCPromotionLink(cPromotionLink manage.CPromotionLink) (err error) {
	err = global.GVA_DB.Model(&manage.CPromotionLink{}).Where("id = ?", cPromotionLink.Id).Updates(&cPromotionLink).Error
	return err
}

// GetCPromotionLink 根据id获取cPromotionLink表记录
// Author [yourname](https://github.com/yourname)
func (cPromotionLinkService *CPromotionLinkService) GetCPromotionLink(id string) (cPromotionLink manage.CPromotionLink, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cPromotionLink).Error
	return
}

// GetCPromotionLinkInfoList 分页获取cPromotionLink表记录
// Author [yourname](https://github.com/yourname)
func (cPromotionLinkService *CPromotionLinkService) GetCPromotionLinkInfoList(info manageReq.CPromotionLinkSearch) (list []manage.CPromotionLink, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.CPromotionLink{})
	var cPromotionLinks []manage.CPromotionLink
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name = ?", *info.Name)
	}
	if info.AppId != nil {
		db = db.Where("app_id = ?", *info.AppId)
	}
	if info.TemplateId != nil {
		db = db.Where("template_id = ?", *info.TemplateId)
	}
	if info.IsAllowEditUrl != nil {
		db = db.Where("is_allow_edit_url = ?", *info.IsAllowEditUrl)
	}
	if info.DomainId != nil {
		db = db.Where("domain_id = ?", *info.DomainId)
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

	err = db.Find(&cPromotionLinks).Error
	return cPromotionLinks, total, err
}
func (cPromotionLinkService *CPromotionLinkService) GetCPromotionLinkPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
