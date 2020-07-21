package n3log

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLoggly(t *testing.T) {
	LrInit()
	EnableLoggly(true)
	SetLogglyToken("54290728-93e0-425a-a49c-c9e834288026")
	if _, err := LrOut(logrus.Errorf, "loggly test Errorf"); err != nil {
		fmt.Println(err)
	}
}
