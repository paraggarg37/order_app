#!/bin/bash
./setup.sh
source .env
go test -v ./... -cover -race -integration