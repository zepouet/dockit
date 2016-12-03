CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
COMMIT=`git rev-parse --short HEAD`
APP=dockit
REPO?=treeptik/$(APP)
TAG?=latest

all: image

deps:
	glide install

build: build-dockit

build-dockit: deps
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.buildTime=`date -u +%Y.%m.%d-%Hh%Mm%S`" -o dockit-linux main.go

build-image:
	@docker build -t $(REPO):$(TAG) .
	@echo "Image created: $(REPO):$(TAG)"

image: build-dockit build-image clean

clean:
	@rm dockit-linux

.PHONY: deps build build-docker build-app build-image image clean