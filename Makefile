# Makefile

# Variables
REGISTRY = harbor.hyades.io
IMAGE_NAME = chloe
IMAGE_VERSION = 1.0

.PHONY: all build tag push clean

all: push

build:
 go build -o ./cmd/chloe/chloe ./cmd/chloe/main.go

tag:
 docker tag $(IMAGE_NAME):$(IMAGE_VERSION) $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_VERSION)

push:
 docker push $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_VERSION)

clean:
 rm ./cmd/chloe/chloe