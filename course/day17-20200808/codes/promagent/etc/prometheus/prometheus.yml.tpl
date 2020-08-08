global:
  scrape_interval: 15s
  evaluation_interval: 15s
  query_log_file: "prometheus.log"

alerting:
  alertmanagers:
    - static_configs:
      - targets:
        - "localhost:9093"

rule_files:
  - "rules/*.yml"

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']