package utils

import (
	"html/template"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
)

func FormatEmailBody(path string, data interface{}) string {
	builder := &strings.Builder{}
	funcs := map[string]interface{}{
		"dateformat": func(t *time.Time) string {
			if t == nil {
				return ""
			}
			return t.Format("2006-01-02 15:04:05")
		},
	}
	tpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles(path))
	tpl.ExecuteTemplate(builder, filepath.Base(path), data)
	return builder.String()
}

func SendMail(tos []string, subject, content string) error {
	smtpAddr := beego.AppConfig.DefaultString("smtp::host", "smtp.qq.com")
	smtpPort := beego.AppConfig.DefaultInt("smtp::port", 465) // 465或587
	smtpUser := beego.AppConfig.DefaultString("smtp::user", "782874382@qq.com")
	smtpPassword := beego.AppConfig.DefaultString("smtp::password", "wbxjliuwtlzdbehe")

	// fmt.Println(smtpAddr, smtpPort, smtpUser, smtpPassword)
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", tos...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(smtpAddr, smtpPort, smtpUser, smtpPassword)

	return d.DialAndSend(m)
}
