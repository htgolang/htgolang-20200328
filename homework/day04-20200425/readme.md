1. 整理知识
2. TODO LIST
    a. 认证
        密码 =》 内置在程序中
        密码 => md5 + salt

        打开程序
            认证: 密码
            password + salt => hashed => 程序内置的hash值比较

            认证成功可以进行后续操作，如果连续输入密码超过3次失败，直接退出程序

        输入密码: fmt.Scan()

        第三方包: gopass
    b. sort
        查询 输入排序方式(name:名字 或 time: 开始时间) 升序排列

        第三方包: tablewriter

        | ID | 名字 | xxx| xxx|
        |xxx | xxxx | xx | xxx|