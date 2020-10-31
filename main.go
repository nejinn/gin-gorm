package main

import (
	"note/config"
	"note/global"
	"note/initialize"
)

func main() {
	// 初始化配置
	global.NLY_CONFIG = config.Server(initialize.InitConf())
	//初始化zap logger
	global.NLY_LOG = initialize.Zap()
	// 初始化数据库
	global.NLY_DB = initialize.InitPostgres()
	// 注册model
	initialize.RegisterModel(global.NLY_DB)
	db := global.NLY_DB.DB()
	defer db.Close()
}
