// Package configs provides configuration loading and management for the application.
package configs

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/wangzhione/sbp/util/tuml"
)

var G *Config

// Config 结构体对应 resource/etc/prod.toml
type Config struct {
	Env struct {
		Mode string `toml:"mode"`
	} `toml:"env"`

	Serve struct {
		Port        uint16        `toml:"port"`
		StopTime    time.Duration `toml:"stoptime"`    // 停止时间配置格式 7s 或 2000ms
		GOMAXPROCS  int           `toml:"gomaxprocs"`  // GOMAXPROCS 数量
		PprofBearer string        `toml:"pprofbearer"` // Pprof Bearer Token, 为空则禁用 pprof 访问
	} `toml:"serve"`

	Log struct {
		Level string `toml:"level"`
	} `toml:"log"`

	MySQL struct {
		Main string `toml:"main"`
	} `toml:"mysql"`
}

// Init 初始化配置 G
func Init(ctx context.Context, path string) (err error) {
	cfg, err := tuml.ReadFile[*Config](path) // 解析 TOML 配置文件
	if err != nil {
		slog.ErrorContext(ctx, "Failed to read config file", "path", path, "error", err)
		return
	}

	if cfg.Env.Mode == "" {
		cfg.Env.Mode = "local"
	} else {
		cfg.Env.Mode = strings.ToLower(cfg.Env.Mode)
	}

	if cfg.Log.Level == "" {
		cfg.Log.Level = "INFO"
	} else {
		cfg.Log.Level = strings.ToUpper(cfg.Log.Level)
	}

	// 处理默认值
	if cfg.Serve.Port <= 0 {
		cfg.Serve.Port = 8089 // 默认的服务端口, 可以随着配置改变
	}

	if cfg.Serve.StopTime <= 0 {
		cfg.Serve.StopTime = 7000 * time.Millisecond // 服务等待时间, 默认 7 s
	}

	// 构建全局的是否是线上环境变量
	Online = cfg.IsOnline()

	G = cfg
	return
}

// IsOnline 判断当前是否为线上环境
func (cfg *Config) IsOnline() bool {
	switch cfg.Env.Mode {
	case "prod", "production", "online", "preview", "pre":
		return true
	default:
		return false
	}
}
