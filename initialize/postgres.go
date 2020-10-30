package initialize

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"note/model"
	"note/utils"
)

func RegisterModel(db *gorm.DB) {
	err := db.AutoMigrate(model.User{})

	if err != nil {
		fmt.Print("迁移数据库成功")
	}

	fmt.Print("迁移数据库成功")

}

func InitPostgres() *gorm.DB {
	DbYamlConf := utils.GetDbConf()
	hostname := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		DbYamlConf.Postgres.Host,
		DbYamlConf.Postgres.Port,
		DbYamlConf.Postgres.Username,
		DbYamlConf.Postgres.DataBase,
		DbYamlConf.Postgres.Password,
		DbYamlConf.Postgres.Sslmode,
	)
	db, err := gorm.Open("postgres", hostname)
	if err != nil {
		fmt.Print("postgres启动失败")
		return nil
	} else {
		db.DB().SetMaxIdleConns(DbYamlConf.Postgres.MaxIdleConns)
		db.DB().SetMaxOpenConns(DbYamlConf.Postgres.MaxOpenConns)
		return db
	}
}
