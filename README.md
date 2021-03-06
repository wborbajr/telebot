![CI](https://github.com/wborbajr/telebot/workflows/CI/badge.svg)

![License](https://img.shields.io/github/license/wborbajr/telebot?style=plastic)

![TeleBot](./docs/images/telebot.jpg)

# telebot

Telegram BOT using GOLang

## Start Server

```bash
source .env && go run main.go
```

## Generating Security Certs

If you need, you may generate a self signed certficate, as this requires HTTPS / TLS. The above example tells Telegram that this is your certificate and that it should be trusted, even though it is not properly signed.

```bash
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 3560 -subj "//O=Org\CN=Test" -nodes
```

```bash
openssl genrsa -out server.key
openssl req -new -sha256 -key server.key -out server.csr
openssl x509 -req -days 3650 -in server.csr -out server.crt -signkey server.key
```

## Docker

### Portainer.io

```bash
$ docker volume create portainer_data

$ docker run -d -p 8000:8000 -p 9000:9000 --name=portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce
```

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

### Logs

```bash
docker-compose logs -f
```

### Cleanup <none> images

```bash
docker rmi $(docker images -f "dangling=true" -q) --force
```

## References

https://sohlich.github.io/post/go_makefile/

https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3

https://medium.com/tech-at-wildlife-studios/write-backend-systems-50aae8db849e

https://ewanvalentine.io/

https://sosedoff.com/2019/02/12/go-github-actions.html

https://blog.alexellis.io/introducing-functions-as-a-service/

https://madeddu.xyz/posts/goa-docker-multistage/

https://codeburst.io/using-rabbitmq-for-microservices-communication-on-docker-a43840401819

https://www.linode.com/docs/applications/containers/deploying-microservices-with-docker/

https://www.linode.com/docs/applications/containers/how-to-create-a-docker-swarm-manager-and-nodes-on-linode/

https://www.linode.com/docs/applications/containers/docker-container-communication/

https://www.freecodecamp.org/news/build-a-blockchain-in-golang-from-scratch/

https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc

https://www.cloudreach.com/en/resources/blog/cts-build-golang-dockerfiles/

### Reference Docker

https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

### Best Practice

https://www.reddit.com/r/golang/comments/fah0v8/best_practice_to_dockerize_golang_113/

https://blog.berkgokden.com/continuous-integration-practices-with-docker-part-2-3573da77a9d8

https://dev.to/nodepractices/docker-best-practices-with-node-js-4ln4

https://github.com/goldbergyoni/nodebestpractices

## iTerm2 shortcut

Expand split
```
Cmd + Shift + Enter
```

Split Horizontal
```
Cmd + Shift + D 
```

Split Vertical
```
Cmd + Shift + D 
```

Switch between  splits
```
Cmd + [ or Cmd + ]
```

Save layout
```
Cmd + Shift + S
```

Restore layout
```
Cmd + Shift + R
```
