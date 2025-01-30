
# === CONFIG 8< ================================================================
#
#

OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
NO_COLOR=\033[0m
 
# === CONSTANTS =============================================================

BUILD_DIR ?= build
APP_BINARY_SRC=cmd/main.go

app_name=synmail-api


.PHONY: build
build:
	@printf "$(OK_COLOR)==> Building synmail-api binaries $(NO_COLOR)\n"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/${app_name} ${APP_BINARY_SRC}

.PHONY: test
test:
	@printf "$(OK_COLOR)==> Running Test $(NO_COLOR)\n"
	mkdir -p "${BUILD_DIR}"
	go test -v -race ./...

.PHONY: lint
lint:
	@printf "$(OK_COLOR)==> Running Linter following https://github.com/golang/go/wiki/CodeReviewComments $(NO_COLOR)\n"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.1
	golangci-lint run --timeout=2m ./...

.PHONY: fmt
fmt:
	go install mvdan.cc/gofumpt@latest
	gofumpt -w .
