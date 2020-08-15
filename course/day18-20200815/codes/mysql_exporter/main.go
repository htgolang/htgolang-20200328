package main

import (
	"database/sql"
	"mysql_exporter/collectors"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := ":9999"
	mysqlAddr := "localhost:3306"
	dsn := "golang:golang@2020@tcp(localhost:3306)/mysql?charset=utf8mb4&loc=PRC&parseTime=true"

	//
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logrus.Fatal(err)
	}
	// mysqlInfo := prometheus.NewGauge(prometheus.GaugeOpts{
	// 	Name:        "mysql_info",
	// 	Help:        "mysql info",
	// 	ConstLabels: prometheus.Labels{"addr": mysqlAddr},
	// })
	// mysqlInfo.Set(1)

	// 定义指标
	// 注册指标
	// 1. 时间触发 2. 业务请求触发 3. metrics请求触发
	// 可以选择的方案 1, 3
	// 2 不可以: exporter跟业务没关系
	prometheus.MustRegister(collectors.NewUpCollector(db))
	prometheus.MustRegister(collectors.NewSlowQuriesCollector(db))
	prometheus.MustRegister(collectors.NewTrafficCollector(db))
	prometheus.MustRegister(collectors.NewConnectionCollector(db))
	prometheus.MustRegister(collectors.NewCommandCollector(db))
	// prometheus.MustRegister(mysqlInfo)
	prometheus.MustRegister(prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name:        "mysql_info",
		Help:        "mysql info",
		ConstLabels: prometheus.Labels{"addr": mysqlAddr},
	}, func() float64 {
		return 1
	}))

	// 注册控制器
	http.Handle("/metrics", promhttp.Handler())
	// 启动web服务
	http.ListenAndServe(addr, nil)
}
