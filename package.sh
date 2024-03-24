#!/bin/bash

go run package.go -assets heroicons/optimized
templ generate .
templ fmt .
