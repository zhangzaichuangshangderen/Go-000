学习笔记

第1课，Go架构实践-微服务（微服务概览与治理）
微服务概览

微服务设计

gRoc&服务发现

多集群&多租户


SOA和微服务什么关系？
SOA：Service-Oriented Architecture，面向服务的结构
微服务可以认为是SOA的一种实践
小即是美
单一职责
尽可能早的创建原型
可移植性比效率更重要


微服务定义
1. 原子服务  关注单一业务
2. 独立进程
3. 隔离部署
4. 去中心化服务治理
    数据去中心化
    治理去中心化
    技术去中心化

缺点：
基础设施的建设，复杂度高


组件服务化
kit：一个微服务的基础库
service：业务代码+kit依赖+第三方依赖组成的业务微服务
rpc+message queue:轻量级通讯


Design For Failure
    隔离
    超时控制
    负载保护
    限流
    降级
    重试
    负载均衡



API Gateway
API Gateway 负责安全认证、限流熔断等
BFF Backend for front  面向前端的后端（这不是我现在做的事情吗= =）


linux 环境高级编程
