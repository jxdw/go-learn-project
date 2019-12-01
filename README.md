# 说明
## 1.本仓库依赖
### 1.1.1 基本依赖
- 开发工具: goland2019.3 (https://download.jetbrains.8686c.com/go/goland-2019.3.exe)
- sdk版本 : 1.13.4
- 代理辅助: https://goproxy.cn

### 1.1.2 依赖的中间件
- 缓存技术: redis
- 存储技术: mysql
- 消息队列: rabbitmq
- 注册发现: etcd
- 链路监控: jaeger+elasticsearch
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