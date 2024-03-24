package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/beevik/etree"
)

func main() {
	var assetsPath string
	flag.StringVar(&assetsPath, "assets", "", "path to assets directory")
	flag.Parse()

	fmt.Println(assetsPath)
	if assetsPath == "" {
		fmt.Println("specify assets directory")
		fmt.Println()
		fmt.Println("usage:")
		fmt.Println("go run package.go -assets <path to icons dir>")
		fmt.Println("go run package.go -assets heroicons/optimized")
		fmt.Println()
		os.Exit(1)
	}
	embedSVGFiles(assetsPath, "")
}

func embedSVGFiles(dir, pkgName string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("error reading directory:", err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			subPkgName := pkgName + strings.Title(file.Name()) + "/"
			embedSVGFiles(filepath.Join(dir, file.Name()), subPkgName)
		} else if strings.HasSuffix(file.Name(), ".svg") {
			filePath := filepath.Join(dir, file.Name())

			fileData, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			svgFileName := strings.TrimSuffix(file.Name(), ".svg")
			fmt.Println(pkgName + file.Name())

			path := strings.TrimRight(strings.ReplaceAll(pkgName, "-", "_"), "/")
			os.MkdirAll("pkg"+"/"+"s"+strings.ToLower(path), os.ModePerm)

			fileName := "pkg" + "/" + "s" + strings.ToLower(path) + "/" + strings.ToLower(svgFileName)
			fnName := strings.ToLower(strings.Title(svgFileName))
			pkg := strings.Split(strings.TrimRight(pkgName, "/"), "/")
			fnBody := body(strings.ToLower(pkg[len(pkg)-1]), fnName, fileData)

			if err := os.WriteFile(fileName+".templ", []byte(fnBody), 0666); err != nil {
				panic(err)
			}
		}
	}
}

func body(pkg, fn string, data []byte) string {
	doc := etree.NewDocument()

	err := doc.ReadFromBytes(data)
	if err != nil {
		panic(err)
	}

	doc.Root().CreateAttr("class", "{{{styleClass}}}")

	str, err := doc.WriteToString()
	if err != nil {
		panic(err)
	}
	str = strings.ReplaceAll(str, "\"{{{styleClass}}}\"", "{ styleClass }")

	builder := strings.Builder{}
	builder.WriteString("package " + pkg + "\n")
	builder.WriteString("\n")
	builder.WriteString("templ " + strings.ReplaceAll(strings.Title(fn), "-", "") + "(styleClass string) {\n")
	builder.WriteString(strings.TrimSpace(str))
	builder.WriteString("\n}\n")

	return builder.String()
}
