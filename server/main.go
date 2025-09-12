package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger API接口文档
// @version                     v2.7.7
// @description                 A full-stack development platform using gin+vue for rapid development
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() // Initialize Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // Initialize the zap log library
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // init MySQL for landingpage project
	//global.GVA_PG_DB = initialize.GormPgSqlForBxapi() // init PostgreSQL for bxapi
	utils.Init()
	initialize.Redis()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // Initialization table
		// Close the database connection before the program ends
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
