#!/bin/bash

go install github.com/a-h/templ/cmd/templ@latest
go get -u github.com/a-h/templ
go run package.go -assets heroicons/optimized
templ generate .
templ fmt .
