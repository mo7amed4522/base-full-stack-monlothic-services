
package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage"
    manageReq "github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
)

type CPageTemplateService struct {}
// CreateCPageTemplate 创建cPageTemplate表记录
// Author [yourname](https://github.com/yourname)
func (cPageTemplateService *CPageTemplateService) CreateCPageTemplate(cPageTemplate *manage.CPageTemplate) (err error) {
	err = global.GVA_DB.Create(cPageTemplate).Error
	return err
}

// DeleteCPageTemplate 删除cPageTemplate表记录
// Author [yourname](https://github.com/yourname)
func (cPageTemplateService *CPageTemplateService)DeleteCPageTemplate(id string) (err error) {
	err = global.GVA_DB.Delete(&manage.CPageTemplate{},"id = ?",id).Error
	return err
}

// DeleteCPageTemplateByIds 批量删除cPageTemplate表记录
// Author [yourname](https://github.com/yourname)
func (cPageTemplateService *CPageTemplateService)DeleteCPageTemplateByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]manage.CPageTemplate{},"id in ?",ids).Error
	return err
}

// UpdateCPageTemplate 更新cPageTemplate表记录
// Author [yourname](https://github.com/yourname)
func (cPageTemplateService *CPageTemplateService)UpdateCPageTemplate(cPageTemplate manage.CPageTemplate) (err error) {
	err = global.GVA_DB.Model(&manage.CPageTemplate{}).Where("id = ?",cPageTemplate.Id).Updates(&cPageTemplate).Error
	return err
}

// GetCPageTemplate 根据id获取cPageTemplate表记录
// Author [yourname](https://github.com/yourname)
func (cPageTemplateService *CPageTemplateService)GetCPageTemplate(id string) (cPageTemplate manage.CPageTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cPageTemplate).Error
	return
}
// GetCPageTemplateInfoList 分页获取cPageTemplate表记录
// Author [yourname](https://github.com/yourname)
func (cPageTemplateService *CPageTemplateService)GetCPageTemplateInfoList(info manageReq.CPageTemplateSearch) (list []manage.CPageTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&manage.CPageTemplate{})
    var cPageTemplates []manage.CPageTemplate
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Id != nil {
        db = db.Where("id = ?",*info.Id)
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name = ?",*info.Name)
    }
    if info.Code != nil && *info.Code != "" {
        db = db.Where("code = ?",*info.Code)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&cPageTemplates).Error
	return  cPageTemplates, total, err
}
func (cPageTemplateService *CPageTemplateService)GetCPageTemplatePublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
