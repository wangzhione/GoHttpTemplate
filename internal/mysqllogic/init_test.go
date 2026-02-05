// Package mysqllogic 的单元测试：验证 Init 行为及连接是否可用。
package mysqllogic

import (
	"testing"

	"github.com/wangzhione/sbp/chain"
)

var ctx = chain.Context()

// testMySQLCommand 固定测试用 MySQL 命令串（与 mysql 命令行一致）：
const testMySQLCommand = "mysql -uroot -p{passwd} -h{hostname or ip} -P{port} {database} --default-character-set=utf8mb4"

// TestInit_ValidCommand_CanConnect 使用固定 DSN 初始化并 Ping，连接成功则输出成功。
func TestInit_ValidCommand_CanConnect(t *testing.T) {
	err := Init(ctx, testMySQLCommand)
	if err != nil {
		t.Fatalf("Init 失败: %v", err)
	}
	if Main == nil {
		t.Fatal("Init 成功时 Main 不应为 nil")
	}

	db := Main.DB()
	if db == nil {
		t.Fatal("Main.DB() 不应为 nil")
	}
	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("数据库 Ping 失败: %v", err)
	}
	t.Log("MySQL 连接成功")
}
