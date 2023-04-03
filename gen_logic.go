package main

import (
	"fmt"
	"os"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	logxPackage = protogen.GoImportPath("github.com/zeromicro/go-zero/core/logx")
)

func genZeroLogic(outDir string, gen *protogen.Plugin, file *protogen.File) {
	if len(file.Services) == 0 {
		return
	}

	rootPackage := protogen.GoImportPath(goPackage(outDir))
	for _, service := range file.Services {
		genZeroLogicFiles(outDir, rootPackage, file, gen, service)
	}
	return
}

func genZeroLogicFiles(outDir string, rootPackage protogen.GoImportPath, file *protogen.File, gen *protogen.Plugin, s *protogen.Service) {
	uniMap := make(map[string]struct{})
	for _, method := range s.Methods {
		if _, ok := uniMap[method.GoName]; ok {
			continue
		}
		uniMap[method.GoName] = struct{}{}
		generateZeroLogicFile(outDir, rootPackage, method, gen, s)
	}
}

func generateZeroLogicFile(outDir string, rootPackage protogen.GoImportPath, method *protogen.Method, gen *protogen.Plugin, s *protogen.Service) {
	logicPackage := rootPackage + "/internal/logic"
	svcPackage := rootPackage + "/internal/svc"
	filename := outDir + "/internal/logic/" + strings.ToLower(method.GoName) + ".logic.go"
	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		return
	}
	g := gen.NewGeneratedFile(filename, logicPackage)
	g.P("package logic")
	g.P()
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the Ccheers/protoc-gen-zeroapi package it is being compiled against.")

	if s.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
	g.P(fmt.Sprintf("type %sLogic struct {", method.GoName))
	g.P("logger ", logxPackage.Ident("Logger"))
	g.P("ctx ", contextPkg.Ident("Context"))
	g.P("svcCtx *", svcPackage.Ident("ServiceContext"))
	g.P("}")

	g.P("func New", method.GoName, "Logic(ctx ", contextPkg.Ident("Context"), ", svcCtx *", svcPackage.Ident("ServiceContext"), ") *", method.GoName, "Logic {")
	g.P("return &", method.GoName, "Logic{")
	g.P("logger: ", logxPackage.Ident("WithContext"), "(ctx),")
	g.P("ctx: ctx,")
	g.P("svcCtx: svcCtx,")
	g.P("}")
	g.P("}")

	g.P("func (l *", method.GoName, "Logic) ", method.GoName, "(req *", g.QualifiedGoIdent(method.Input.GoIdent), ") (*", g.QualifiedGoIdent(method.Output.GoIdent), ", error) {")
	g.P("panic(\"implement me\")")
	g.P("}")
}
