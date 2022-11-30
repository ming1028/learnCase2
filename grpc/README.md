* protobuf

```
brew install protobuf ## 默认安装到/usr/local/bin/protoc
```

* grpc

```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest ## 代码生成插件 $GOBIN会有二进制文件
```

* grpc-gateway

```
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
```

* proto声称命令

```
protoc -I ./proto --go_out ../ --go-grpc_out ../ search.proto
```