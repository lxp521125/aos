package panicHandle

import (
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"aos/pkg/setting"
	"aos/pkg/errors"
)

type HTTPError interface {
	HTTPStatus() int
}

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stackInfo := "stack info:" + string(debug.Stack())
				setting.Logger.WithField("stackInfo", stackInfo).Infof("", stackInfo)
				c.JSON(200, errors.New(errors.SYSERR, errors.GetInfo()[errors.SYSERR]))

				c.Abort()
			}
		}()
		c.Next()
	}
}
