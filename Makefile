FMT_TARGET = $(shell find . -type f -name "*.go")
LINT_TARGET = $(shell go list ./...)
TEST_TARGET = ./...

.PHONY: default
default: run

.PHONY: setup
setup:
	go install golang.org/x/lint/golint@latest
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: checkfmt
checkfmt:
	test ! -n "$(shell goimports -l $(FMT_TARGET))"

.PHONY: lint
lint:
	go vet $(LINT_TARGET)
	golint -set_exit_status $(LINT_TARGET)

.PHONY: test
test: checkfmt lint
	go test $(TEST_TARGET)

.PHONY: build
build:
	go build

.PHONY: install
install:
	go install

.PHONY: clean
clean:
	rm quizlet

.PHONY: run
run:
	@go run main.go
