package n3toml

import "testing"

func TestRmAttrL1(T *testing.T) {
	fPln(
		RmFileAttrL1(
			"../_data/toml/test.toml",
			"./test",
			// "WebService",
			"LogFile",
			"Storage",
			"File",
		),
	)
}
