# ====== 可调参数 ======
APP_PKG     := ./cmd/main		     # 你的 main.go 所在包
HC_PKG      := ./cmd/healthcheck     # 健康检查小程序包（极简）
OUT_DIR     := $(or $(OUT_DIR),./out)
VERSION     := $(or $(VERSION),dev)

# Go 编译参数（静态二进制、精简符号）
GOFLAGS     :=
LDFLAGS     := -s -w -X main.version=$(VERSION)

.PHONY: compile clean main healthcheck

default: docker_compose_run

docker_compose_run:
	docker-compose up -d --build

# 编译静态文件
compile: main healthcheck

main: openapi
	CGO_ENABLED=0 GO111MODULE=on \
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(OUT_DIR)/main $(APP_PKG)

healthcheck:
	CGO_ENABLED=0 GO111MODULE=on \
	go build -trimpath -ldflags "-s -w" -o $(OUT_DIR)/healthcheck $(HC_PKG)

openapi:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
	-generate types,chi-server \
	-package openapi \
	-o internal/openapi/api.gen.go \
	api/openapi.yaml

# run for test
run: openapi
	go run cmd/main/main.go

setEnv:
	set -a
	. ./test/env.dev
	set +a

test: setEnv run

clean:
	rm -f $(OUT_DIR)/main $(OUT_DIR)/healthcheck