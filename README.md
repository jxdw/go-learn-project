# 1.项目依赖
## 1.1 项目基本依赖
- 开发工具: goland 2019.3候选版本 (https://download.jetbrains.8686c.com/go/goland-193.5233.69.exe)
- sdk版本 : 1.13.4
- 代理辅助: https://goproxy.cn

## 1.2 依赖的中间件
- 缓存技术: redis
- 存储技术: mysql
- 消息队列: rabbitmq
- 注册发现: etcd
- 链路监控: jaeger+elasticsearch
- 应用监控: prometheus
- 任务调度: etcd+web ui

# 2.使用的框架和中间件的go client
## 2.1 使用的框架
- 配置框架 : yaml
- MVC框架 : gin
- rpc框架 : grpc
- job框架 ：cron+etcd client

## 2.2 中间件的go client
- 缓存go client: redis client
- 存储go client: mysql client
- 消息go client: rabbitmq client
- 注册go client: etcd client
- 链路go client: jaeger client
- 监控go client: prometheus client
