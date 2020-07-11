登录成功后显示用户列表
url->用户列表页面展示
Controller => Model(获取用户数据) => View => Router


用户认证
记录用户状态? 记录在哪里?
HTTP无状态? 下一次请求

cookie session机制
状态记录 => session
    在什么时间记录(代码什么位置)?
    登录成功 记录状态 (session) sessionid
    setcookie sessionid
状态的跟踪 => (sessionid) => cookie


登录状态
    无sessionid
    有sessionid sessionid无对应session信息
    有sessionid sessionid无session登录状态

    未登录(无session登录标识)
        跳转到登录页面
    已登录 => 正常逻辑

    用户鉴权


beego
    开启: 配置 SessionOn=true/false
    存储位置: 内存，文件，数据库
            SessionProvider: file/mysql/redis
    存储的地址
            SessionProviderConfig
    cookie中存储sessionid的名字
            SessionName
    失效时间
        SessionGCMaxLifetime = 3600 s

    操作
        存储session
            controller: SetSession key value 可以是任意类型的
                        持久化的编码方式 gob 注册
        获取session
            controller: GetSession key => value interface{}
                        断言
        销毁session
            key1
            key2

            controller: DelSession(key)
            DestorySession()

1. session(登录检查)
    在任何需要登录以后才能访问的action执行之前都需要进行检查
2. 如果访问登录页面
    检查session已存在(用户已登录，就不在打开登录页面，直接跳转到首页)


1. 公共地方检查
    beego的执行过程


数据操作
存储: Table
数据 增/删/改/查

数据定义 Table => 列，类型 => 数据 => 增，删，改，查
面向对象 类 => 属性(属性名, 类型) => 实例 => 方法调用
ORM