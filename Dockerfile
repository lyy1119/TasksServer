ARG VERSION=dev
########## build ##########
FROM golang:1.24-bookworm AS build

WORKDIR /src
# 安装make
RUN apt-get update
RUN apt-get install make
# 构建缓存（加速 go mod / go build）
ENV CGO_ENABLED=0 GO111MODULE=on \
    GOCACHE=/root/.cache/go-build \
    GOMODCACHE=/go/pkg/mod

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# 版本号注入
ARG VERSION
ENV VERSION=${VERSION}

# 再拷源码与 Makefile
COPY . .

# 输出到根目录下的 /out
ENV OUT_DIR=/out

# 编译
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    make -j compile

RUN chmod +x /out/.

########## runtime ##########
FROM scratch AS runtime
ARG VERSION
ENV VERSION=${VERSION}
WORKDIR /app
# 防root权限
USER 65532:65532

# 复制静态编译文件
COPY --from=build /out/main main
COPY --from=build /out/healthcheck healthcheck

CMD ["./main"]