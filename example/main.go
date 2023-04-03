package main

import (
	"flag"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/Ccheers/protoc-gen-zeroapi/example/internal/config"
	"github.com/Ccheers/protoc-gen-zeroapi/example/internal/handler"
	"github.com/Ccheers/protoc-gen-zeroapi/example/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/config.local.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
