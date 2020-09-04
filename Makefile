APP_NAME = hello-world-example
PROCESS_TYPE = web

all:	test

setup:
	heroku plugins:install heroku-container-registry
	heroku container:login

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o artifact .

image: build
	docker build -t registry.heroku.com/$(APP_NAME)/$(PROCESS_TYPE) .

push: image
	docker login --email=_ --username=_ --password=$(shell heroku auth:token) registry.heroku.com
	docker push registry.heroku.com/$(APP_NAME)/$(PROCESS_TYPE)

deploy: build setup
	heroku container:push $(PROCESS_NAME) --app $(APP_NAME)

run:
	docker run registry.heroku.com/$(APP_NAME)/$(PROCESS_TYPE)

help:
	@printf "\033[32m\xE2\x9c\x93 Run the smallest and secured golang docker image based on scratch\n\033[0m"

clean: