package handler

// ErrRequest error 多用于偷懒, 定位是 client error
var ErrRequest = &BizError{Code: "ERROR_REQUEST", Message: "request param error"}

// ErrService 服务内部错误
var ErrService = &BizError{Code: "ERROR_SERVICE", Message: "service internal error"}

// ErrorResponse error server 不是服务内部错误, 也不是 client 错误, 就是单纯想回复请求方错误
func ErrorResponse(message string) *BizError {
	return &BizError{Code: "ERROR_RESPONSE", Message: message}
}

func ErrorParam(message string) *BizError {
	return &BizError{Code: "ERROR_PARAM", Message: message}
}

// ErrorInternal 内部错误定义
func ErrorInternal(message string) *BizError {
	return &BizError{Code: "ERROR_INTERNAL", Message: message}
}

// 后面是自定义的业务 Error
