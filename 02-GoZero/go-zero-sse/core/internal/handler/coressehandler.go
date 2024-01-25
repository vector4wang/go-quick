package handler

import (
	"bufio"
	"core/internal/svc"
	"fmt"
	"log"
	"net/http"
	"strings"
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

		req, _ := http.NewRequest("GET", "http://localhost:8080/events", nil)
		req.Header.Set("Accept", "text/event-stream")
		client := &http.Client{}
		res, _ := client.Do(req)
		defer res.Body.Close()

		reader := bufio.NewReader(res.Body)

		for {
			line, _ := reader.ReadString('\n')
			if line == "\n" || line == "\r\n" {
				continue
			}

			fields := strings.SplitN(line, ":", 2)
			if len(fields) < 2 {
				continue
			}
			_, err := fmt.Fprintf(w, "data: %s\n", fields[1])
			w.(http.Flusher).Flush()
			time.Sleep(2 * time.Second)
			if err != nil {
				return
			}

		}

	}
}
