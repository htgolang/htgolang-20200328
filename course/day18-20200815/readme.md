1. prometheus client_golang
    client
        a. prometheus client(java, go, php, python)
        b. http暴露 => 处理器
2. mysqld_exporter
3. web basic auth
4. 应用监控
5. alertmanager
6. 告警管理
    a. 告警接收&存储
    b. 告警查询
    c. 分页
7. 通知
    a. email
    b. 腾讯sms



开发功能 => 满足业务需求的 + 技术(go, beego)
           业务逻辑 + 技术
           1. 重复手动的工作(设计) => 自动化(开发)
           2. 需求 => 需求分析 => 设计 =>开发


监控:
1. 可用性
2. 延迟
    请求消耗时间
    操作使用时间
3. 错误次数
4. 容量
    当前请求多少/总请求多少
    当前连接数量/总的连接数量


mysql => exporter =>
    监控对象api => 获取指标信息(计算)
    sql查询 =>
                show global status

mysql 可用性
    操作失败
        select 1;
        ping
慢查询次数
