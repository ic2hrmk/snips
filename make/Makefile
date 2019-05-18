TARGET_OS = linux
TARGET_ARCH = amd64

PROJECT = project
PROJECT_BIN_DIR = bin

VERSION = 0.1.0
STAGING = dev
GITHASH = $(shell git rev-parse --short HEAD)

LDFLAGS += -X $(PROJECT)/version/common.Version=$(VERSION)
LDFLAGS += -X $(PROJECT)/version/common.Staging=$(STAGING)
LDFLAGS += -X $(PROJECT)/version/common.Githash=$(GITHASH)
LDFLAGS += -X $(PROJECT)/version/common.Created=$(shell date --rfc-3339=seconds | sed 's/ /T/')

default: build

deps:
	dep init -v .

build:
	CGO_ENABLED=0 GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) \
		go build -ldflags "$(LDFLAGS)" -a -o $(PROJECT_BIN_DIR)/$(PROJECT) main.go

.PHONY: \
	build deps

.SILENT: \
	build