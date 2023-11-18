FROM golang:1.21.4-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/main .

FROM scratch
ENV GIN_MODE=release
COPY --from=builder /app/main /app/main
ENTRYPOINT [ "/app/main" ]

