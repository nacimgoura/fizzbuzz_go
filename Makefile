ROOT := $(PWD)
GOLANG_DOCKER_IMAGE := "golang:1.21-alpine"

# Format correctly all files
# Usage:
#       make fmt
fmt:
	docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} go fmt ./...

# Check problem in source code
# Usage:
#       make vet
vet:
	docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} go vet ./...
