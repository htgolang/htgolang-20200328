FROM ubuntu:18.04

LABEL user="kk"
LABEL email="imsilence@outlook.com"

# 生命
EXPOSE 80

# web 编译结果
ADD web /opt/web/

ENTRYPOINT [ "/opt/web/web" ]