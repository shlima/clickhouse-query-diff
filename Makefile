REGISTRY=ghcr.io/shlima/clickhouse-query-diff
GIT_COMMIT=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION=0.0.1
BUILD_INFO_PACKAGE="github.com/shlima/clickhouse-query-diff/internal/pkg/buildinfo"
LDFLAGS="-s -w -X ${BUILD_INFO_PACKAGE}.BuildArgTime=$(BUILD_TIME) -X ${BUILD_INFO_PACKAGE}.BuildArgGitCommit=$(GIT_COMMIT) -X ${BUILD_INFO_PACKAGE}.BuildArgVersion=$(VERSION)"

docker-build:
	@echo "Build a Docker container for the current platform (locally)"
	$(MAKE) buildx ARG=--load

docker-push:
	@echo "Build Docker images for all available platforms and push them to the registry <$(REGISTRY)>"
	$(MAKE) buildx ARG="--push --platform linux/arm64,linux/amd64"

buildx: ARG=
buildx:
	@echo ""
	@echo ""
	@echo "building for the current platform..."
	@echo "👉 docker buildx create --use"
	@echo ""
	@echo ""

	docker buildx build \
	$(ARG) \
	--no-cache \
	--build-arg LDFLAGS=$(LDFLAGS) \
	--build-arg GIT_COMMIT=$(GIT_COMMIT) \
	--build-arg VERSION=$(VERSION) \
	-t $(REGISTRY):$(VERSION) \
	-t $(REGISTRY):latest \
	-f Dockerfile .

	@echo "👌 OK"
	@echo "docker run --rm -i ${REGISTRY}"
	@echo "docker run --rm -i --entrypoint server ${REGISTRY}"
