FROM golang:1.22.5 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch as final

EXPOSE 8080

COPY --from=builder /app/main /main

CMD ["/main"]