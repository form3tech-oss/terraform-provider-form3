TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build test testacc

build: fmtcheck
	@go install
	@mkdir -p ~/.terraform.d/plugins/
	@cp $(GOPATH)/bin/terraform-provider-form3 ~/.terraform.d/plugins/
	@echo "Build succeeded"

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

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

build-release-deps:
	go get -u github.com/mitchellh/gox

build-release: build-release-deps fmtcheck vet
	gox -osarch="linux/amd64 windows/amd64 darwin/amd64" \
	-output="pkg/{{.OS}}_{{.Arch}}/terraform-provider-form3" .

release: build-release
	@test "${VERSION}" || (echo 'VERSION name required' && exit 1)
	rm -f pkg/darwin_amd64/terraform-provider-form3_${VERSION}_darwin_amd64.zip
	zip pkg/darwin_amd64/terraform-provider-form3_${VERSION}_darwin_amd64.zip pkg/darwin_amd64/terraform-provider-form3 -j
	rm -f pkg/linux_amd64/terraform-provider-form3_${VERSION}_linux_amd64.zip
	zip pkg/linux_amd64/terraform-provider-form3_${VERSION}_linux_amd64.zip pkg/linux_amd64/terraform-provider-form3 -j
	rm -f pkg/windows_amd64/terraform-provider-form3_${VERSION}_windows_amd64.zip
	zip pkg/windows_amd64/terraform-provider-form3_${VERSION}_windows_amd64.zip pkg/windows_amd64/terraform-provider-form3.exe -j

.PHONY: build test testacc vet fmt fmtcheck errcheck vendor-status test-compile release build-release build-release-deps

