package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	_ "golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"zhangyudevops.com/config"
)

var (
	GVA_CONFIG config.Server
	GVA_LOG    *zap.Logger
	GVA_DB     *gorm.DB
	GVA_VP                *viper.Viper
	GvaConcurrencyControl = &singleflight.Group{}
)
