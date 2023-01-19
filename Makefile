.DEFAULT_GOAL := all

.PHONY: all
all: install lint build test sast

.PHONY: pre_commit
pre_commit: lint test_fast sast

.PHONY: install
install: install_go install_pnpm

.PHONY: install_go
install_go:
	@go mod vendor

.PHONY: install_pnpm
install_pnpm:
	@pnpm install

.PHONY: lint
lint: lint_pnpm lint_golang_ci

.PHONY: lint_pnpm
lint_pnpm:
	@pnpm run lint

.PHONY: lint_golang_ci
lint_golang_ci:
	@golangci-lint run

.PHONY: format
format:
	@pnpm run format && golangci-lint run --fix

.PHONY: test
test:
	@go test -race -cover -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: test_fast
test_fast:
	@CGO_ENABLED=0 go test ./...

.PHONY: build
build: build_go

.PHONY: build_go
build_go: clean_go
	@CGO_ENABLED=0 go build -mod vendor -o ./bin/tosarif ./cmd/tosarif

.PHONY: clean_go
clean_go:
	@rm -rf ./bin

.PHONY: upgrade
upgrade:
	@pnpm dlx npm-check-updates -u && pnpm upgrade && go get -u all && go mod tidy

.PHONY: sast
sast: sast_snyk sast_osv

.PHONY: sast_snyk
sast_snyk:
	snyk test --all-projects --detection-depth=1

.PHONY: sast_osv
sast_osv:
	osv-scanner ./
