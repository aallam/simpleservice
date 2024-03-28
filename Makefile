# Makefile for building and pushing Docker images with multi-architecture support

IMAGE_NAME := aallam/simpleservice
VERSION := 0.2.0
BUILDER_NAME := builder
PORT=8080

# Default target
all: publish

# Create a new builder instance
create-builder:
	docker buildx create --name $(BUILDER_NAME) --use

# Start up the builder instance
bootstrap-builder:
	docker buildx inspect --bootstrap

# Build the Docker image for local platform and tag it with latest and version
build:
	docker build -t simpleservice --build-arg=VERSION=$(VERSION) -f build/Dockerfile .

# Build the Docker image for multiple platforms and tag it with latest and version
publish:
	docker buildx build --platform linux/amd64,linux/arm64 \
		--build-arg VERSION=$(VERSION) \
		-t $(IMAGE_NAME):latest \
		-t $(IMAGE_NAME):$(VERSION) \
		--push .

docker-run:
	@docker run --rm -p $(PORT):$(PORT) -e PORT=$(PORT) simpleservice

docker-boot: build docker-run

# Clean up builder instance
clean:
	docker buildx rm $(BUILDER_NAME)

.PHONY: all create-builder bootstrap-builder publish clean
