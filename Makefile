#!/bin/bash

export PKGS=$(shell go list ./...)
export GO111MODULE=on

vet:
	@echo "---- VERIFY ----"
	@go vet ./... ${PKGS}

build:
	@echo "---- BUILD -----"
	@go mod vendor
	@go build -o ./cmd/toolkit/toolkit ./cmd/toolkit/
	@GOOS=linux go build -o ./cmd/toolkit/nix_toolkit ./cmd/toolkit/

vault: vet build
	@echo "---- GENERATE VAULT SECRET -----"
	@./cmd/toolkit/toolkit  -vault -output=conf.json -secret=vaultkv/data/yourdata
	@./cmd/toolkit/toolkit  -vault  -env -secret=vaultkv/data/envdata
	@./cmd/toolkit/toolkit  -vault  -raw -secret=vaultkv/data/rawtextfile

git: vet build
	@echo "---- VERIFY PULL REQUEST -----"
	@./cmd/toolkit/toolkit  -debug -approval -repo=ujunglangit-id/some-repo -id=23