GO_VERSION := $(shell cat ${MAKEFILE_DIR}.go-version)

## docker-compose
define docker-compose
	GO_VERSION=${GO_VERSION} \
	docker compose $1
endef

## protoのコード生成
.PHONY: bufgen
bufgen:
	$(call docker-compose, run --rm --entrypoint sh buf ./scripts/buf.sh)
	goimports -w -local "github.com/appstore-notify-sample" pkg/pb/

## イメージをビルド
.PHONY: build
build:
	$(call docker-compose, build $(target))

## サーバーを起動
.PHONY: api-run
api-run:
	go run cmd/api/main.go
