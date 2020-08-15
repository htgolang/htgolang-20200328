package response

type JSONResposne struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func NewJsonResponse(code int, msg string, result interface{}) *JSONResposne {
	return &JSONResposne{code, msg, result}
}

var (
	Unauthorization = NewJsonResponse(401, "unauthorization", nil)
	Ok              = NewJsonResponse(200, "ok", nil)
	BadRequest      = NewJsonResponse(400, "bad request", nil)
)
