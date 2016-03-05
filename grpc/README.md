# Naga GRPC boilerplate

This example generates GRPC method stubs using the Go protobuf library. The Naga module here implements the methods in the GRPC service. In the module's `Setup` phase the GRPC server is created and the module is registered as a service handler. In the `Start` phase the GRPC module is started.

See the Makefile for an example of how to generate Go code from `.proto` files.

## Usage

```
make prepare  # install requirements. requires homebrew.
make protos   # regenerate the pb.go files
```
