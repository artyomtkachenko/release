DEPS = $(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
VETARGS?=-asmdecl -atomic -bool -buildtags -copylocks -methods \
         -nilfunc -printf -rangeloops -shift -structtags -unsafeptr
all: test

deps: 
	@go get -u github.com/artyomtkachenko/release
	@go get -u github.com/gorilla/mux

test: vet 
	@go test ./...

build: test 
	@go build .

vet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@echo "--> Running go tool vet $(VETARGS) ."
	@go tool vet $(VETARGS) . ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "[LINT] Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
	fi

.PHONY: all deps test build vet
