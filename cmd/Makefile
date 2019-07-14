#! /usr/bin/make
SHELL := /bin/bash

GOBUILD=go build
BINARY_NAME=logistics
SETUP_BOILER=go get -u -t github.com/volatiletech/sqlboiler && go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql

all: build run

boiler:
	$(SETUP_BOILER)

generate:
	go generate

build:
	$(GOBUILD) -o $(BINARY_NAME)

run:
	./$(BINARY_NAME)

docker-run:
	docker build -t logistics ../ && docker run -p 8080:8080 logistics

docker-dev-start:
	docker-compose -f ../docker-compose.yml up --build

