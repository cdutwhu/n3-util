package n3log

import (
	"testing"
)

func TestLoggly(t *testing.T) {
	// e := echo.New()
	// defer e.Close()
	// defer e.Start(fSf(":%d", 1500))

	// e.GET("/", func(c echo.Context) error {
	// 	e.Logger.Infof("e.GET")
	// 	return c.String(http.StatusOK, "echo return")
	// })

	SetLoggly(true, "54290728-93e0-425a-a49c-c9e834288026", "n3log")
	SyncBindLog(false)
	Bind(logger, lPf, Loggly("info")).Do("%s", "TestLogglyBind ASyncBindLog")
	SyncBindLog(true)
	Bind(logger, lPf, Loggly("info")).Do("%s", "TestLogglyBind SyncBindLog")
}
