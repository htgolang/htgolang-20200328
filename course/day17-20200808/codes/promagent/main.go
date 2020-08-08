package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/imroc/req"
	uuid "github.com/satori/go.uuid"

	"promagent/config"
	"promagent/task"

	_ "promagent/task/init"
)

func initConfig(path string) *config.AgentConfig {
	var config config.AgentConfig

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PROM_AGENT")

	viper.SetConfigFile(path)
	// viper.SetConfigName("promagent")
	// viper.AddConfigPath("./etc/")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Fatal(err)
	}

	uuidPath := "promagent.uuid"
	if ctx, err := ioutil.ReadFile(uuidPath); err == nil {
		config.UUID = string(ctx)
	} else if os.IsNotExist(err) {
		config.UUID = uuid.NewV4().String()
		ioutil.WriteFile(uuidPath, []byte(config.UUID), os.ModePerm)
	} else {
		logrus.Fatal(err)
	}

	return &config
}

func initLog(verbose bool, config *config.AgentConfig) {
	logger := &lumberjack.Logger{
		Filename:   config.LogConfig.Filename,
		MaxSize:    config.LogConfig.Maxsize,
		MaxBackups: config.LogConfig.Maxbackups,
		Compress:   config.LogConfig.Compress,
	}
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(logger)
	}
}

func main() {
	var (
		verbose bool
		help    bool
		path    string
	)
	flag.BoolVar(&verbose, "verbose", false, "verbose")
	flag.BoolVar(&help, "help", false, "help")
	flag.StringVar(&path, "path", "./etc/promagent.yaml", "config path")

	flag.Usage = func() {
		fmt.Println("usage: promagent [--verbose] [--config file]")
		flag.PrintDefaults()
	}

	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}

	// 初始化配置
	config := initConfig(path)
	// 初始化日志
	initLog(verbose, config)
	if verbose {
		req.Debug = true
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT)
	errChan := make(chan error, 1)

	// 启动
	go func() {
		task.Run(config, errChan)
	}()
	logrus.WithFields(logrus.Fields{
		"pid": os.Getpid(),
	}).Info("promagent is running")

	select {
	case <-stop:
		logrus.Info("promagent stopped")
	case err := <-errChan:
		e := "someone goroutine is stopped"
		if err != nil {
			e = err.Error()
		}
		logrus.Error(e)
	}
}
