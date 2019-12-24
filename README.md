# 说明
## 1.本仓库依赖
### 1.1.1 基本依赖
- 开发工具: goland2019.3 (https://download.jetbrains.8686c.com/go/goland-2019.3.exe)
- sdk版本 : 1.13.4
- 代理辅助: https://goproxy.cn
- 调试工具: postman、grpcui

### 1.1.2 依赖的DB、中间件
- 缓存技术: redis
- 存储技术: mysql
- 消息队列: rabbitmq
- 注册发现: etcd
- 链路监控: jaeger(elasticsearch做存储)
- 应用监控: prometheus
- 任务调度: etcd+web ui

## 1.2.使用的框架和中间件 go client
### 1.2.1 使用的框架
- 配置框架 : yaml
- MVC框架 : gin
- rpc框架 : grpc
- job框架 ：cron+etcd client

### 1.2.2 中间件的go client
- redis go client<br>
 https://github.com/go-redis/redis
 
- mysql go client<br> 
https://github.com/go-sql-driver/mysql

- rabbitmq go client<br> 
https://github.com/streadway/amqp

- etcd go client<br>
 https://github.com/coreos/etcd

- jaegar go client<br> 
https://github.com/uber/jaeger-client-go

- promtheus go client<br>
https://github.com/prometheus/client_golang

# 2.各个project的参考资料
## 2.1 golang01_basic_api工程
- 1.unkown的go语言入门:https://github.com/unknwon/the-way-to-go_ZH_CN
- 2.go http编程:https://cizixs.com/2016/08/17/golang-http-server-side/
- 3.golang版设计模式: https://github.com/senghoo/golang-design-pattern