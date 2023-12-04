### How to generate .proto

#### Pre-requisite

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

#### Command

```bash
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:proto --go-grpc_opt=paths=source_relative proto/*.proto
```
