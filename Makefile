service.proto:
	protoc --proto_path=proto --proto_path=third-party --go_out=plugins=grpc:proto service.proto

.PHONY: all
all: service.proto

