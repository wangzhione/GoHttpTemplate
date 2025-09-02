package handler

import (
	"fmt"
)

// BizError 业务错误结构体
type BizError struct {
	Code    string // 错误代码
	Message string // 错误消息
	Err     error  // 首次 error
}

// Error 实现 error 接口
func (e *BizError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *BizError) Unwrap() error { return e.Err }

// GetCode 获取错误代码
func (e *BizError) GetCode() string {
	return e.Code
}

// GetMessage 获取 error message
func (e *BizError) GetMessage() string {
	if len(e.Message) != 0 {
		return e.Message
	}

	if e.Err != nil {
		return fmt.Sprintf("biz error code [%s] %s", e.Code, e.Err.Error())
	}
	return fmt.Sprintf("biz error code [%s]", e.Code)
}

// NewBizError 创建新的 BizError
func NewBizError(code, message string, errargs ...error) *BizError {
	bizerr := &BizError{
		Code:    code,
		Message: message,
	}

	for _, err := range errargs {
		if err != nil {
			bizerr.Err = err // 默认只纪录第一个, 就是个可变参数写法
			break
		}
	}

	return bizerr
}
