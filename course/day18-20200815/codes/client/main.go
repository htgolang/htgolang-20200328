package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// counter
	requestTotal := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "request_total",
		Help: "request total",
	})

	codeStatus := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "status_code_total",
		Help: "status code total",
	}, []string{"code"})
	// guage

	// 固定label
	cpu := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "cpu",
		Help:        "cpu total",
		ConstLabels: prometheus.Labels{"a": "xxx"},
	})

	// 非固定label
	disk := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "disk",
		Help: "disk total",
	}, []string{"mount"})

	// historgram
	requestTime := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "request_time",
		Help:    "request time",
		Buckets: prometheus.LinearBuckets(0, 3, 5),
	}, []string{"url"})
	// summary
	requestSummary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "request_time_summary",
		Help:       "request time summary",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"url"})

	// metrics_name{label=label_value} metrics_value

	// 有lable
	// label/label_value value固定
	// label/label_value value不固定
	// 无lable => label都是空 => 固定

	cpu.Set(2)

	disk.WithLabelValues("c:").Set(100)
	disk.WithLabelValues("e:").Set(200)

	// 注册指标信息
	prometheus.MustRegister(cpu)
	prometheus.MustRegister(disk)
	prometheus.MustRegister(requestTotal)
	prometheus.MustRegister(codeStatus)
	prometheus.MustRegister(requestTime)
	prometheus.MustRegister(requestSummary)

	requestTotal.Add(10)
	codeStatus.WithLabelValues("200").Add(10)
	codeStatus.WithLabelValues("500").Add(3)
	codeStatus.WithLabelValues("500").Add(1)

	requestTime.WithLabelValues("/aaaa").Observe(6)
	requestSummary.WithLabelValues("/aaaa").Observe(6)
	requestTime.WithLabelValues("/aaaa").Observe(2)
	requestSummary.WithLabelValues("/aaaa").Observe(2)

	// 值的修改
	// 修改的时间 => 触发
	// 时间触发
	// 磁盘使用，cpu使用，内存使用
	go func() {
		for range time.Tick(time.Second) {
			disk.WithLabelValues("c:").Set(float64(rand.Int()))
		}
	}()
	// 事件触发
	// 业务请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestTotal.Inc()
		codeStatus.WithLabelValues(strconv.Itoa(rand.Intn(5) * 100)).Add(1)
		requestTime.WithLabelValues(r.URL.Path).Observe(float64(rand.Intn(20)))
		requestSummary.WithLabelValues(r.URL.Path).Observe(float64(rand.Intn(20)))
		fmt.Fprint(w, "hi")
	})

	// 在metics接口访问的时候
	call := prometheus.NewCounterFunc(prometheus.CounterOpts{
		Name: "xxx",
		Help: "xxxxx",
	}, func() float64 {
		fmt.Println("call")
		return rand.Float64()
	})

	// prometheus.NewGaugeFunc()
	// Collector
	prometheus.MustRegister(call)

	// 暴露
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":9999", nil)

}
