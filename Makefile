NAME 		= embdr
MAIN		= cmd/${NAME}/main.go
SOURCES 	= $(wildcard *.go) ${MAIN}
COVERAGE	= target/coverage.txt
LINT_REPORT	= target/lint.txt
BINARY		= target/${NAME}
MODULES		= go.sum
MOD_FILE	= go.mod

export GOBIN	?= $(shell go env GOPATH)/bin
TOOLS_DIR		:= target/tools/
GOLANGCI_LINT	:= ${TOOLS_DIR}github.com/golangci/golangci-lint/cmd/golangci-lint@v1.23.8
REVIEWDOG		:= ${TOOLS_DIR}github.com/reviewdog/reviewdog/cmd/reviewdog@v0.9.17
GORELEASER		:= ${TOOLS_DIR}github.com/goreleaser/goreleaser@v0.129.0


${GOLANGCI_LINT} ${REVIEWDOG}:
	@echo Installing ${@}...
	@cd $(shell mktemp -d); GO111MODULE=on go get $(@:${TOOLS_DIR}%=%)
	@mkdir -p $(dir ${@})
	@cp ${GOBIN}/$(firstword $(subst @, ,$(notdir ${@}))) ${@}


all: ${MODULES} ${COVERAGE} ${BINARY}

build:${BINARY}
${BINARY}: ${SOURCES}
	@mkdir -p target
	go build -o ${@} ${MAIN}

modules: ${MODULES}
${MODULES}: ${SOURCES} ${MOD_FILE}
	go mod tidy
	go get ./...

test: ${COVERAGE}
${COVERAGE}: ${SOURCES}
	@mkdir -p target
	go test -race -coverprofile=${@} -covermode=atomic ./...

lint: ${LINT_REPORT}
${LINT_REPORT}: ${SOURCES} ${GOLANGCI_LINT}
	${GOLANGCI_LINT} run --out-format=line-number --enable-all --disable=wsl,gomnd,lll ./... 2>&1 | tee ${LINT_REPORT}

.PHONY: review
review: ${LINT_REPORT} ${REVIEWDOG}
	cat ${LINT_REPORT} | ${REVIEWDOG} -f=golangci-lint -reporter=github-check

.PHONY: verify-no-changes
verify-no-changes:
	git diff --name-status --exit-code

.PHONY: release
release: ${COVERAGE} ${BINARY}

clean:
	rm -rf target
