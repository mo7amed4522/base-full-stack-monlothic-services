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
		global.GVA_LOG.Info("PostgreSQL数据库连接成功，专门用于bxapi模块")
	} else {
		global.GVA_LOG.Error("PostgreSQL数据库连接失败，bxapi模块无法使用")
	}
	return db
}
