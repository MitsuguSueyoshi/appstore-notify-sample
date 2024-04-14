GO_VERSION := $(shell cat ${MAKEFILE_DIR}.go-version)

## docker-compose
define docker-compose
	GO_VERSION=${GO_VERSION} \
	docker compose $1
endef

## bufのコード生成
.PHONY: buf
buf:
	$(call docker-compose, run --rm --entrypoint sh buf ./scripts/buf.sh)
	goimports -w -local "github.com/appstore-notify-sample" pkg/pb/

## イメージをビルド
.PHONY: build
image-rebuild:
	$(call docker-compose, build $(target))
