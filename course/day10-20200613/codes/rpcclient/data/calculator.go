package data

// calculator service请求对象
type CalculatorRequest struct {
	Left int
	Right int
}

// calculator service响应对象
type CalculatorResponse struct {
	Result int
}
