package main

import (
	"fmt"
	"os"
)

// 生成骨架
//api
//├── bookstore.go                   // main入口定义
//├── etc
//│   └── bookstore-api.yaml         // 配置文件
//└── internal
//├── config
//│   └── config.go              // 定义配置
//├── handler
//│   ├── addhandler.go          // 实现addHandler
//│   ├── checkhandler.go        // 实现checkHandler
//│   └── routes.go              // 定义路由处理
//├── logic
//│   ├── addlogic.go            // 实现AddLogic
//│   └── checklogic.go          // 实现CheckLogic
//├── svc
//│   └── servicecontext.go      // 定义ServiceContext

func genZeroLayout(outDir string) {
	err := os.MkdirAll(outDir+"/internal/handler", os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(outDir+"/internal/logic", os.ModePerm)
	if err != nil {
		panic(err)
	}
	buildConfigFile(outDir)
	buildConfig(outDir)
	buildSvc(outDir)
	buildMain(outDir)
}

func buildSvc(outDir string) {
	err := os.MkdirAll(outDir+"/internal/svc", os.ModePerm)
	if err != nil {
		panic(err)
	}
	filename := outDir + "/internal/svc/servicecontext.go"
	_, err = os.Stat(filename)
	if !os.IsNotExist(err) {
		return
	}

	fw, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	const content = `package svc

import (
	"%s/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
`
	_, err = fw.Write([]byte(fmt.Sprintf(content, goPackage(outDir))))
	if err != nil {
		panic(err)
	}
}

func buildMain(outDir string) {
	filename := outDir + "/main.go"
	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		return
	}
	fw, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	const content = `package main

import (
	"flag"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"%s/internal/config"
	"%s/internal/handler"
	"%s/internal/svc"

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

	fmt.Printf("Starting server at %%s:%%d...\n", c.Host, c.Port)
	server.Start()
}
`
	packageName := goPackage(outDir)
	_, err = fw.Write([]byte(fmt.Sprintf(
		content,
		packageName,
		packageName,
		packageName,
	)))
	if err != nil {
		panic(err)
	}
}

func buildConfig(outDir string) {
	filename := outDir + "/internal/config/config.go"
	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		return
	}
	err = os.MkdirAll(outDir+"/internal/config", os.ModePerm)
	if err != nil {
		panic(err)
	}
	fw, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	const content = `package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
}
`
	_, err = fw.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}

func buildConfigFile(outDir string) {
	filename := outDir + "/etc/config.local.yaml"
	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		return
	}
	err = os.MkdirAll(outDir+"/etc", os.ModePerm)
	if err != nil {
		panic(err)
	}
	fw, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	const content = `Name: api
Host: 0.0.0.0
Port: 8888
Mode: dev
`
	_, err = fw.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}
