package logic

import (
	context "context"
	v1 "github.com/Ccheers/protoc-gen-zeroapi/example/api/product/app/v1"
	svc "github.com/Ccheers/protoc-gen-zeroapi/example/internal/svc"
	logx "github.com/zeromicro/go-zero/core/logx"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the Ccheers/protoc-gen-zeroapi package it is being compiled against.
type CreateArticleLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *CreateArticleLogic) CreateArticle(req *v1.Article) (*v1.Article, error) {
	panic("implement me")
}
