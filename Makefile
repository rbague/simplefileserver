.PHONY: all prepare build build-current clean format tidy setup

NAME=sfs
GO_FILES=$(shell go list -f '{{range .GoFiles}}{{.}} {{end}}' ./... | xargs)

all: setup build

prepare: format clean tidy

build: prepare
	gox --osarch "!darwin/386" --output "bin/$(NAME)-{{.OS}}-{{.Arch}}"

build-current: prepare
	go build -o $(NAME) $(GO_FILES)

clean:
	@rm -f $(NAME)
	@mkdir -p bin
	@find bin -type f -exec rm {} +

format:
	gofmt -w $(GO_FILES)
	goimports -w $(GO_FILES)
	go vet ./...

tidy:
	@go mod tidy

setup:
	go mod download
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
