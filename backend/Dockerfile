FROM golang:alpine AS builder

WORKDIR /rec

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GOOS linux

RUN go build -o rec-server ./cmd/main.go

FROM alpine AS runner

WORKDIR /root/

COPY --from=builder /rec/.env .

COPY --from=builder /rec/rec-server .

COPY --from=builder /rec/migrations ./migrations

EXPOSE 3000

CMD ["./rec-server"]