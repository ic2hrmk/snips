SYSARCH = arm
MCKARCH = amd64

PROJECT = project

VERSION = 0.1.0
STAGING = devel
GITHASH = $(shell git rev-parse --short HEAD)

LDFLAGS += -X $(PROJECT)/version/common.Version=$(VERSION)
LDFLAGS += -X $(PROJECT)/version/common.Staging=$(STAGING)
LDFLAGS += -X $(PROJECT)/version/common.Githash=$(GITHASH)
LDFLAGS += -X $(PROJECT)/version/common.Created=$(shell date --rfc-3339=seconds | sed 's/ /T/')

default: build

deps:
	dep init

build:
	GOOS=linux CGO_ENABLED=0 GOARCH=$(SYSARCH) \
		go build -ldflags "$(LDFLAGS)" -a -o $(DIST)/bin/$(PROJECT) main.go

.PHONY: all deps build