VERSION:=0.1.0

SOURCE:=$(wildcard *.go) pkged.go

build: go.mod pkged.go mocha

tidy: go.mod

clean:
	rm -f mocha

mocha: $(SOURCE)
	CGO_ENABLED=0 go build -trimpath -ldflags '-s -w -X main.Version=$(VERSION)'

go.mod: $(SOURCE)
	go mod tidy

pkged.go:
	pkger

.DEFAULT: build

.PHONY: build clean tidy
