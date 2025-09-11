package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type MBillPayDailyService struct{}

// CreateMBillPayDaily 创建mBillPayDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayDailyService *MBillPayDailyService) CreateMBillPayDaily(mBillPayDaily *system.MBillPayDaily) (err error) {
	err = global.GVA_DB.Create(mBillPayDaily).Error
	return err
}

// DeleteMBillPayDaily 删除mBillPayDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayDailyService *MBillPayDailyService) DeleteMBillPayDaily(id string) (err error) {
	err = global.GVA_DB.Delete(&system.MBillPayDaily{}, "id = ?", id).Error
	return err
}

// DeleteMBillPayDailyByIds 批量删除mBillPayDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayDailyService *MBillPayDailyService) DeleteMBillPayDailyByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.MBillPayDaily{}, "id in ?", ids).Error
	return err
}

// UpdateMBillPayDaily 更新mBillPayDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayDailyService *MBillPayDailyService) UpdateMBillPayDaily(mBillPayDaily system.MBillPayDaily) (err error) {
	err = global.GVA_DB.Model(&system.MBillPayDaily{}).Where("id = ?", mBillPayDaily.Id).Updates(&mBillPayDaily).Error
	return err
}

// GetMBillPayDaily 根据id获取mBillPayDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayDailyService *MBillPayDailyService) GetMBillPayDaily(id string) (mBillPayDaily system.MBillPayDaily, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&mBillPayDaily).Error
	return
}

// GetMBillPayDailyInfoList 分页获取mBillPayDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayDailyService *MBillPayDailyService) GetMBillPayDailyInfoList(info systemReq.MBillPayDailySearch) (list []system.MBillPayDaily, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.MBillPayDaily{})
	var mBillPayDailys []system.MBillPayDaily
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	db = db.Order("id desc")
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mBillPayDailys).Error
	return mBillPayDailys, total, err
}
func (mBillPayDailyService *MBillPayDailyService) GetMBillPayDailyPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
