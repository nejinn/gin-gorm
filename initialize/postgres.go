package initialize

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/zap"
	"note/global"
	"note/model"
	//"note/utils"
)

// 迁移数据库
func RegisterModel(db *gorm.DB) {
	global.NLY_LOG.Info("开始迁移数据库")
	db.AutoMigrate(&model.User{})
	global.NLY_LOG.Info("迁移表 User 成功")
	db.AutoMigrate(&model.Nav{})
	global.NLY_LOG.Info("迁移表 Nav 成功")
	db.AutoMigrate(&model.VisitApi{})
	global.NLY_LOG.Info("迁移表 VisitApi 成功")
	db.AutoMigrate(&model.VisitPage{})
	global.NLY_LOG.Info("迁移表 VisitPage 成功")
	db.AutoMigrate(&model.VisitUser{})
	global.NLY_LOG.Info("迁移表 VisitUser 成功")
	global.NLY_LOG.Info("迁移数据库成功")
}

// 构造 数据库初始化函数
func InitPostgres() *gorm.DB {
	global.NLY_LOG.Info("开始初始化数据库")
	hostname := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		global.NLY_CONFIG.Postgres.Host,
		global.NLY_CONFIG.Postgres.Port,
		global.NLY_CONFIG.Postgres.Username,
		global.NLY_CONFIG.Postgres.DataBase,
		global.NLY_CONFIG.Postgres.Password,
		global.NLY_CONFIG.Postgres.Sslmode,
	)
	db, err := gorm.Open("postgres", hostname)
	if err != nil {
		global.NLY_LOG.Info("数据库连接失败", zap.String("hostname", hostname))
		return nil
	} else {
		global.NLY_LOG.Info("postgres启动成功: ", zap.String("hostname", hostname))
		db.DB().SetMaxIdleConns(global.NLY_CONFIG.Postgres.MaxIdleConns)
		db.DB().SetMaxOpenConns(global.NLY_CONFIG.Postgres.MaxOpenConns)
		RegisterModel(db)
		return db
	}
}
