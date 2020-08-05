package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitializeRoutes(engine *gin.Engine,logger *logrus.Logger){
	logrus.Info("begin init router")

	//可放置全局中间件
	engine.Use(gin.RecoveryWithWriter(logger.Out))

	initDemoRouter(engine)




	logrus.Info("finished init router")
}