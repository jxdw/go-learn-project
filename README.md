# <h3>1 本仓库依赖概述</h3>
## <h4>1.1 golang sdk版本</h4>
golang 1.13.5

## <h4>1.3 DB、中间件的依赖</h4>
- 缓存技术: redis
- 存储技术: mysql
- 消息队列: kafka、rabbitmq（延迟队列）
- 注册发现: etcd
- 链路监控: jaeger(elasticsearch做存储)
- 应用监控: prometheus
- 任务调度: etcd

## <h4>1.2 工具依赖</h4>
- 开发工具: goland2019.3 (https://download.jetbrains.8686c.com/go/goland-2019.3.exe)
- 代理辅助: https://goproxy.cn
- 调试工具: postman、grpcui
- redis可视化工具:redis desktop manger
- mysql可视化工具:navicat
- etcd可视化工具:etcdkeeper
- jaeger可视化工具:jeager ui
- prometheus可视化工具:prometheus
- 接口文档工具:gin-swagger
# <h3>2 本仓库使用的框架概述</h3>
## <h4>2.1 使用的框架</h4>
- 配置框架:yaml
- MVC框架:gin
- rpc框架:grpc
- job框架:cron+etcd client
- 日志框架:file-rotatelogs
## <h4>2.2 中间件的go client</h4>
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

- kafka go client<br>
https://github.com/Shopify/sarama