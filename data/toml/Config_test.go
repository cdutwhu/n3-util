package toml

import (
	"fmt"
	"testing"

	"github.com/cdutwhu/gotil/rflx"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
)

func TestNewCfg(t *testing.T) {
	cfg := NewCfg(
		"Config1",
		map[string]string{
			"[p]": "Port",
			"[s]": "Service",
			"[v]": "WebService.Version",
		},
		"./test.toml",
	).(*Config1)
	spew.Dump(*cfg)

	fmt.Println(" ------------------------- ")

	cfg1 := rflx.Env2Struct("Config1", &Config1{}).(*Config1)
	spew.Dump(*cfg1)

	if cmp.Equal(*cfg, *cfg1) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
