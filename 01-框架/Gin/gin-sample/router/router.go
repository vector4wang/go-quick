package router

import (
	"encoding/json"
	"fmt"
	"gin-sample/common"
	"gin-sample/controller/v1"
	v2 "gin-sample/controller/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter(r *gin.Engine) {
	r.GET("/hello", hello)
	r.GET("/extHello", hello)

	r.POST("/log/send", saveLog)

	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/say", v1.Say)
	}

	GroupV2 := r.Group("/v2", func(context *gin.Context) {
		method := context.Request.Method
		if method != http.MethodGet {
			common.RetJson("500", "method not support", nil, context)
			return
		}
	}, common.VerifySt)
	{
		GroupV2.Any("/say", v2.Say)
	}

}

func saveLog(context *gin.Context) {
	jm := make(map[string]interface{}) //注意该结构接受的内容
	context.BindJSON(&jm)
	js, err := json.Marshal(jm)
	if err != nil {
		fmt.Println("2 json err", err)
	}

	fmt.Println("req body: ", string(js))
	//log.Printf("%v", &json)
	common.RetJson2("200", "", context)
}

func hello(c *gin.Context) {
	res := make(map[string]interface{})
	res["sn"] = "hello"
	res["ts"] = time.DateTime
	common.RetJson("200", "", res, c)
}
