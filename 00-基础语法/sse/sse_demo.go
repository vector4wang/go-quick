package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/antage/eventsource.v1"
	"log"
	"net/http"
	"strings"
)

func main() {
	es := eventsource.New(nil, nil)
	defer es.Close()
	r := gin.Default()

	{
		r.POST("/events", esSSE)
		r.Run(":9090")
	}
}

type SSERes struct {
	Text string `json:"text"`
}

func esSSE(c *gin.Context) {
	w := c.Writer

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	_, ok := w.(http.Flusher)

	if !ok {
		log.Panic("server not support") //浏览器不兼容
	}

	jm := make(map[string]string) //注意该结构接受的内容
	c.BindJSON(&jm)
	//js, err := json.Marshal(jm)
	//js["text"]
	fmt.Println(jm)
	txt := jm["text"]

	for _, v := range strings.Split(txt, ",") {
		sr := SSERes{
			Text: v,
		}
		tmp, _ := json.Marshal(&sr)
		fmt.Fprintf(w, "data: %s\n\n", tmp)
		_, ok := w.(http.Flusher)
		if !ok {
			return
		}
	}

	fmt.Fprintf(w, "data: %s\n\n", "[DONE]")
	_, ok = w.(http.Flusher)
	if !ok {
		return
	}
}
