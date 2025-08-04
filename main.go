package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/wangzhione/gohttptemplate/global"
	"github.com/wangzhione/gohttptemplate/global/config"
	"github.com/wangzhione/gohttptemplate/handler/middleware"
	"github.com/wangzhione/gohttptemplate/register"
	"github.com/wangzhione/sbp/system"
)

// pathG 默认配置文件地址
var pathG = flag.String("f", "resource/etc/prod.toml", "The config file")

func main() {
	flag.Parse() // flag 参数初始化

	defer system.End(global.BC)

	// init 如果失败, 程序会直接退出
	err := register.Init(global.BC, *pathG)
	if err != nil {
		os.Exit(-1)
	}

	system.ServeLoop(
		global.BC,
		fmt.Sprintf("0.0.0.0:%d", config.G.Serve.Port), // 0.0.0.0 默认 ipv4 绑定本机地址
		middleware.MainMiddleware(http.DefaultServeMux),
		config.G.Serve.StopTime,
	)
}
