.DEFAULT_GOAL = bin/broker-darwin

# enable module support across all go commands.
export GO111MODULE = on
# enable consistent Go 1.12/1.13 GOPROXY behavior.
export GOPROXY = https://proxy.golang.org

# Binary

bin/broker-darwin:
	env GOOS=darwin GOARCH=amd64 go build -o broker ./main.go
.PHONY: bin/broker-darwin

bin/broker-linux:
	env GOOS=linux GOARCH=amd64 go build -o broker ./main.go
.PHONY: bin/broker-linux

# Docker

docker/build: bin/broker-linux
	docker build -t mszostok/impersonate-helm-broker:0.1.0 .
.PHONY: docker/build

docker/push: docker/build
	docker push mszostok/impersonate-helm-broker:0.1.0
.PHONY: docker/push
