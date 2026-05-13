# -------------------------
# Base
# -------------------------
FROM golang:1.25.6-alpine AS base

WORKDIR /app
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

# -------------------------
# Development stage
# -------------------------
FROM base AS dev

RUN go install github.com/air-verse/air@latest

COPY . .

EXPOSE 8080

CMD ["air"]

# -------------------------
# Build stage
# -------------------------
FROM base AS build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -buildvcs=false -ldflags="-s -w" \
    -o server ./cmd/api

# -------------------------
# Production stage
# -------------------------
FROM scratch AS prod

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/server /server

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
  CMD ["/server", "--healthcheck"]

ENTRYPOINT ["/server"]