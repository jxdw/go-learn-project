# <h3>01_development_kit使用概述</h3>
## <h4>1.1 golang版本</h4>
- golang 1.14.3
- 开发工具: goland2019.3 (https://download.jetbrains.8686c.com/go/goland-2019.3.exe)
- 代理辅助: https://goproxy.cn

# <h3>02_middleware_sdk使用概述</h3>
主要记录应用开发中常用的13个中间件的client的用法。
## <h4>2.1 依赖中间件</h4>
- 缓存技术: redis
- 存储技术: mysql
- 消息队列: kafka、rabbitmq（延迟队列）
- 注册发现: etcd
- 链路监控: jaeger(elasticsearch做存储)
- 应用监控: prometheus
- 任务调度: etcd
- 配置管理： nacos

## <h4>2.2 使用的框架</h4>
- config框架:yaml、viper
- MVC框架:gin
- rpc框架:grpc
- job框架:cron+etcd client
- log框架:file-rotatelogs
- cache框架:redis go client（https://github.com/go-redis/redis）
- db库框架:mysql go client(https://github.com/go-sql-driver/mysql) <br>
- MQ框架:rabbitmq go client(https://github.com/streadway/amqp) <br>
         kafka go client https://github.com/Shopify/sarama
- SD框架：etcd go client( https://github.com/coreos/etcd <br>
         nacos go client( http://github.com/nacos-group/nacos-sdk-go )
- trace框架:jaegar go client(https://github.com/uber/jaeger-client-go
- metrics框架:promtheus go client(https://github.com/prometheus/client_golang)
- distribute事务框架:seata go client
- jwt框架:

## <h4>2.3 工具依赖</h4>
- http调试工具: postman、
- grpc框架调试工具:grpcui、grpcurl
- redis可视化工具:redis desktop manger
- mysql可视化工具:navicat
- etcd可视化工具:etcdkeeper
- jaeger可视化工具:jeager ui
- prometheus可视化工具:prometheus
- 接口文档工具:swag

