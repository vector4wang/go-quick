package main

import (
	"gin-sample/config"
	"gin-sample/middleware"
	"gin-sample/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(middleware.LoggerToFile())
	router.InitRouter(engine)
	err := engine.Run(config.PORT)
	if err != nil {
		return
	}
}
