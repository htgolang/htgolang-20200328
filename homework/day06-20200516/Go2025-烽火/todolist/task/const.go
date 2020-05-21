package task

const (

	
	TaskFile      = "db/tasks.txt"
	GobTaskFile   = "db/tasks.gob"
	JsonTaskFile  = "db/tasks.json"
	CsvTaskFile   = "db/tasks.csv"
	RetainCsvNum  = 3
)

const (
	statusNew      = "未执行"
	statusBegin    = "开始执行"
	statusPause    = "暂停"
	statusComplete = "完成"
)

var (
	// StatusChoices = []string{statusNew, statusBegin, statusComplete, statusPause}
	Header     = []string{"ID", "Name", "StartTime", "EndTime", "Status", "User"}
	TimeLayout = "2006/01/02 15:04:05"
	StatusMap  = map[string]string{
		"1": statusNew,
		"2": statusBegin,
		"3": statusPause,
		"4": statusComplete,
	}
)
