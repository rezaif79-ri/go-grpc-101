# GO GRPC 101

Summary:
- Install protobuf-compiler (libprotoc)
- Run `go get -u github.com/golang/protobuf/protoc-gen-go`
- Run `go get -u google.golang.org/golang/protobuf`
- Create your proto definition in proto folders
- Run `protoc --go_out=. .\proto\user.proto` & `protoc --go-grpc_out=. ./proto/*.proto`
- 