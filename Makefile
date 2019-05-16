.DEFAULT_GOAL := default

swagger_codegen_version := "0.15.0"

ifeq (${platform},Darwin)
swagger_binary := "swagger_darwin_amd64"
else
swagger_binary := "swagger_linux_amd64"
endif

TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build testacc

build: fmtcheck vet
	go install

testacc: fmtcheck
	FORM3_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

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


.PHONY: build testacc vet fmt fmtcheck errcheck vendor-status
