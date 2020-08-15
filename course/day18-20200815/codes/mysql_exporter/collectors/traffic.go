package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type TrafficCollector struct {
	*baseCollector
	desc *prometheus.Desc
}

func NewTrafficCollector(db *sql.DB) *TrafficCollector {
	desc := prometheus.NewDesc("mysql_traffic_total", "MySQL Traffic Total", []string{"direction"}, nil)
	return &TrafficCollector{newBaseCollector(db), desc}
}

func (c *TrafficCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *TrafficCollector) Collect(metrics chan<- prometheus.Metric) {
	in := c.status("Bytes_received")
	out := c.status("Bytes_send")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, in, "in")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, out, "out")
}
