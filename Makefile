GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install

BINARY_PATH=./bin
BINARY_NAME=memezis
MAIN_NAME=./cmd/memezis

PROJECT_PATH=$(shell pwd)
PROTO_NAME=memezis.proto
GOBIN_PATH=$(GOPATH)/bin

PKGMAP:=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,$


all: test build

build:
	$(GOBUILD) -o $(BINARY_PATH)/$(BINARY_NAME) -v $(MAIN_NAME)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	$(BINARY_PATH)/$(BINARY_NAME)

.gen-deps:
	$(GOINSTALL) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/golang/protobuf/protoc-gen-go \
		github.com/utrack/clay/v2/cmd/protoc-gen-goclay \
		github.com/gogo/protobuf/protoc-gen-gofast


CLIENT_DIR=pkg/memezis
REL_PATH_TO_ROOT=$(shell echo $(CLIENT_DIR) | perl -F/ -lane 'print "../"x scalar(@F)')
generate: .gen-deps
	mkdir -p $(CLIENT_DIR) && cd $(CLIENT_DIR) && protoc -I/usr/local/include -I. \
      -I $(GOPATH)/src \
      -I $(REL_PATH_TO_ROOT)vendor.pb \
      -I $(REL_PATH_TO_ROOT)api/ \
      --plugin=protoc-gen-goclay=$(GOBIN_PATH)/protoc-gen-goclay \
      --plugin=protoc-gen-gofast=$(GOBIN_PATH)/protoc-gen-gofast \
      --gofast_out=$(PKGMAP)plugins=grpc:. \
      --goclay_out=$(PKGMAP)impl=true,impl_path=$(REL_PATH_TO_ROOT)internal/app/memezis,swagger=true,swagger_path=$(REL_PATH_TO_ROOT)web/swaggerui,impl_type_name_tmpl=Memezis:. \
      $(REL_PATH_TO_ROOT)api/memezis.proto

.swagger-deps:
	$(GOINSTALL) github.com/rakyll/statik

generate-swaggerui: .swagger-deps
	$(GOBIN_PATH)/statik --src $(PROJECT_PATH)/web/swaggerui --dest ./web -include=*.png,*.html,*.css,*.js,*.json
