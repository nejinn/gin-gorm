package main

import (
	"note/config"
	"note/core"
	"note/global"
	"note/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	// 初始化配置
	global.NLY_CONFIG = config.Server(initialize.InitConf())

	//初始化zap logger
	global.NLY_LOG = initialize.Zap()

	// 初始化数据库
	global.NLY_DB = initialize.InitPostgres()
	db := global.NLY_DB.DB()
	defer db.Close()
	core.RunServer()
}
