FROM golang:1.18-alpine3.16 as builder
WORKDIR /demoapp
COPY . .

RUN go env -w GO111MODULE=on \
   && go env -w GOPROXY=https://goproxy.cn,direct \
   && go env -w CGO_ENABLED=0 \
   && go env \
   && go mod tidy \
   && go build -o demoapp .

FROM alpine:3.17.3
LABEL MAINTAINER="staryjie@163.com"

WORKDIR /
COPY --from=builder /demoapp/demoapp .

EXPOSE 8080

CMD ["/demoapp"]