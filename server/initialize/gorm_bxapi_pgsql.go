package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// GormPgSqlForBxapi 初始化 PostgreSQL 数据库专门用于bxapi模块
func GormPgSqlForBxapi() *gorm.DB {
	// 直接使用现有的 GormPgSql() 函数
	db := GormPgSql()
	if db != nil {
		global.GVA_LOG.Info("PostgreSQL database connection is successful, specifically for the bxapi module")
	} else {
		global.GVA_LOG.Error("PostgreSQL database connection failed and the bxapi module cannot be used")
	}
	return db
}
