package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go_project_framework/conf"
	"go_project_framework/global"
	"go_project_framework/internal"
	"go_project_framework/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	Version   = "not set"
	CommitID  = "not set"
	BuildTime = "not set"
	v         = flag.Bool("v", false, "display version")
)

func main() {
	if *v {
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("CommitID: %s\n", CommitID)
		fmt.Printf("BuildTime: %s\n", BuildTime)
		return
	}

	//加载配置文件,并且接收命令行参数
	err := conf.InitConfig("conf/goproject.yaml")
	if err != nil {
		panic(fmt.Sprintf("init config err:%v", err))
	}

	fmt.Println(global.Global)

	//log level, 2:ERROR 3:WARN 4:INFO 5:DEBUG
	err = internal.LogSet(4, true, "")
	if err != nil {
		panic(fmt.Sprintf("set log err:%v", err))
	}

	//链接数据库
	logrus.StandardLogger().Info("link database")

	//gin router
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	//加载路由
	router.InitializeRoutes(engine, logrus.StandardLogger())
	logrus.Info("Initialize Router OK!")

	//启动http server
	srv := &http.Server{
		Addr:           "127.0.0.1:8900",
		Handler:        engine,
		ReadTimeout:    25 * time.Second,
		WriteTimeout:   25 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		logrus.Info("server start at ", srv.Addr, " OK!")
	}()

	// 监听退出信号,如果
	go internal.ListeningExitSignal()

	//阻止主groutine退出
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	//优雅的关闭所有groutine
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	os.Exit(0)
}
