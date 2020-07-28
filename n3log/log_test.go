package n3log

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestLoggly(t *testing.T) {
	e := echo.New()
	defer e.Close()
	defer e.Start(fSf(":%d", 1500))

	LrInit()
	EnableLoggly(true)
	SetLogglyToken("54290728-93e0-425a-a49c-c9e834288026")
	Bind(logger, lPf, e.Logger.Infof, Loggly("info")).Do("%s", "echo starting")
}
