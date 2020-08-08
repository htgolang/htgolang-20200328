
问题:
    1. prometheus.yaml => 修改 => prometheus reload
        => api => 开启lifecycle nginx => req
            Authorization: Basic base64(name:password)
        => systemctl reload =>
            需要有系统命令 systemctl
    2. job: 其他配置， target其他配置 根据自己业务需求添加
    3. 每间隔 每次重新写问题 每次重新加载 当内容变化后才加载
        prometheus:
            job: job1 job2 无变化不需要reload
        target:
            无变化不许写入文件

        jobs targets 排序 sort json.Marshal => []byte => 写入