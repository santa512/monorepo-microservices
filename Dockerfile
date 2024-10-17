#build stage
FROM golang:1.23.2 AS builder
ARG SERVICE_NAME
WORKDIR $GOPATH/src/github.com/santa512/monorepo-microservices
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o app cmd/$SERVICE_NAME/main.go

#final stage
FROM alpine:latest
WORKDIR /root/
RUN mkdir -p ./cmd/bookings
COPY --from=builder /go/src/github.com/santa512/monorepo-microservices/app .
COPY --from=builder /go/src/github.com/santa512/monorepo-microservices/config/config.yaml ./config/
CMD ["./app"]

EXPOSE 8080