.DEFAULT_GOAL := build-all

export GO15VENDOREXPERIMENT=1

build-all: pb

pb:
	@protoc -I ./proto helloworld.proto --go_out=plugins=grpc:helloworld
	@ls helloworld

clean:
	@echo "清理"