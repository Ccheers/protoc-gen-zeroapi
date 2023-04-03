package main

import (
	"net/http"
	"regexp"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

const (
	contextPkg         = protogen.GoImportPath("context")
	netHTTPPkg         = protogen.GoImportPath("net/http")
	deprecationComment = "// Deprecated: Do not use."
)

var methodSets = make(map[string]int)

func genMethod(m *protogen.Method, g *protogen.GeneratedFile) []*method {
	var methods []*method
	// 存在 http rule 配置
	rule, ok := proto.GetExtension(m.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
	if rule != nil && ok {
		for _, bind := range rule.AdditionalBindings {
			methods = append(methods, buildHTTPRule(m, bind, g))
		}
		methods = append(methods, buildHTTPRule(m, rule, g))
		return methods
	}

	// 不存在走默认流程
	methods = append(methods, defaultMethod(m, g))
	return methods
}

// defaultMethodPath 根据函数名生成 http 路由
// 例如: GetBlogArticles ==> get: /blog/articles
// 如果方法名首个单词不是 http method 映射，那么默认返回 POST
func defaultMethod(m *protogen.Method, g *protogen.GeneratedFile) *method {
	names := strings.Split(toSnakeCase(m.GoName), "_")
	var (
		paths      []string
		httpMethod string
		path       string
		body       string
	)

	switch strings.ToUpper(names[0]) {
	case http.MethodGet, "FIND", "QUERY", "LIST", "SEARCH":
		httpMethod = http.MethodGet
	case http.MethodPost, "CREATE":
		httpMethod = http.MethodPost
		body = "*"
	case http.MethodPut, "UPDATE":
		httpMethod = http.MethodPut
		body = "*"
	case http.MethodPatch:
		httpMethod = http.MethodPatch
	case http.MethodDelete:
		httpMethod = http.MethodDelete
	default:
		httpMethod = "POST"
		paths = names
		body = "*"
	}

	if len(paths) > 0 {
		path = strings.Join(paths, "/")
	}

	if len(names) > 1 {
		path = strings.Join(names[1:], "/")
	}

	md := buildMethodDesc(m, httpMethod, path, g)
	md.Body = body
	return md
}

func buildHTTPRule(m *protogen.Method, rule *annotations.HttpRule, g *protogen.GeneratedFile) *method {
	var (
		path   string
		method string
	)
	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		path = pattern.Get
		method = "GET"
	case *annotations.HttpRule_Put:
		path = pattern.Put
		method = "PUT"
	case *annotations.HttpRule_Post:
		path = pattern.Post
		method = "POST"
	case *annotations.HttpRule_Delete:
		path = pattern.Delete
		method = "DELETE"
	case *annotations.HttpRule_Patch:
		path = pattern.Patch
		method = "PATCH"
	case *annotations.HttpRule_Custom:
		path = pattern.Custom.Path
		method = pattern.Custom.Kind
	}
	md := buildMethodDesc(m, method, path, g)
	md.Body = rule.Body
	return md
}

func buildMethodDesc(m *protogen.Method, httpMethod, path string, g *protogen.GeneratedFile) *method {
	defer func() { methodSets[m.GoName]++ }()
	md := &method{
		Name:            m.GoName,
		Num:             methodSets[m.GoName],
		Request:         g.QualifiedGoIdent(m.Input.GoIdent),
		Reply:           g.QualifiedGoIdent(m.Output.GoIdent),
		Path:            path,
		Method:          httpMethod,
		Comment:         clearComment(string(m.Comments.Leading)),
		MiddlewareNames: parseMiddleware(m.Comments.Leading.String()),
	}
	md.initPathParams()
	return md
}

var (
	matchFirstCap = regexp.MustCompile("([A-Z])([A-Z][a-z])")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func toSnakeCase(input string) string {
	output := matchFirstCap.ReplaceAllString(input, "${1}_${2}")
	output = matchAllCap.ReplaceAllString(output, "${1}_${2}")
	output = strings.ReplaceAll(output, "-", "_")
	return strings.ToLower(output)
}

func clearComment(s string) string {
	return strings.TrimSpace(strings.ReplaceAll(s, "\n", ""))
}

var middleWareMatch = regexp.MustCompile("@[A-Za-z0-9_]+")

func parseMiddleware(s string) []string {
	strs := middleWareMatch.FindAllString(s, -1)
	for i, str := range strs {
		strs[i] = strings.TrimPrefix(str, "@")
	}
	return strs
}
