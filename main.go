package main

import (
	"database/sql"
	"zhangyudevops.com/core"
	"zhangyudevops.com/global"
	"zhangyudevops.com/initial"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initial.Gorm()

	if global.GVA_DB != nil {
		initial.MysqlTables(global.GVA_DB)
		db, _ := global.GVA_DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
	}
	core.RunWindowsServer()
}