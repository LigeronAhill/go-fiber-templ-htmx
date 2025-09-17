# BUILDER
FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum .env ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . .
RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/main.go

# RUNNER
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/public ./public

EXPOSE 3000

CMD ["./main"]
