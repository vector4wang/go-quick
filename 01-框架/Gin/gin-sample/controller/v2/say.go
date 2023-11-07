package v2

import (
	"gin-sample/common"
	"github.com/gin-gonic/gin"
)

func Say(c *gin.Context) {
	msg := c.Query("msg")
	common.RetJson2("200", msg, c)
}
