FROM golang:1.16.3-alpine AS builder

WORKDIR /build
COPY . .

RUN apk add build-base

RUN go mod download
RUN go build -o gateway cmd/gateway/main.go

# Runtime image
FROM alpine:latest
WORKDIR /app
ENV PATH=/app/bin/:$PATH

COPY --from=builder /build/gateway ./bin/gateway

EXPOSE 8000

VOLUME ["/app/config.json"]

CMD ["gateway"]