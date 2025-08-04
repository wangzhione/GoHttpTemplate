// Package middleware 拦截器主要用于所有请求, 注入 trace id, 统一处理 panic
package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/wangzhione/sbp/chain"
)

// ResponseWriterPanicError 自定义 5xx 当成 panic error
func ResponseWriterPanicError(w http.ResponseWriter) {
	w.WriteHeader(599)
	fmt.Fprintf(w, `{"code":599, "message:"Internal Server Panic Error"}`)
}

// MainMiddleware 拦截器, 默认 serve 所以拦截器集中在这里
func MainMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // 记录请求开始时间

		// 获取或生成 requestID
		r, requestID := chain.Request(r)

		// 使用 `defer` 捕获 `panic`
		defer func() {
			if err := recover(); err != nil {
				// 记录 `panic` 错误日志
				slog.ErrorContext(r.Context(), "HTTP request handler panic error",
					slog.String("method", r.Method),
					slog.String("path", r.URL.Path),
					slog.String("start", start.Format(time.RFC3339Nano)),
					slog.Float64("elapsed", time.Since(start).Seconds()),
					slog.Any("error", err),
					slog.String("stack", string(debug.Stack())), // 记录详细的堆栈信息
				)

				// 返回 自定义 panic error 服务器内部错误响应
				ResponseWriterPanicError(w)
			}

			// 确保 r.Body close, 避免 底层 net.Conn 长时间占用, 影响连接复用
			r.Body.Close()

			// 记录请求完成日志
			slog.InfoContext(r.Context(), "HTTP request handler completed",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("start", start.Format(time.RFC3339Nano)),
				slog.Float64("elapsed", time.Since(start).Seconds()),
			)
		}()

		// set 具体的 X-Request-Id
		w.Header().Set(chain.XRquestID, requestID)

		// Step 1: Header 特殊处理逻辑, 主要为了放开跨域问题

		// 默认本服务处理の类型
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		// 防止 MIME 类型嗅探（MIME-Type Sniffing）
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// 默认全添加 CORS 头, 默认允许跨域访问
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin) // 允许所有来源
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // 允许跨域携带 Cookie

		// 处理预检请求
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Step 2: serve http call logic 之前额外操作

		next.ServeHTTP(w, r)

		// Step 3: serve http call logic 之后额外操作
	})
}
