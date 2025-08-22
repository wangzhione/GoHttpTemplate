package handler

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrRequestEmpty 返回的数据是 empty
var ErrRequestEmpty = errors.New("error: request is empty")

// ErrRequestInvalid request 参数异常
var ErrRequestInvalid = errors.New("error: request is invalid")

// ErrTimeout 内部超时
var ErrTimeout = errors.New("error: timeout")

// ErrRequestURLInvalid request URL 参数异常
var ErrRequestURLInvalid = errors.New("error: request url is invalid")

// ErrCallInvalid 请求异常
var ErrCallInvalid = errors.New("error: call is invalid")

// Response 统一返回 response base model
type Response[T any] struct {
	Code    int    `json:"code"`              // [200, 299] 都认为是 OK, 默认都是 200
	Message string `json:"message,omitempty"` // 补充信息
	Data    T      `json:"data,omitzero"`
}

// Coder 定义接口 GetCode 和 GetMessage 接口
type Coder interface {
	GetCode() int
}

type Messager interface {
	GetMessage() string
}

func NewResponse[T any](resp T, err error) (response *Response[T]) {
	response = &Response[T]{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    resp,
	}

	if err == nil {
		return
	}

	// 获取 error message {text} and code {status}
	if coder, ok := err.(Coder); ok {
		response.Code = coder.GetCode()
	} else {
		// 默认 400
		response.Code = 400
	}

	if messager, ok := err.(Messager); ok {
		response.Message = messager.GetMessage()
	} else {
		response.Message = err.Error()
	}

	return
}

// OK 请求是否成功
func (resp *Response[T]) OK() bool {
	return resp.Code >= http.StatusOK && resp.Code < http.StatusMultipleChoices
}

// Error 构建 error
func (resp *Response[T]) Error() error {
	if resp.OK() {
		return nil
	}
	return errors.New(resp.ErrorString())
}

// ErrorString 得到 Error 错误文本
func (resp *Response[T]) ErrorString() string {
	if resp.OK() {
		return "OK"
	}

	if len(resp.Message) != 0 {
		return resp.Message
	}

	message := http.StatusText(resp.Code)
	if len(message) != 0 {
		return message
	}

	return fmt.Sprintf("response error code <%d>", resp.Code)
}

// ResponseWriterMethodError ResponseError 构建 error string, code 必须是严格 http code
func ResponseWriterMethodError(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"code":%d, "message:"%s"}`, code, http.StatusText(code))
}

func ResponseWriterMessage(w http.ResponseWriter, message string) {
	fmt.Fprintf(w, `{""message:"%s"}`, message)
}

func (resp *Response[T]) SetOK(data T) {
	resp.Code = http.StatusOK
	resp.Message = "OK"
	resp.Data = data
}
