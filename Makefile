DEPS = $(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
VETARGS?=-asmdecl -atomic -bool -buildtags -copylocks -methods \
         -nilfunc -printf -rangeloops -shift -structtags -unsafeptr
all: test

test: deps
	@sh -c "'$(PWD)/scripts/test.sh'"
	@$(MAKE) vet

deps:
	@echo "--> Installing build dependencies"
	# @DEP_ARGS="-d -v" sh -c "'$(PWD)/scripts/deps.sh'"


updatedeps: deps
	@echo "--> Updating build dependencies"
	# @DEP_ARGS="-d -f -u" sh -c "'$(PWD)/scripts/deps.sh'"

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

	@git grep -n `echo "log"".Print"` ; if [ $$? -eq 0 ]; then \
		echo "[LINT] Found "log"".Printf" calls. These should use Nomad's logger instead."; \
	fi


.PHONY: all deps test vet
