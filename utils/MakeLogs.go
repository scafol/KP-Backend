package utils

import (
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func MakeLogs(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	return log.WithFields(log.Fields{
		"user agent": c.Request().UserAgent(),
		"method":     c.Request().Method,
		"uri":        c.Request().URL.String(),
		"ip":         c.Request().RemoteAddr,
		"at":         time.Now().Format("2006-01-02 15:04:05"),
	})
}
