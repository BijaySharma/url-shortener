# Build Stage
FROM golang:1.26.1-alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o write-service cmd/main.go


# Runtime Stage
FROM alpine:3.19
RUN apk add --no-cache ca-certificates && \
    adduser -D -u 10001 appuser
WORKDIR /app
COPY --from=build /app/write-service .
USER appuser
EXPOSE 8080
CMD ["./write-service"]