package configs

import (
	"testing"

	"github.com/wangzhione/sbp/chain"
	"github.com/wangzhione/sbp/util/jsou"
	"github.com/wangzhione/sbp/util/tuml"
)

var ctx = chain.Context()

func TestConfigParse(t *testing.T) {
	path := `../../resource/etc/prod.toml`

	cfg, err := tuml.ReadFile[*Config](ctx, path)
	if err != nil {
		t.Fatal("ConfigParse is fatal", path, err)
	}

	t.Log("Success config parse", jsou.String(cfg))
}
