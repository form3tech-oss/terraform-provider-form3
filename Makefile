.DEFAULT_GOAL := default
GOFILES = $(shell go list ./...)
swagger_codegen_version := "v0.23.0"

export GOFLAGS=-mod=vendor

ifeq ($(shell uname),Darwin)
swagger_binary := "swagger_darwin_amd64"
else
swagger_binary := "swagger_linux_amd64"
endif

default: build test testacc

env:
	@go env

build: env vet fmtcheck
	@go install
	@mkdir -p ~/.terraform.d/plugins/
	@cp $(GOPATH)/bin/terraform-provider-form3 ~/.terraform.d/plugins/
	@echo "Build succeeded"

test: fmtcheck
	go test -v -i $(GOFILES) || exit 1
	echo $(GOFILES) | \
		xargs -t -n4 go test -count 1 -v -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 FORM3_ACC=1 go test -v -timeout 120m $(GOFILES)

testacc_debug: fmtcheck
	TF_LOG=DEBUG TF_ACC=1 FORM3_ACC=1 go test $(GOFILES) -run $(test) -v -count 1 -timeout 1m


install-swagger:
	@sudo curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/${swagger_codegen_version}/${swagger_binary} && chmod +x /usr/local/bin/swagger; \

generate-swagger-model:
	@swagger generate client -f ./swagger.yaml

vet:
	@echo "go vet ."
	@go vet $(GOFILES) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

.PHONY: vendor
vendor:
	@go mod tidy && go mod vendor && go mod verify

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./form3"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

release:
	@curl -sL http://git.io/goreleaser | bash

.PHONY: build test testacc vet fmt fmtcheck errcheck vendor-status test-compile release

