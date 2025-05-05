package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/colin-404/ID-Generator/biz"
	"github.com/colin-404/logx"
	"github.com/gin-gonic/gin"
)

var sig = make(chan os.Signal, 1)

func init() {
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	logOpts := logx.Options{
		LogFile:    "./logs/log.log",
		MaxSize:    5,
		MaxAge:     3,
		MaxBackups: 3,
	}
	loger := logx.NewLoger(&logOpts)
	logx.InitLogger(loger)

}

func main() {
	//start server
	go ServerStart()
	<-sig
}

func ServerStart() {
	logx.Infof("ServerStart")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	biz.RegisterRouter(router)
	port := 6069
	logx.Infof("Listening and serving on %d", port)
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		logx.Errorf("SRV_ERROR: %v", err)
	}
}
