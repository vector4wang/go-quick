package main

import (
	"gin-sample/config"
	"gin-sample/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	router.InitRouter(engine)
	err := engine.Run(config.PORT)
	if err != nil {
		return
	}
}
