package handler

import (
	http1 "github.com/Ccheers/bind/http"
	v1 "github.com/Ccheers/protoc-gen-zeroapi/example/api/product/app/v1"
	logic "github.com/Ccheers/protoc-gen-zeroapi/example/internal/logic"
	svc "github.com/Ccheers/protoc-gen-zeroapi/example/internal/svc"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the Ccheers/protoc-gen-zeroapi package it is being compiled against.
func GetArticles_1_Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req v1.GetArticlesReq
		http1.TryMyBestBind(r, &req)
		if validate, ok := (interface{})(&req).(interface{ Validate() error }); ok {
			err := validate.Validate()
			if err != nil {
				svcCtx.ResponseEncodeFunc(r, w, nil, err)
				return
			}
		}
		l := logic.NewGetArticlesLogic(r.Context(), svcCtx)
		resp, err := l.GetArticles(&req)
		svcCtx.ResponseEncodeFunc(r, w, resp, err)
	}
}
