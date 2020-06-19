package config

const (
	StatusNew      = "未执行"
	StatusBegin    = "开始执行"
	StatusPause    = "暂停"
	StatusComplete = "完成"
)

type config struct {
	TaskFile     string
	GobTaskFile  string
	JsonTaskFile string
	CsvTaskFile  string
	RetainCsvNum int
	Header       []string
	TimeLayout   string
	StatusMap    map[string]string
	QueryMap     map[string]string
}

func New() *config {
	return &config{
		TaskFile:     "db/tasks.txt",
		GobTaskFile:  "db/tasks.gob",
		JsonTaskFile: "db/tasks.json",
		CsvTaskFile:  "db/tasks.csv",
		RetainCsvNum: 3,
		Header:       []string{"ID", "Name", "StartTime", "EndTime", "Status", "User"},
		TimeLayout:   "2006-01-02 15:04",
		StatusMap: map[string]string{
			"1": StatusNew,
			"2": StatusBegin,
			"3": StatusPause,
			"4": StatusComplete,
		},
		QueryMap: map[string]string{
			"1": "name",
			"2": "startTime",
		},
	}
}

var Config = New()
