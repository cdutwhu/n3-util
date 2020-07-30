package n3cfg

import "testing"

func TestSelFileAttrL1(T *testing.T) {
	fPln(
		SelFileAttrL1(
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
