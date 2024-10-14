FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

EXPOSE 8000

EXPOSE 50051

EXPOSE 8080

COPY . .

CMD ["go", "run", "cmd/ordersystem/main.go", "cmd/ordersystem/wire_gen.go"]