package register

import (
	"context"
	"log/slog"
	"time"

	"github.com/wangzhione/sbp/system"

	// init logic server register
	_ "github.com/wangzhione/gohttptemplate/internal/logic"
)

func initlogic(ctx context.Context) (err error) {
	defer func() {
		end := time.Now()

		slog.InfoContext(ctx, "logic init end",
			slog.Any("reason", err),
			slog.Float64("elapsed_seconds", end.Sub(system.BeginTime).Seconds()),
		)
	}()

	// do something register logic init here 👇

	// mysql init
	// if err = mysqllogic.Init(ctx, configs.G.MySQL.Main); err != nil {
	// 	return err
	// }

	return
}
