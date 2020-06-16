package logging

import (
	"homework0613/tools"
	"log"
	"os"
	"time"
)

type Logger struct {
	filegroup []string
}

func NewLogger() *Logger {
	//Only leave 4 logfile in the logdir.
	filegroup := make([]string, 4)
	filegroup[0] = "tasktool_" + time.Now().Format("2006_01_02_150405") + ".log"
	return &Logger{filegroup: filegroup}
}

func (l *Logger) Logging(prefix string,logline string) {
	finfo,err := os.Stat(tools.LOGDIR + l.filegroup[0])
	if !os.IsNotExist(err){
		if finfo.Size() > tools.LOGFILESIZE {
			fdel := l.filegroup[3]
			for i := 2; i >= 0; i-- {
				l.filegroup[i+1] = l.filegroup[i]
			}
			l.filegroup[0] = "tasktool_" + time.Now().Format("2006_01_02_150405") + ".log"
			_=os.Remove(tools.LOGDIR + fdel)
		}
	}

	f, _ := os.OpenFile(tools.LOGDIR+l.filegroup[0], os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer f.Close()
	logger := log.New(f, prefix, log.LstdFlags)
	logger.Println(logline)
}
