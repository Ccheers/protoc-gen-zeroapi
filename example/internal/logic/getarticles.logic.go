package logic

import (
	context "context"

	v1 "github.com/Ccheers/protoc-gen-zeroapi/example/api/product/app/v1"
	svc "github.com/Ccheers/protoc-gen-zeroapi/example/internal/svc"
	logx "github.com/zeromicro/go-zero/core/logx"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the Ccheers/protoc-gen-zeroapi package it is being compiled against.
type GetArticlesLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesLogic {
	return &GetArticlesLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesLogic) GetArticles(req *v1.GetArticlesReq) (*v1.GetArticlesResp, error) {
	panic("implement me222")
}
