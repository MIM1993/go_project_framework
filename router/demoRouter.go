package router

import (
	"github.com/gin-gonic/gin"
	"go_project_framework/controller"
	"net/http"
)

func initDemoRouter(engine *gin.Engine) {
	engine.Handle(http.MethodGet,"hello/world",controller.DemoController)
}
