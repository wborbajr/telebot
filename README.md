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

## References

```link
https://sohlich.github.io/post/go_makefile/
```
