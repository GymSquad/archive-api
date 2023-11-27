FROM golang:1.21.4-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/main .

FROM scratch
WORKDIR /app
ENV GIN_MODE=release
COPY --from=builder /app/main /app/main
COPY docs /app/docs/
ENTRYPOINT [ "/app/main" ]

