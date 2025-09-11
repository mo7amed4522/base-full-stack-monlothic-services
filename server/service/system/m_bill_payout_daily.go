
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type MBillPayoutDailyService struct {}
// CreateMBillPayoutDaily 创建mBillPayoutDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayoutDailyService *MBillPayoutDailyService) CreateMBillPayoutDaily(mBillPayoutDaily *system.MBillPayoutDaily) (err error) {
	err = global.GVA_DB.Create(mBillPayoutDaily).Error
	return err
}

// DeleteMBillPayoutDaily 删除mBillPayoutDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayoutDailyService *MBillPayoutDailyService)DeleteMBillPayoutDaily(id string) (err error) {
	err = global.GVA_DB.Delete(&system.MBillPayoutDaily{},"id = ?",id).Error
	return err
}

// DeleteMBillPayoutDailyByIds 批量删除mBillPayoutDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayoutDailyService *MBillPayoutDailyService)DeleteMBillPayoutDailyByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.MBillPayoutDaily{},"id in ?",ids).Error
	return err
}

// UpdateMBillPayoutDaily 更新mBillPayoutDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayoutDailyService *MBillPayoutDailyService)UpdateMBillPayoutDaily(mBillPayoutDaily system.MBillPayoutDaily) (err error) {
	err = global.GVA_DB.Model(&system.MBillPayoutDaily{}).Where("id = ?",mBillPayoutDaily.Id).Updates(&mBillPayoutDaily).Error
	return err
}

// GetMBillPayoutDaily 根据id获取mBillPayoutDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayoutDailyService *MBillPayoutDailyService)GetMBillPayoutDaily(id string) (mBillPayoutDaily system.MBillPayoutDaily, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&mBillPayoutDaily).Error
	return
}
// GetMBillPayoutDailyInfoList 分页获取mBillPayoutDaily表记录
// Author [yourname](https://github.com/yourname)
func (mBillPayoutDailyService *MBillPayoutDailyService)GetMBillPayoutDailyInfoList(info systemReq.MBillPayoutDailySearch) (list []system.MBillPayoutDaily, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&system.MBillPayoutDaily{})
    var mBillPayoutDailys []system.MBillPayoutDaily
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&mBillPayoutDailys).Error
	return  mBillPayoutDailys, total, err
}
func (mBillPayoutDailyService *MBillPayoutDailyService)GetMBillPayoutDailyPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
