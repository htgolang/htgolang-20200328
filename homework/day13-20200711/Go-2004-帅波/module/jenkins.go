package module

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/strive-after/go-cmdb/gojenkins"
)

type JenkinsInfo struct {
	user string
	password string
	host string
	token string
}

func NewJenkins() JenkinsInfo {
	return JenkinsInfo{
		user:     beego.AppConfig.String("jenkins_user"),
		password: beego.AppConfig.String("jenkins_password"),
		host:     beego.AppConfig.String("jenkins_host"),
		token:    beego.AppConfig.String("jenkins_token"),
	}
}


var (
	jenkinsclient *gojenkins.Jenkins
	request *gojenkins.Requester
)

func init() {
	jenkins := NewJenkins()
	jenkinsclient, err = gojenkins.CreateJenkins(nil,jenkins.host,jenkins.user,jenkins.token).Init()
	if err != nil {
		fmt.Println(err)
		return
	}
}


