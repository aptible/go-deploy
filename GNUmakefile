TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=aptible
TEST_COUNT?=1
CUR_DIR = $(shell echo "${PWD}")

default: build

build: fmtcheck
	go install

gen:
	@curl -sSfLO https://documentation-staging.s3.amazonaws.com/swagger/v1/swagger.json
	docker run --rm -v $(CURDIR):/src -v $(CURDIR):/out/api/ swaggerapi/swagger-codegen-cli generate -i src/swagger.json -l go -o /out
	@rm swagger.json

gen-mac:
	wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.28/swagger-codegen-cli-2.4.28.jar
	java -jar swagger-codegen-cli-2.4.28.jar generate -i https://documentation-staging.s3.amazonaws.com/swagger/v1/swagger.json -l go -o ./tmp
	cp tmp/api/swagger.yaml . && rm -rf tmp
	bin/swagger generate client

test: fmtcheck
	go test $(TEST) -timeout=120s -parallel=4

testapi: fmtcheck
	go test $(TEST) -timeout=120s -parallel=4 -tags=api

fmt:
	gofmt -s -w .

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

lint:
	@./bin/golangci-lint run

tools:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0
	@scripts/install-swagger.sh

.PHONY: build gen test testacc fmt fmtcheck lint tools testapi
