FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY .env . 
COPY . .

RUN apk add --no-cache make 

RUN wget -O /usr/local/bin/migrate https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz \
    && chmod +x /usr/local/bin/migrate

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .



FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY .env .      

EXPOSE 8086

ENTRYPOINT ["./main"]

CMD ["make", "migrate"]
