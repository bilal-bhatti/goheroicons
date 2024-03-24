package main

import (
	"flag"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var assetsPath string
	flag.StringVar(&assetsPath, "assets", "", "path to assets directory")
	// assetsPath := flag.String("assets", "", "Path to the assets directory")
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
			// fileData, err := fs.ReadFile(content, filePath)
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			svgFileName := strings.TrimSuffix(file.Name(), ".svg")
			fmt.Println(pkgName + file.Name())

			path := strings.TrimRight(strings.ReplaceAll(pkgName, "-", "_"), "/")
			// path = "pkg" + "/" + strings.ToLower(path)
			os.MkdirAll("pkg"+"/"+"s"+strings.ToLower(path), os.ModePerm)

			fileName := "pkg" + "/" + "s" + strings.ToLower(path) + "/" + strings.ToLower(svgFileName)
			fnName := strings.ToLower(strings.Title(svgFileName))
			pkg := strings.Split(strings.TrimRight(pkgName, "/"), "/")
			fnBody := body(strings.ToLower(pkg[len(pkg)-1]), fnName, fileData)

			if err := os.WriteFile(fileName+".go", []byte(fnBody), 0666); err != nil {
				panic(err)
			}
		}
	}
}

func body(pkg, fn string, data []byte) string {
	builder := strings.Builder{}

	builder.WriteString("package " + pkg + "\n")
	builder.WriteString("\n")
	builder.WriteString("func " + strings.ReplaceAll(strings.Title(fn), "-", "") + "() string {\n")
	builder.WriteString("  return `" + strings.TrimSpace(string(data)) + "`")
	builder.WriteString("}\n")

	source, err := format.Source([]byte(builder.String()))
	if err != nil {
		panic(err)
	}

	return string(source)
}
