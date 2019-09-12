.DEFAULT_GOAL := default

swagger_codegen_version := "v0.19.0"

ifeq (${platform},Darwin)
swagger_binary := "swagger_darwin_amd64"
else
swagger_binary := "swagger_linux_amd64"
endif

TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build test testacc

build: vet fmtcheck
	@go install
	@mkdir -p ~/.terraform.d/plugins/
	@cp $(GOPATH)/bin/terraform-provider-form3 ~/.terraform.d/plugins/
	@echo "Build succeeded"

test: fmtcheck
	go test -i $(TEST) -v || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -count 1 -v -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 FORM3_ACC=1 go test $(TEST) $(TESTARGS) -v -count 1 -timeout 120m

install-swagger:
	@sudo curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/${swagger_codegen_version}/${swagger_binary} && chmod +x /usr/local/bin/swagger; \

generate-swagger-model:
	@swagger generate client -f ./swagger.yaml

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
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

vendor-status:
	@govendor status

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

