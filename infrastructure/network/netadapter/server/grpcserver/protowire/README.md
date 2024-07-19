protowire
=========

1. Download and place in your PATH:
   https://github.com/protocolbuffers/protobuf-go/releases/download/v1.34.2/protoc-gen-go.v1.34.2.linux.amd64.tar.gz
2. `go install google.golang.org/protobuf/protoc-gen-go`
3. `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`
4. In the protowire directory: `go generate .`

Documentation
-------------

To generate `rpc.md`:
1. `go install -u github.com/kaspanet/protoc-gen-doc/cmd/protoc-gen-doc`
2. In the protowire directory: `protoc --doc_out=. --doc_opt=markdown,rpc.md rpc.proto`
