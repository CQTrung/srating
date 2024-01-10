FROM golang:1.21-alpine as builder

WORKDIR /builder

ENV GO111MODULE=on CGO_ENABLED=0

COPY . .

RUN go mod download && go build -o /builder/main

# Runner image
FROM gcr.io/distroless/static:latest

WORKDIR /app

COPY --from=builder /builder/main .
COPY .env .

CMD ["./main","server"]
