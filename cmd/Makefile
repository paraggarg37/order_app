#! /usr/bin/make
GOBUILD=go build
SETUP_SWAG=go get -u github.com/paragtokopedia/swag/cmd/swag
GENERATE_DOC=swag init
BINARY_NAME=logistics
MAIN_FILE=cmd/app.go
RUN_INTEG_TEST=go test -v -cover -race ../integration_test
SETUP_BOILER=go get -u -t github.com/volatiletech/sqlboiler && go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql

all: boiler generate swag build run

boiler-setup:
	$(SETUP_BOILER)

generate:
	go generate

swag:
	$(SETUP_SWAG)
	cd .. && $(GENERATE_DOC) -g $(MAIN_FILE)

build:
	$(GOBUILD) -o $(BINARY_NAME)

run:
	./$(BINARY_NAME)

integtest:
	$(RUN_INTEG_TEST)
