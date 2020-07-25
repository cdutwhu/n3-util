package n3cfg

import "testing"

func TestRmAttrL1(T *testing.T) {
	fPln(
		RmFileAttrL1(
			"../_data/toml/test.toml",
			"./config_part",
			"WebService",
			"LogFile",
			"Storage",
			"File",
		),
	)
}
