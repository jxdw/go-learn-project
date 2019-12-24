# 1项目依赖
## 1.1 proto-gen-go工具
### 1.1.1 protoc、proto-gen-go下载<br>
下载protoc二进制版，放到golang sdk的bin目录下，可以重命名为protoc3

### 1.1.2 proto-gen-go下载<br>
  go get github.com/golang/protobuf/protoc-gen-go

### 1.1.3 生成pb.go文件<br>
protoc3 --go_out=plugins=grpc:. greetercenter/greetercenter.proto

### 1.1.4 下载grpc调试工具
官方readme说的使用办法：<br>
go get github.com/fullstorydev/grpcui<br>
go install github.com/fullstorydev/grpcui/cmd/grpcui<br>
在墙的作用下，失败了。因为墙的缘故，这里用的办法是:<br>
github.com/fullstorydev/grpcui放到mod文件,利用https://goproxy.cn下载到本地.<br>
然后IDE命令行执行go install github.com/fullstorydev/grpcui/cmd/grpcui，就会在go path目录生成grpcui二进制文件.
然后就可以用命令执行: ./grpcui.exe -plaintext grpc服务ip:grpc
