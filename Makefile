GO_VERSION := $(shell cat ${MAKEFILE_DIR}.go-version)

## docker-compose
define docker-compose
	GO_VERSION=${GO_VERSION} \
	docker compose $1
endef

## buf generate
.PHONY: buf-gen
buf-gen:
	$(call docker-compose, run --rm --entrypoint sh buf ./scripts/buf.sh)
	goimports -w -local "github.com/appstore-notify-sample" pkg/pb/
