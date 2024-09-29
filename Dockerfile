FROM golang:1.22.7-alpine3.20

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o shop cmd/server/main.go

EXPOSE 8080

CMD ["./shop"]