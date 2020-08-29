package profile

import (
	"fmt"
	"os"
	"path/filepath"
	"promagent/utils"

	"gopkg.in/yaml.v2"
)

type FileSdConfig struct {
	Files []string `yaml:"files"`
}

func NewFileSdConfig(files ...string) *FileSdConfig {
	return &FileSdConfig{files}
}

type ScrapConfig struct {
	JobName       string          `yaml:"job_name"`
	StaticConfigs interface{}     `yaml:"static_configs,omitempty"`
	FileSdConfigs []*FileSdConfig `yaml:"file_sd_configs"`
}

func NewScrapConfig(job string) *ScrapConfig {
	paths := []string{
		fmt.Sprintf("sd/file/%s/*.yaml", job),
		fmt.Sprintf("sd/file/%s/*.json", job),
	}
	return &ScrapConfig{
		JobName:       job,
		FileSdConfigs: []*FileSdConfig{NewFileSdConfig(paths...)},
	}
}

type PrometheusConfig struct {
	Global        interface{}    `yaml:"global"`
	Alerting      interface{}    `yaml:"alerting"`
	RuleFiles     []string       `yaml:"rule_files"`
	ScrapeConfigs []*ScrapConfig `yaml:"scrape_configs"`
}

func writePrometheus(tpl, path string, jobs []*Job) error {
	utils.MkPdir(path)
	file, err := os.Open(tpl)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var config PrometheusConfig
	if err := decoder.Decode(&config); err != nil {
		return err
	}

	for _, job := range jobs {
		scrape := NewScrapConfig(job.Name)
		config.ScrapeConfigs = append(config.ScrapeConfigs, scrape)
		path := filepath.Join(filepath.Dir(path), fmt.Sprintf("sd/file/%s/%s.yaml", job.Name, job.Name))
		writeTarget(path, job.Targets)
	}

	output, err := os.Create(path)
	if err != nil {
		return err
	}
	defer output.Close()

	encoder := yaml.NewEncoder(output)
	if err := encoder.Encode(config); err != nil {
		return err
	}
	// reload prometheus
	return nil
}
