package register

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"promagent/config"
	"strings"
	"time"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
)

type Register struct {
	config *config.AgentConfig
}

func (r *Register) Init(config *config.AgentConfig) {
	r.config = config
}

func (r *Register) Run() {
	ticker := time.NewTicker(r.config.TaskConfig.Register.Interval)
	defer ticker.Stop()

	api := fmt.Sprintf("%s/v1/prometheus/register", strings.TrimRight(r.config.ServerConfig.Addr, "/"))
	hostname, _ := os.Hostname()
	params := req.Param{
		"uuid":     r.config.UUID,
		"addr":     r.config.Addr,
		"hostname": hostname,
	}
	headers := req.Header{
		"Authorization": fmt.Sprintf("Token %s", r.config.ServerConfig.Token),
	}
	request := req.New()
	transport, _ := request.Client().Transport.(*http.Transport)
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	for {
		// 先执行
		response, err := request.Post(api, req.BodyJSON(params), headers)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("error register")
		} else {
			body, _ := response.ToString()
			logrus.WithFields(logrus.Fields{
				"response": body,
			}).Debug("success register")
		}
		<-ticker.C
	}
}
