FROM golang:1.20.3-alpine AS builder

COPY . /github.com/Murat993/chat-server/source/
WORKDIR /github.com/Murat993/chat-server/source/

RUN go mod download
RUN go build -o ./bin/server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Murat993/chat-server/source/bin/server .

CMD ["./server"]