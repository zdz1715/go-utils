FROM golang:1.20
LABEL authors="zdz1715/go-utils"

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct


RUN set -eux; \
    go install golang.org/x/tools/gopls@latest;
