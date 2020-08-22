package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"time"
)

type AlertGroupForm struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Alerts            []*AlertForm      `json:"alerts"`
}

type AlertForm struct {
	Fingerprint  string            `json:"fingerprint"`
	Status       string            `json:"status"`
	StartsAt     *time.Time        `json:"startsAt"`
	EndsAt       *time.Time        `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
}

func (f *AlertForm) IsNew() bool {
	return f.Status == "firing"
}

func (f *AlertForm) AlertName() string {
	return f.Labels["alertname"]
}

func (f *AlertForm) LabelsString() string {
	if bytes, err := json.Marshal(f.Labels); err == nil {
		return string(bytes)
	}
	return "{}"
}

func (f *AlertForm) AnnotationsString() string {
	if bytes, err := json.Marshal(f.Annotations); err == nil {
		return string(bytes)
	}
	return "{}"
}

func (f *AlertGroupForm) AlertName() string {
	return f.GroupLabels["alertname"]
}

func (f *AlertGroupForm) GroupLabelsString() string {
	if bytes, err := json.Marshal(f.GroupLabelsString); err == nil {
		return string(bytes)
	}
	return "{}"
}

func main() {
	text := `{"receiver":"web\\.hook","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"target is down","env":"test","instance":"10.0.0.1:9999","job":"mysql"},"annotations":{"description":"节点离线","summary":"节点离线"},"startsAt":"2020-08-22T03:22:05.608589483Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://10.0.0.2:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=1","fingerprint":"591a193d73cbc0cc"}],"groupLabels":{"alertname":"target is down"},"commonLabels":{"alertname":"target is down","env":"test","instance":"10.0.0.1:9999","job":"mysql"},"commonAnnotations":{"description":"节点离线","summary":"节点离线"},"externalURL":"http://centos:9093","version":"4","groupKey":"{}:{alertname=\"target is down\"}","truncatedAlerts":0}`

	var form AlertGroupForm
	fmt.Println(json.Unmarshal([]byte(text), &form))

	funcs := map[string]interface{}{
		"dateformat": func(t *time.Time) string {
			if t == nil {
				return ""
			}
			return t.Format("2006-01-02 15:04:05")
		},
	}
	tpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles("email.html"))
	fmt.Println(tpl.ExecuteTemplate(os.Stdout, "email.html", &form))
	// smtpAddr := "smtp.qq.com"
	// smtpPort := 465 // 465或587
	// smtpUser := "782874382@qq.com"
	// smtpPassword := "wbxjliuwtlzdbehe"

	// from := "782874382@qq.com"
	// tos := []string{"yxzhanlin@163.com", "190100183@qq.com", "imsilence@outlook.com"}
	// subject := "kk找你们收邮件啦"
	// content := `邀请你们加入到<span style="color: red; font-size: 20px">KK群组</span>`

	// m := gomail.NewMessage()
	// m.SetHeader("From", from)
	// m.SetHeader("To", tos...)
	// // m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// m.SetHeader("Subject", subject)
	// m.SetBody("text/html", content)
	// // m.Attach("/home/Alex/lolcat.jpg")

	// d := gomail.NewDialer(smtpAddr, smtpPort, smtpUser, smtpPassword)

	// if err := d.DialAndSend(m); err != nil {
	// 	panic(err)
	// }
}
