GOOS=linux
GOARCH=386

.PHONY: build cui

GIT_COMMIT := $(shell git rev-list -1 HEAD)
DATE			?= $(shell date +%FT%T%z)
#uses https://github.com/markchalloner/git-semver
VERSION			?= $(shell git semver get)
HASH			?= $(shell git rev-parse HEAD)
DBNAME ?= pickleit.db
PRODUCTION ?= FALSE

LDFLAGS			+= -X "main.dbName=${DBNAME}"
LDFLAGS			+= -X "main.production=${PRODUCTION}"
LDFLAGS			+= -X "main.buildVersion=${VERSION}"
LDFLAGS			+= -X "main.buildHash=${HASH}"
LDFLAGS			+= -X "main.buildDate=${DATE}" 
# LDFLAGS			+= -s -w
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.gitCommit=$(GIT_COMMIT)" .

cui:
	go build -ldflags "$(LDFLAGS)"