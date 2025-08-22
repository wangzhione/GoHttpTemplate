package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/wangzhione/gohttptemplate/configs"
	"github.com/wangzhione/gohttptemplate/handler/middleware"
	"github.com/wangzhione/gohttptemplate/register"
	"github.com/wangzhione/sbp/chain"
	"github.com/wangzhione/sbp/system"
)

func main() {
	ctx := chain.Context()

	// path 默认配置文件地址
	resourcePath := flag.String("f", "resource/etc/prod.toml", "The config file")

	flag.Parse() // flag 参数初始化

	defer system.End(ctx)

	// init 如果失败, 程序会直接退出
	err := register.Init(ctx, *resourcePath)
	if err != nil {
		os.Exit(-1)
	}

	system.ServeLoop(
		ctx,
		fmt.Sprintf("0.0.0.0:%d", configs.G.Serve.Port), // 0.0.0.0 默认 ipv4 绑定本机地址
		middleware.MainMiddleware(http.DefaultServeMux),
		configs.G.Serve.StopTime,
	)
}
