package responses

type ErrorResponse struct {
	Detail []ErrorInfo `json:"detail"`
}

type ErrorInfo struct {
	Msg  string `json:"msg"`
	Type string `json:"type"`
}
