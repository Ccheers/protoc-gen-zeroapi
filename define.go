package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

type service struct {
	Name      string // Greeter
	FullName  string // helloworld.Greeter
	FilePath  string // api/helloworld/helloworld.proto
	Comment   string // 注释
	Methods   []*method
	MethodSet methodSet
}

// InterfaceName service interface name
func (s *service) InterfaceName() string {
	return s.Name + "GinServer"
}

type methodSet map[string]*method

func (x methodSet) toSortedArray() []*method {
	arr := make([]*method, 0, len(x))
	for _, m := range x {
		arr = append(arr, m)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Name > arr[j].Name
	})
	return arr
}

type method struct {
	Name    string // SayHello
	Num     int    // 一个 rpc 方法可以对应多个 http 请求
	Request string // SayHelloReq
	Reply   string // SayHelloResp
	Comment string // 注释
	// http_rule
	Path            string // gin 路由
	Method          string // HTTP Method
	Body            string // rule 定义的 Body
	MiddlewareNames []string
}

// HandlerName for gin handler name
func (m *method) HandlerName() string {
	return fmt.Sprintf("%s_%d_Handler", m.Name, m.Num)
}

// HasPathParams 是否包含路由参数
func (m *method) HasPathParams() bool {
	paths := strings.Split(m.Path, "/")
	for _, p := range paths {
		if len(p) > 0 && (p[0] == '{' && p[len(p)-1] == '}' || p[0] == ':') {
			return true
		}
	}
	return false
}

// initPathParams 转换参数路由 {xx} --> :xx
func (m *method) initPathParams() {
	paths := strings.Split(m.Path, "/")
	for i, p := range paths {
		if len(p) > 0 && (p[0] == '{' && p[len(p)-1] == '}' || p[0] == ':') {
			paths[i] = ":" + p[1:len(p)-1]
		}
	}
	m.Path = strings.Join(paths, "/")
}
