package n3cfg

import "testing"

func TestRmCfgAttrL1(T *testing.T) {
	fPln(
		RmCfgAttrL1(
			"../_data/toml/test.toml",
			"./config_rm",
			"Path",
			"WebService",
			"Port",
			"Storage",
			"File",
		),
	)
}

func TestSelCfgAttrL1(T *testing.T) {
	fPln(
		SelCfgAttrL1(
			"../_data/toml/test.toml",
			"./config_sel",
			"Path",
			"WebService",
			"LogFile",
			"Storage",
			"File",
		),
	)
}
