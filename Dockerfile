FROM golang:1.22 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/main ./cmd/main.go

FROM alpine:latest
EXPOSE 8080
COPY --from=builder /app/bin/main .

CMD ["/main"]