package svc

import (
	"net/http"

	httpx "github.com/zeromicro/go-zero/rest/httpx"

	"github.com/Ccheers/protoc-gen-zeroapi/example/internal/config"
)

type ServiceContext struct {
	Config config.Config
	ResponseEncodeFunc func(*http.Request, http.ResponseWriter,interface{}, error)
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ResponseEncodeFunc: func(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
			if err != nil {
				httpx.Error(w, err)
				return
			}
			httpx.WriteJson(w, http.StatusOK, resp)
		},
	}
}
