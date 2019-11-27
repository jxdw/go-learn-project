项目依赖
# 1.proto-gen-go工具
## 1.1 proto-gen-go下载<br>
  go get github.com/golang/protobuf/protoc-gen-go
## 1.2 生成pb.go文件<br>
protoc --go_out=plugins=grpc:. proto/greeter.proto