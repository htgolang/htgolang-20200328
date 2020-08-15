package collectors

import (
	"database/sql"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type CommandCollector struct {
	*baseCollector
	desc *prometheus.Desc
}

func NewCommandCollector(db *sql.DB) *CommandCollector {
	desc := prometheus.NewDesc("mysql_command_total", "MySQL Command Total", []string{"cmd"}, nil)
	return &CommandCollector{newBaseCollector(db), desc}
}

func (c *CommandCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *CommandCollector) Collect(metrics chan<- prometheus.Metric) {
	cmds := []string{"insert", "select", "delete", "update", "replace"}

	for _, cmd := range cmds {
		metrics <- prometheus.MustNewConstMetric(
			c.desc,
			prometheus.CounterValue,
			c.status(fmt.Sprintf("com_%s", cmd)),
			cmd)
	}
}
