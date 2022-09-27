FROM golang:1.19.1-alpine as builder

WORKDIR /app
COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN go build -o ./proxy ./main.go

from alpine:3.16

WORKDIR /app
COPY --from=builder /app/proxy ./proxy
COPY --from=builder /app/public ./public

CMD ["./proxy"]
