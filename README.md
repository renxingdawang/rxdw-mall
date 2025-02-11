# rxdw-mall


## 服务分类

auth 认证服务![image.png](images/consul/image.png)

* cart 购物车服务
* checkout 结算服务
* order 订单服务
* payment 支付服务
* product 商品服务
* user 用户服务

***商品->购物车->结算&支付->订单***

## 技术栈

Go+Kitex+Hertz+Consul+Gorm+MySQL+Redis+Jaeger+OpenTelemetry🚀️ 🚀️ 🚀️

[kItex和Hertz的doc](https://www.cloudwego.io/)

指标 链路 日志

## 数据库表设计

## WebUI

[consul](http://xxx.xxx.xx.x:8500/)

[minio](http://xxx.xx.xx.x:19001/)

[jaeger](http://xxx.xx.xx.x:16686/)

[prometheus](http://xxx.xx.x.xx:3000/)

[rabbitmq-console](http://xx.xxx.xxx.xx:15672/)

## docker compose

command

`docker compose -p <project_name> up <service_name>`
