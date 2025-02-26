NAME := gncompany

LINT_VERSION := v1.63.4
BUILD_DIR := ${NAME}

PKG := `go list -mod=mod -f {{.Dir}} ./...`
MAIN := cmd/${NAME}/main.go

build:
	@go build -mod=mod -o ${BUILD_DIR} ./cmd/${NAME}/main.go

fmt:
	@go run golang.org/x/tools/cmd/goimports@latest -local ${NAME} -l -w $(PKG)

lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@$(LINT_VERSION) run -c .golangci.yml

mod-download:
	@go mod download all

mod-tidy:
	@go mod tidy

mod: mod-tidy mod-download

pre-commit: fmt lint
