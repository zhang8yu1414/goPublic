package core

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
	"zhangyudevops.com/global"
	"zhangyudevops.com/initial"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initial.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	s := &http.Server{
		Addr: address,
		Handler: Router,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		MaxHeaderBytes: 1<<20,
	}

	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}