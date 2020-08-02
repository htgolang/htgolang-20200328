yaml

key: value

{
    key: value
}

key string
value int/string/boolean/slice/map

注释 #开头
缩进 2个空格/4个空格

name: kk
age: 30
sex: true
scores: [1, 2, 3]
hobbies:
  - 足球
  - 乒乓球
mscores: {"math" : 1, "chinese" : 3}
mscores2:
  math:
    - 1
    - 2
  chinese: 2



scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']
  - job_name: "node"
    basic_auth:
        username: ""
        password: ""
    static_configs:
      - targets:
        - "localhost:9100"


{
    scrape_configs: [
        {
            job_name : "prometheus".
            basic_auth : {
                username: "",
                password : ""
            }
            static_configs: [
                {
                    targets: ["locahost:9090"]
                }
            ]
        }
    ]
}