package res

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

const (
	errorCode = -1
	success   = 0
	msg       = "操作成功"
)

func Ok() *Response {
	return newResponse(success, msg, nil)
}

func OkCode(code int) *Response {
	return newResponse(code, msg, nil)
}
func OkMsg(msg string) *Response {
	return newResponse(success, msg, nil)
}
func OkMsgCode(msg string, code int) *Response {
	return newResponse(code, msg, nil)
}
func OkMsgData(code int, msg string, data interface{}) *Response {
	return newResponse(code, msg, data)
}
func OkData(data interface{}) *Response {
	return newResponse(success, msg, data)
}
func OkDataCode(data interface{}, code int) *Response {
	return newResponse(code, msg, data)
}
func Err(msg string) *Response {
	return newResponse(errorCode, msg, nil)
}
func newResponse(code int, message string, data interface{}) *Response {
	if message == "" {
		message = msg
	}
	return &Response{
		Code: code,
		Data: data,
		Msg:  message,
	}
}
