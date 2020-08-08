用户管理
1. 登录
2. 用户增、删、改、查

用户信息
id parimary key
staff_id
name
nickname
password
gender
tel
addr
email
department
status

created_at
updated_at
deleted_at

部门管理


Data[nav]=home
Data[nav]=user


name like ? or addr like ? and deleted_at is null
(name like ?) or (addr like ? and deleted_at is null)

(name like ? or addr like ?) and deleted_at is null


api => json

controller => json
beego:
    CopyRequestBody=true
    c.Input.RequestBody json => json.Unmarshal => Form `json`

beego: xsrf检查 => 关闭xsrf检查
    认证: 登录=> set-cookie: sessionid, cookie: sessionid
            Token: 固定的随机字符串 V
                    JWT token (认证)
                    id key => id params(timestamp) 签名token

            header: Authrozation: Token xxx
                                  Bearer xxxx



v1
    PrometheusController
v2
    PRometheusController

v1/prometheus
v2/prometheus