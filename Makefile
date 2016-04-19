export PATH := /usr/local/go/bin:$(PATH)
TEST?=$$(go list ./... | grep -v /vendor/)
export TMPDIR := $(PWD)/tmp
DEPS = $(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
VETARGS=-asmdecl -atomic -bool -buildtags -copylocks -methods \
									-nilfunc -printf -rangeloops -shift -structtags -unsafeptr

VERSION=$(shell [ -n "${GO_PIPELINE_COUNTER}" ] && echo "${GO_PIPELINE_COUNTER}" || echo "dev" )

all: test

fmt:
		mkdir -p tmp
		go fmt .

vet: fmt
		go tool vet -v ${VETARGS} $$(ls -d */ | grep -v vendor) 

test: vet
		go test -v $(TEST) -timeout=30s -parallel=4

cover: test
		go test -coverprofile=cover.out  .
		go tool cover -html=cover.out

bench: test
		go test . -bench=. -cpuprofile=cpu.pprof
		go tool pprof --pdf release.test cpu.pprof >cpu_stats.pdf
		open cpu_stats.pdf

build: test
		go install -ldflags "-X main.version=${VERSION}" .
		rm -rf tmp

.PHONY: all deps vet test cover bench build fmt

