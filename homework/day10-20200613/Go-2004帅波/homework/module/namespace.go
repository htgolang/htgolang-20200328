package module

import (
"time"
)

type NameSpace struct {
	CreateTime time.Time
	RunTime time.Time
	Deploy  string
	Sts  string
	ConfigMap  string
	Job  string
	CronJob string
	Pod string
}
