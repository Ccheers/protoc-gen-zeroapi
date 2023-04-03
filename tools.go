package main

import (
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func modName() string {
	filename := "go.mod"
	modBytes, err := os.ReadFile(filename)
	failed := 0
	for err != nil {
		filename = "../" + filename
		modBytes, err = os.ReadFile(filename)
		if failed > 10 {
			panic("mod 深度太深")
		}
		failed++
	}
	return modfile.ModulePath(modBytes)
}

func modFileDir() string {
	filename := "go.mod"
	failed := 0

	_, err := os.Stat(filename)
	for err != nil {
		filename = "../" + filename
		_, err = os.Stat(filename)
		if failed > 10 {
			panic("mod 深度太深")
		}
		failed++
	}
	abs, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}
	return filepath.Dir(abs)
}

func goPackage(dir string) string {
	dir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}

	mfd := modFileDir()
	mn := modName()
	path := strings.TrimPrefix(dir, mfd)
	return mn + strings.ReplaceAll(path, "\\", "/")
}

func toUpperCamelCase(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = cases.Title(language.Und, cases.NoLower).String(s)
	return strings.ReplaceAll(s, " ", "")
}
