protos:
	PROTO_FILES=$$(find . -name "*.proto"); \
	  PROTO_DIR=$$(find . -name "*.proto" -exec dirname {} \; | sort -u | sed -e 's/^/-I/'); \
	  protoc $$PROTO_DIR $$PROTO_FILES \
	    --go_out=plugins=setter+grpc:./proto

prepare:
	brew install grpc/grpc/google-protobuf
	go get -u github.com/golang/protobuf/protoc-gen-go
