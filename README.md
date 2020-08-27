# telebot

Telegram BOT using GOLang

## Start Server

```bash
source .env && go run main.go
```

## Generating Security Certs

```bash
openssl genrsa -out server.key
openssl req -new -sha256 -key server.key -out server.csr
openssl x509 -req -days 3650 -in server.csr -out server.crt -signkey server.key
```

## Docker

### Images Size

```bash
golang  alpine 402MB
golang  latest 839MB
```

### Build

```bash
docker-compose up -d --build --no-cache --pull
```

### Run

```bash
docker-compose up -d
docker-compose down
```

## References

```link
https://sohlich.github.io/post/go_makefile/
https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3
https://medium.com/tech-at-wildlife-studios/write-backend-systems-50aae8db849e
https://ewanvalentine.io/
```
