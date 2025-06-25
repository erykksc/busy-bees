FROM golang:1.24-alpine AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Development stage with hot reload
FROM base AS dev
RUN go install github.com/air-verse/air@latest
COPY . .
COPY .air.toml ./
ENV PORT=8080
ENV GIN_MODE=release
EXPOSE 8080
CMD ["air"]

# Production stage (minimal build)
FROM base AS builder
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /gin-app ./cmd/webserver

FROM scratch AS prod
COPY --from=builder /gin-app /
ENV PORT=8080
ENV GIN_MODE=release
EXPOSE 8080
CMD ["/gin-app"]

