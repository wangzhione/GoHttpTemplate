// Package configs provides configuration loading and management for the application.
package configs

import (
	"context"
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
		Port     uint16        `toml:"port"`
		StopTime time.Duration `toml:"stoptime"` // 停止时间配置格式 7s 或 2000ms
		PNumber  int           `toml:"pnumber"`
	} `toml:"serve"`

	Log struct {
		Level string `toml:"level"`
	} `toml:"log"`

	MySQL struct {
		Main string `toml:"main"`
	} `toml:"mysql"`
}

var (
	// DefaultServePort 默认的服务端口, 可以随着配置改变
	DefaultServePort uint16 = 8089

	// DefaultServeStopTime 服务等待时间, 默认 7 s
	DefaultServeStopTime time.Duration = 7000 * time.Millisecond
)

// Init 初始化配置 G
func Init(ctx context.Context, path string) (err error) {
	cfg, err := tuml.ReadFile[*Config](ctx, path) // 解析 TOML 配置文件
	if err != nil {
		return
	}

	// 处理默认值
	if cfg.Serve.Port <= 0 {
		cfg.Serve.Port = DefaultServePort
	}

	if cfg.Serve.StopTime <= 0 {
		cfg.Serve.StopTime = DefaultServeStopTime
	}

	G = cfg
	return
}
