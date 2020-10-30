package global

import (
	"gin-vue-admin/config"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
	GVA_LOG    *zap.Logger
)
