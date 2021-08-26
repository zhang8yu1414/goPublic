package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"zhangyudevops.com/config"
)

var (
	GVA_CONFIG config.Server
	GVA_LOG    *zap.Logger
	GVA_VP     *viper.Viper
)
