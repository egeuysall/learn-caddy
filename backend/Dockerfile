FROM golang:1.25.2-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o api ./main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates
RUN adduser -D -g '' appuser
USER appuser

WORKDIR /app
COPY --from=builder /app/api .

EXPOSE 8080
CMD ["./api"]
