all: build

APP?=paraTest

build:
	dep ensure
	CGO_ENABLED=0 go build -a -installsuffix cgo \
		-o ./bin/${APP} ./


local: build

run: local
	./bin/${APP}