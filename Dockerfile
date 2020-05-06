FROM golang:1.13.1 AS builder
WORKDIR /go/src/github.com/ksmt88/grpc-web-chat
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
RUN go build -o app main.go

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/ksmt88/grpc-web-chat/app /app
EXPOSE 50051
ENTRYPOINT ["/app"]
