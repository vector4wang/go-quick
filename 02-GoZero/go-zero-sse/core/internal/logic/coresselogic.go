package logic

import (
	"bufio"
	"bytes"
	"context"
	"core/internal/svc"
	"core/internal/types"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// 在对接其他sse接口并往外提供是，需要判断截止字符串，主动进行return，否则会死循环
// sse 消息格式体为："data: %s\n\n" 其他系统过来的sse已经包含了\n,需注意
type CoreSseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreSseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreSseLogic {
	return &CoreSseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreSseLogic) CoreSse(w http.ResponseWriter, r *http.Request) {
	WriteStreamHeaders(w)
	_, ok := w.(http.Flusher)

	if !ok {
		//log.Panic("server not support") //浏览器不兼容
		fmt.Println("server not support")
		return
	}

	// 参数校验
	var reqBody types.SSERequest
	if err := httpx.Parse(r, &reqBody); err != nil {
		_, _ = fmt.Fprintf(w, "data: %s\n\n", err.Error())
		w.(http.Flusher).Flush()
		return
	}

	body := make(map[string]string)
	body["text"] = reqBody.Text
	jsonData, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "http://localhost:9090/events", bytes.NewBuffer(jsonData))
	req.Header.Set("Accept", "text/event-stream")
	client := &http.Client{}
	res, _ := client.Do(req)
	defer res.Body.Close()

	reader := bufio.NewReader(res.Body)

	for {
		line, _ := reader.ReadString('\n')
		fmt.Printf("line: %s", line)
		if line == "\n" || line == "\r\n" {
			continue
		}

		fields := strings.SplitN(line, ":", 2)
		if len(fields) < 2 {
			continue
		}

		_, err := fmt.Fprintf(w, "data: %s\n", fields[1])
		if err != nil {
			return
		}
		w.(http.Flusher).Flush()
		if strings.HasPrefix(strings.TrimSpace(fields[1]), "[DONE]") {
			return
		}

	}
	return

}

func WriteStreamHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")
}
