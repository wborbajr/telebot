![Docker](https://github.com/wborbajr/telebot/blob/master/img/telebot.jpg)

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
docker-compose build --no-cache --pull
or
docker-compose up -d --build --no-cache --pull
```

### Run

```bash
docker-compose up  && docker-compose down
or
docker-compose up -d
docker-compose down
```

### Cleanup <none> images

```bash
docker rmi $(docker images -f "dangling=true" -q) --force
```

## References

```link
https://sohlich.github.io/post/go_makefile/
https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3
https://medium.com/tech-at-wildlife-studios/write-backend-systems-50aae8db849e
https://ewanvalentine.io/
https://sosedoff.com/2019/02/12/go-github-actions.html
https://blog.alexellis.io/introducing-functions-as-a-service/

https://madeddu.xyz/posts/goa-docker-multistage/
https://codeburst.io/using-rabbitmq-for-microservices-communication-on-docker-a43840401819

```
