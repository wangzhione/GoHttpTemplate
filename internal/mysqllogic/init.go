// Package mysqllogic provides initialization and management for the main MySQL DB connection.
package mysqllogic

import (
	"context"

	"github.com/wangzhione/sbp/helper/sqler"
	"github.com/wangzhione/sbp/helper/sqler/mysql"
)

// Main Main MySQL DB 对象
var Main *sqler.DB

func Init(ctx context.Context, command string) (err error) {
	// 重复执行会有资源泄露风险, 依赖自行 close Init 管理
	db, err := mysql.NewDB(ctx, command)
	if err != nil {
		return
	}

	if Main != nil {
		Main.Close(ctx)
	}

	Main = db
	return
}
