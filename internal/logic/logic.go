// Package logic provides the core business logic for the GoHttpTemplate application.
package logic

import (
	"net/http"
	"time"

	"github.com/wangzhione/gohttptemplate/handler"
	"github.com/wangzhione/sbp/system"
)

func init() {
	http.HandleFunc("/health", handler.GET(Health))
}

// HealthResp 定义健康检查的返回结构
type HealthResp struct {
	Version   string `json:"version"`   // 版本标识
	Timestamp string `json:"timestamp"` // 服务器运行当前时间
}

// Health 处理健康检查请求
func Health(*http.Request) (resp *HealthResp, err error) {
	resp = &HealthResp{
		Timestamp: time.Now().Format(time.RFC3339Nano),
		Version:   system.GitVersion,
	}
	return
}
