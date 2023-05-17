package handler

import (
	http "net/http"

	http1 "github.com/Ccheers/bind/http"
	v1 "github.com/Ccheers/protoc-gen-zeroapi/example/api/product/app/v1"
	logic "github.com/Ccheers/protoc-gen-zeroapi/example/internal/logic"
	svc "github.com/Ccheers/protoc-gen-zeroapi/example/internal/svc"
)

/// asdasd
// This is a compile-time assertion to ensure that this generated file
// is compatible with the Ccheers/protoc-gen-zeroapi package it is being compiled against.
func CreateArticle_0_Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req v1.Article
		http1.TryMyBestBind(r, &req)
		if validate, ok := (interface{})(&req).(interface{ Validate() error }); ok {
			err := validate.Validate()
			if err != nil {
				svcCtx.ResponseEncodeFunc(r, w, nil, err)
				return
			}
		}
		l := logic.NewCreateArticleLogic(r.Context(), svcCtx)
		resp, err := l.CreateArticle(&req)
		svcCtx.ResponseEncodeFunc(r, w, resp, err)
	}
}
