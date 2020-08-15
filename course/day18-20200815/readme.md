1. prometheus client_golang
    client
        a. prometheus client(java, go, php, python)
        b. http暴露 => 处理器
2. mysqld_exporter
3. web basic auth
    a. 服务端响应告知浏览器需要认证 => 弹出用户名密码输入框
        401
        www-authenticate: basic realm="my site"
    b. Authorization: Basic xx => 读取消息并进行验证
        Basic => xxx => Base64解码 用:分割，验证用户名:密码(明文)

        用户名:密码 => 配置文件
            hash => bcrypt/md5
            md5

    cmdb 用户登录
        form
            user/password(明文)
        db: 用户/密码(bcrypt=>hash)

4. 应用监控
    http 请求的总数 Counter
    http url 响应时间的统计 histogram/summary
    http 每个响应状态码出现次数 Counter+labels

    Filter:

5. alertmanager
6. 告警管理
    a. 告警接收&存储
        认证：
            Authorization: Token xxxxxx
            basic auth: Basic xxxx
            bearer auth: Bearer xxxxx
        webhook
            API => alertmanager => json => db

            id => labels 生成 （相同labels生成的id相同）
            1点 => instance 1.1.1.1:9999 离线
                已恢复
            10点 => instance 1.1.1.1:9999 离线 id

            a => a
            a,b => a,b

        通知：groupkey 分组为单位通知
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



告警 => Prometheus 关联(Node)

promagent => prometheus.yaml labels => uuid => xx


json
form  => controller => parseForm/unmarshal => object => server.insert/


                      unmarshal => ALertForm


groupby
alertname

a => a
    labels =>

b => a,b
    a => labels =>
    b =>

分页
    pageSize => 5
    pageNum
    查询条件


    Offset Limit
    pageNum => 1,2,3
    (pageNum) - 1  * pageSize
    limit PageSize

    1 = 0 limit 5
    2 => 5 limit 5

querytable.SetCond(cond).Offset().limit()
querytable.SetCond(conf).Count()

Page
    datas
    分页相关数据 页面 URL参数