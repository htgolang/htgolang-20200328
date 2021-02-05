0. 目标公司
    查找=> 搜索引擎（百度，google），官网，招聘信息(boss)
           朋友圈/微信群/QQ群
    技术&研究方向
    目标公司相关部门人员的博客/微博

1. 准备简历
    实事求是
    工作量化，成果量化
    工作内容: 1个月的性能优化工作, 通过方式应用程序，数据库SQL, 达到的效果

2. 投简历
    招聘网站+官网
    朋友推荐
    目标公司相关部门人员的博客/微博 => 技术讨论 => 联系方式
    猎头

3. 准备面试
    熟悉 -> 做的事情
        项目 -> 细节 (2面)
        技术
            a.项目（面）
                -> 项目架构（了解）
                -> 自己负责的工作 -> 自己负责 => 系统开发 (熟悉)
                                -> 对接 => 谁来用， 怎么用 (了解)
                -> 难点, 怎么解决的, 成果
                招聘目的：解决问题
            a. 点:
                技术:
                    语法相关
                    算法 => 思路+语法（手写）
                    数据库
                    网络
                        tcp/ip
                        http
                        网络问题
                            网络设备 交换机（三层交换机），路由器
                                    DHCP NTP
                        安全问题
                            icmp, 防火墙, 堡垒机
                能力:
                    性能调优
                    设计能力
                    处理问题能力

4. 面试
    提前20分钟
        缓解情绪
        了解公司的工作氛围/文化氛围
    主动
        面试官的性格
            -> 可以被面试者引导:
                强硬的:

半年/一年
面试目的：
    1. 找工作换工作
    2. 了解市场情况
        a. 技术
            前沿技术
            套方案
        b. 薪资

golang
web开发
    html/css/javascript
    web框架

    前后端分离：
        nodejs
        前端开发技术
            js: vuejs
            css: sass => 设计/美工

        syncd

    通信
        http 提交/处理/响应
            json/ /json

        restapi/http api


安全问题：
    用户提交的任何数据都是不可信的

    数据与操作未分离

    a desc;



    a="delete table user";
    a = `" && delete table user where "" ="`

    `select * from where a="` + a + `"`;


    select * from where a="" && delete table user where "" =""

    sql注入:
        不要自己拼写sql
            数据检查 => 类型检查
                        格式检查
                        长度检查
                        范围检查
                        存在性检查
                        ...
        预处理方式 ?
    xss:
        用户提交数据当成html/js/css再浏览器执行

        text/template => 编码处理

        html/template => 数据 => html => html编码
                                 js => js编码
                                 css => css编码
                                 ....
                        1. 不要让用户提交
                        2. template.HTML => csrf_token => input