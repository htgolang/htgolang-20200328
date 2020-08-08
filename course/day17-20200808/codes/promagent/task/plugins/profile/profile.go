package profile

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"promagent/config"
	"strings"
	"time"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
)

type Profile struct {
	config *config.AgentConfig
}

func (r *Profile) Init(config *config.AgentConfig) {
	r.config = config
}

func (r *Profile) Run() {
	ticker := time.NewTicker(r.config.TaskConfig.Profile.Interval)
	defer ticker.Stop()

	api := fmt.Sprintf("%s/v1/prometheus/config", strings.TrimRight(r.config.ServerConfig.Addr, "/"))
	params := req.Param{
		"uuid": r.config.UUID,
	}
	headers := req.Header{
		"Authorization": fmt.Sprintf("Token %s", r.config.ServerConfig.Token),
	}
	request := req.New()
	transport, _ := request.Client().Transport.(*http.Transport)
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	for {
		// 先执行
		response, err := request.Get(api, params, headers)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("error profile")
		} else {
			var resp Response
			if err := response.ToJSON(&resp); err == nil {
				logrus.WithFields(logrus.Fields{
					"jobs": resp.Result,
				}).Debug("jobs")

				writePrometheus(r.config.TaskConfig.Profile.Tpl,
					r.config.TaskConfig.Profile.Output,
					resp.Result)
			} else {
				body, _ := response.ToString()
				logrus.WithFields(logrus.Fields{
					"response": body,
				}).Error("error unmarshal json profile")
			}
		}
		<-ticker.C
	}
}
