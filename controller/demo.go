package controller

import "github.com/gin-gonic/gin"

func DemoController(c *gin.Context){
	c.Writer.Write([]byte("hello world"))
}
