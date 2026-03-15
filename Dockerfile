FROM golang:1.25.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache tzdata

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]