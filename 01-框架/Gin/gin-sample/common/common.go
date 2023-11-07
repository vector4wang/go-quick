package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RetJson(code, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg, "data": data})
	c.Abort()
}

func RetJson2(code, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": "", "data": data})
	c.Abort()
}

func VerifySt(context *gin.Context) {
	st := context.Query("st")
	if st == "" {
		RetJson("500", "param no st", nil, context)
		return
	}
}
