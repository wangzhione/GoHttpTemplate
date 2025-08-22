// Package register provides initialization routines for the application environment.
package register

import (
	"context"
	"log/slog"
	"runtime"

	_ "net/http/pprof"

	_ "go.uber.org/automaxprocs"

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
		slog.Int("G.Serve.PNumber", configs.G.Serve.PNumber),
	)

	// 在 Docker or Kubernetes 程序获取的是宿主机的 CPU 核数导致 GOMAXPROCS 设置的过大。
	// 比如宿主物理机是 48cores，而实际 container 只有 4cores。P 导致系统线程过多，会增加上线文切换的负担，白白浪费 CPU。
	// automaxprocs 通过读取系统的 CPU 核心数和 Cgroups 限制等信息来动态调整 P 的数量
	// 自适应设置 procs @https://github.com/uber-go/automaxprocs
	if configs.G.Serve.PNumber > 0 {
		runtime.GOMAXPROCS(configs.G.Serve.PNumber)
	}

	// init 操作, 放在这后面 👇

	return
}
