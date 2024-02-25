include .env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 # Для работа с protobuf
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 # Для работа с protobuf
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0 # Goose - для работы с миграциями
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1 # Валидация для http grpc
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2 # GRPC GATEWAY - прокси для работы с http через grpc
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2 # GRPC GATEWAY - прокси для работы с http через grpc
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	mkdir -p pkg/swagger
	make generate-chat-api
	$(LOCAL_BIN)/statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'

generate-chat-api:
	mkdir -p pkg/chat_v1
	protoc --proto_path api/chat_v1 --proto_path vendor.protogen \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/chat_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/chat_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/chat_v1/chat.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/main.go

copy-to-server:
	scp service_linux root@89.23.116.87:

#docker-build-and-push:
#	docker buildx build --no-cache --platform linux/amd64 -t timeweb.nyssambayev.dev/server:v0.0.1 .
#	docker login -u murat -p qwe123 timeweb.nyssambayev.dev
#	docker push timeweb.nyssambayev.dev/server:v0.0.1

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

.PHONY: test
test:
	go clean -testcache
	go test ./... -covermode count -coverpkg=github.com/Murat993/chat-server/internal/service/...,github.com/Murat993/chat-server/internal/api/... -count 5

.PHONY: test-coverage
test-coverage:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out -covermode count -coverpkg=github.com/Murat993/chat-server/internal/service/...,github.com/Murat993/chat-server/internal/api/... -count 5
	grep -v 'mocks\|config' coverage.tmp.out  > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out;
	go tool cover -func=./coverage.out | grep "total";
	grep -sqFx "/coverage.out" .gitignore || echo "/coverage.out" >> .gitignore

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi