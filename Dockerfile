ARG GO_VERSION=1.18
FROM golang:${GO_VERSION}-alpine as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn

ADD . .
RUN go mod download && CGO_ENABLED=0 go build -a -ldflags "-s -w"  -o gin-start-demo

FROM alpine as runner
WORKDIR /opt/app
COPY --from=builder /app/gin-start-demo /opt/app/
CMD ["./gin-start-demo"]