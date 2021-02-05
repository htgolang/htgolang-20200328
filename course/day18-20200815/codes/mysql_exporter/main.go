package main

import (
	"database/sql"
	"mysql_exporter/auth"
	"mysql_exporter/collectors"
	"mysql_exporter/config"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

func initConfig() *config.ExporterConfig {
	// 配置文件解析
	return &config.ExporterConfig{
		Web: &config.WebConfig{
			Addr: ":9999",
			Auth: &config.AuthConfig{"kk", "2f64f20f7ecd620286e3cd1f6f821122"},
		},
	}
}

func main() {

	config := initConfig()

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
	http.Handle("/metrics", auth.BasicAuth(config.Web.Auth, promhttp.Handler()))

	// 启动web服务
	http.ListenAndServe(addr, nil)
}
