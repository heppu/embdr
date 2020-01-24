NAME 		= embdr
MAIN		= cmd/${NAME}/main.go
SOURCES 	= $(wildcard *.go) ${MAIN}
COVERAGE	= target/cover.txt
BINARY		= target/${NAME}
MODULES		= go.sum
MOD_FILE	= go.mod

all: ${COVERAGE} ${BINARY}

test: ${COVERAGE}
${COVERAGE}: ${SOURCES}
	@mkdir -p target
	go test -race -cover -coverprofile=${@} ./...

build:${BINARY}
${BINARY}: ${SOURCES}
	@mkdir -p target
	go build -o ${@} ${MAIN}

modules: ${MODULES}
${MODULES}: ${SOURCES} ${MOD_FILE}
	go mod tidy
	go get ./...

.PHONY: verify-no-changes
verify-no-changes:
	git diff --name-status --exit-code

.PHONY: release
release: ${COVERAGE} ${BINARY}
