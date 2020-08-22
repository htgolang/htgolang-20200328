package filters

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/context"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cmdb_request_total",
		Help: "CMDB request total",
	})

	statusCode = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cmdb_request_status_code_total",
		Help: "CMDB Request Status Code Total",
	}, []string{"status"})

	elapsedTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "cmdb_request_elapsed_time",
		Help:    "CMDB request elapsed time",
		Buckets: prometheus.LinearBuckets(100, 300, 10),
	}, []string{"url"})
)

func init() {
	prometheus.MustRegister(requestTotal, statusCode, elapsedTime)
}

func BeforeExecute(ctx *context.Context) {
	requestTotal.Inc()
	ctx.Input.SetData("stime", time.Now())
}

func AfterExecute(ctx *context.Context) {
	statusCode.WithLabelValues(strconv.Itoa(ctx.ResponseWriter.Status)).Inc()
	stimeValue := ctx.Input.GetData("stime")
	if stimeValue != nil {
		if stime, ok := stimeValue.(time.Time); ok {
			elapsed := time.Now().Sub(stime).Milliseconds()
			elapsedTime.WithLabelValues(ctx.Input.URL()).Observe(float64(elapsed))
		}
	}
}
