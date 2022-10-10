FROM golang:1.19-alpine AS builder

ENV GO111MODULE=on
ENV GOBIN /go/bin

WORKDIR /app
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -s' -o /api-server .

FROM scratch
COPY --from=builder /api-server /api-server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/.env /.env

EXPOSE 8080

ENTRYPOINT ["/api-server"]
