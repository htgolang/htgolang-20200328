package profile

type Response struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result []*Job `json:"result"`
}

type Target struct {
	Addr string `json:"addr"`
}

type Job struct {
	Name    string    `json:"key"`
	Targets []*Target `json:"targets"`
}
