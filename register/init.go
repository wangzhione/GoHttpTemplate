// Package register provides initialization routines for the application environment.
package register

import (
	"context"
	"log/slog"
	"runtime"

	_ "net/http/pprof"

	"github.com/wangzhione/sbp/chain"
	"github.com/wangzhione/sbp/system"

	"github.com/wangzhione/gohttptemplate/configs"

	_ "github.com/wangzhione/gohttptemplate/internal/logic"
)

// Init 启动之前的环境初始化 :)
func Init(ctx context.Context, path string) (err error) {
	// init config
	if err = configs.Init(ctx, path); err != nil {
		return
	}

	// slog 默认配置初始化
	switch configs.G.Log.Level {
	case "info":
		chain.EnableLevel = slog.LevelInfo
	case "warn":
		chain.EnableLevel = slog.LevelWarn
	case "error":
		chain.EnableLevel = slog.LevelError
	}
	if err = chain.InitSlogRotatingFile(); err != nil {
		// 如果 文件 日志有问题, 需要打印相关信息
		slog.ErrorContext(ctx, "chain.InitSlogRotatingFile error", "error", err) // 退化成控制台输出
	}

	// 输出 CPU Core 的数量, 输出处理器 P 的数量, 如果是容器, 像个数据不一定准确
	slog.InfoContext(ctx, "main init start ...",
		slog.Time("SystemBeginTime", system.BeginTime),
		slog.Int("cpunumber", runtime.NumCPU()),
		slog.Int("pnumber", runtime.GOMAXPROCS(0)),
		slog.String("path", path),
		slog.String("GOOS", runtime.GOOS),
		slog.String("BuildVersion", system.BuildVersion),
		slog.String("GitVersion", system.GitVersion),
		slog.String("GitCommitTime", system.GitCommitTime),
	)

	// init 操作, 放在这后面 👇

	return
}
