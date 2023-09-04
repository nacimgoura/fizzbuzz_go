# multi stage build
FROM golang:1.21-alpine as builder
# env variable
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
# install upx for reduce size
RUN apk update --no-cache && apk add upx
# Add a work directory
WORKDIR /build
# Copy files for built binaries
COPY . .
# download dependencies
RUN go mod download
# Build api
RUN go build -ldflags="-s -w" -a -o bin/fizzbuzz_go main.go
# reduce size
RUN upx --lzma /build/bin/fizzbuzz_go

# create scratch image
FROM scratch as app
# Copy certificate for call https (like infura)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy built binary from builder
COPY --from=builder /build/bin/fizzbuzz_go ./usr/local/bin/fizzbuzz_go
# Run app
ENTRYPOINT ["fizzbuzz_go"]