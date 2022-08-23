## RPC remote procedure call 远程过程调用

    解决分布式系统通信，rpc协议属于应用层协议

* 连接管理
* 健康监测
* 负载均衡
* 优雅启停机
* 异常重试
* 业务分组
* 熔断限流

## CACHE MISS

* 先写数据库，后delete缓存
* 先写数据库，后set缓存
    - 优化：使用channel异步set缓存，或者pipeline 将大批数据同步写缓存（延迟换吞吐率）