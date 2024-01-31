package handler

import (
	"core/internal/logic"
	"core/internal/svc"
	"net/http"
)

func CoreSseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c := logic.NewCoreSseLogic(r.Context(), svcCtx)
		c.CoreSse(w, r)

	}
}
