FROM golang:1.22 AS builder

WORKDIR /Cloud-Computing

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/GymLens ./cmd/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder ./bin/GymLens .

EXPOSE 8080

CMD ["./cmd/app"]
