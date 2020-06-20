0. 整理知识点
1. 任务管理
    存储到 mysql
    增/删/改/查 => 数据库

    ID
    Name
    Content
    StartTime
    Deadline
    CompleteTime
    User string

2. 用户管理[选做]
    增删改查

3. 任务与用户关联[选做]
    新建/编辑任务 使用下拉框选择用户

备注：
    增加/编辑 数据验证
        1. name 长度 1 - 32
        2. 备注 长度 < 256
        3. deadline 时间格式 > now

        4. status => 0-3
            开始（正在进行/停止/完成） => 新建
            已经完成的任务 => 不能修改任务

    status => 其他状态 startTime completetime