package handler

import (
	"core/internal/svc"
	"fmt"
	"github.com/zeromicro/go-zero/core/stringx"
	"log"
	"net/http"
	"time"
)

func CoreSseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//l := logic.NewCoreSseLogic(r.Context(), svcCtx)
		//resp, err := l.CoreSse()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		_, ok := w.(http.Flusher)

		if !ok {
			log.Panic("server not support") //浏览器不兼容
		}
		for i := 0; i < 10; i++ {

			_, err := fmt.Fprintf(w, "data: %s\n\n", stringx.Randn(10))
			w.(http.Flusher).Flush()
			time.Sleep(2 * time.Second)
			if err != nil {
				return
			}
		}
		//_, err := fmt.Fprintf(w, "data: %s\n\n", "dsdf")
		//if err != nil {
		//	return
		//}
	}
}
