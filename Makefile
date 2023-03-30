GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date +%FT%T%Z)
PROTO_FILES=$(shell find . -name *.proto)
KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)
PKG_VERSION=$(shell go mod graph |grep go-bamboo/pkg |head -n 1 |awk -F '@' '{print $$2}')
PKG=$(GOPATH)/pkg/mod/github.com/go-bamboo/pkg@$(PKG_VERSION)
CONTRIB_VERSION=$(shell go mod graph |grep go-bamboo/contrib |head -n 1 |awk -F '@' '{print $$2}')
CONTRIB=$(GOPATH)/pkg/mod/github.com/go-bamboo/contrib@$(CONTRIB_VERSION)
PWD := $(shell pwd)

.PHONY: errors
errors:
	protoc --proto_path=. \
		   --proto_path=$(PKG) \
		   --proto_path=$(KRATOS) \
           --proto_path=$(KRATOS)/api \
           --proto_path=$(KRATOS)/third_party \
           --proto_path=$(PWD)/../third_party \
           --go_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
           ./api/layout/ecode.proto

.PHONY: api
api:
	protoc --proto_path=. \
		   --proto_path=$(PKG) \
		   --proto_path=$(KRATOS) \
           --proto_path=$(KRATOS)/api \
           --proto_path=$(KRATOS)/third_party \
           --proto_path=$(PWD)/../third_party \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           --go-http_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
           --validate_out=lang=go,paths=source_relative:. \
           ./api/layout/types.proto ./api/layout/v1/api.proto ./api/layout/admin/admin.proto

.PHONY: proto
proto:
	protoc --proto_path=. \
		   --proto_path=$(PKG) \
		   --proto_path=$(KRATOS) \
           --proto_path=$(KRATOS)/api \
           --proto_path=$(KRATOS)/third_party \
           --proto_path=$(PWD)/../third_party \
           --go_out=paths=source_relative:. \
           --validate_out=lang=go,paths=source_relative:. \
           ./internal/conf/conf.proto

.PHONY: build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION) -X main.Branch=$(BRANCH) -X main.Revision=$(REVISION) -X main.BuildDate=$(BUILD_DATE)" -o ./bin/ ./...

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: run
run:
	./bin/layout -conf file:///configs/conf.yaml

.PHONY: install
install:
	cd cmd/layout && go install
