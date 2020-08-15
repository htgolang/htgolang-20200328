package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type ConnectionCollector struct {
	*baseCollector
	maxConnectionsDesc   *prometheus.Desc
	threadsConnectedDesc *prometheus.Desc
}

func NewConnectionCollector(db *sql.DB) *ConnectionCollector {
	maxConnectionsDesc := prometheus.NewDesc("mysql_global_variables_max_connections", "mysql globa variables max connections", nil, nil)
	threadsConnectedDesc := prometheus.NewDesc("mysql_global_status_threads_connected", "mysql globa status threads connected", nil, nil)
	return &ConnectionCollector{newBaseCollector(db), maxConnectionsDesc, threadsConnectedDesc}
}

func (c *ConnectionCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.maxConnectionsDesc
	descs <- c.threadsConnectedDesc
}

func (c *ConnectionCollector) Collect(metrics chan<- prometheus.Metric) {
	metrics <- prometheus.MustNewConstMetric(c.maxConnectionsDesc, prometheus.GaugeValue, c.variables("max_connections"))
	metrics <- prometheus.MustNewConstMetric(c.threadsConnectedDesc, prometheus.GaugeValue, c.status("threads_connected"))
}
