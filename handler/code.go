package handler

// code : 123 45 6789
// 只用 code 9 位 {3}{2}{4} 相当于 {400, 500} + {biz id | 1 - 99} + { error id | 1-9999}
//
// 粗粒度设计
// 400+{biz id | 1 - 99} + { error id | 1-9999} 一般形容客户端割裂请求出错
// 500+{biz id | 1 - 99} + { error id | 1-9999} 一般是服务器内部各类错误
// biz id = 0 框架自用

// 继承 go/src/net/http/status.go 中 HTTP status codes

// 大部分业务码 code 意义不大, 特殊业务码自行定义, 联动运维, 前端, 监控报警

// bizerror::code
// 200-299 is OK
// 400-499 参数错误
// 		400{xxxxxx} - 449{999999}
//
// 500-599 服务内部错误

// ErrRequest error 多用于偷懒, 定位是 client error
var ErrRequest = &BizError{Code: 400000404, Message: "request param error"}

// ErrService 服务内部错误
var ErrService = &BizError{Code: 500000501, Message: "service internal error"}

// ErrorResponse error server 不是服务内部错误, 也不是 client 错误, 就是单纯想回复请求方错误
func ErrorResponse(message string) *BizError {
	return &BizError{Code: 450000001, Message: message}
}

func ErrorParam(message string) *BizError {
	return &BizError{Code: 400000002, Message: message}
}

// ErrorInternal 内部错误定义
func ErrorInternal(message string) *BizError {
	return &BizError{Code: 500000601, Message: message}
}

// 后面是自定义的业务 Error
