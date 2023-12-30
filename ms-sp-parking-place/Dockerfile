# Estágio de construção
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY .env . 
COPY . .

RUN apk add --no-cache make 

RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz -O /tmp/migrate.tar.gz \
    && tar xvzf /tmp/migrate.tar.gz -C /usr/local/bin/ \
    && rm /tmp/migrate.tar.gz

# Construir o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# Estágio final
FROM alpine:latest
RUN apk --no-cache add ca-certificates make

WORKDIR /app

# Copiar o executável e definir permissões
COPY --from=builder /app/main .
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

# Copiar os arquivos necessários, incluindo as migrações
COPY .env .    
COPY Makefile .
COPY sql/migrations ./sql/migrations

# Garantir que o caminho do migrate está no $PATH
ENV PATH="/usr/local/bin:${PATH}"

RUN wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-alpine-linux-amd64-v0.6.1.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-v0.6.1.tar.gz \
    && rm dockerize-alpine-linux-amd64-v0.6.1.tar.gz

EXPOSE 8086

ENTRYPOINT ["./main"]

CMD ["make", "migrate"]


