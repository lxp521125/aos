package setting

import (
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
)

type hl string

func (*hl) HandleLog(e *log.Entry) error {
	return nil
}

var l hl = "nil"

func GinLogger(c *gin.Context) *log.Entry {
	if l, ok := c.Get("logger"); ok {
		 return l.(*log.Entry)
	} else {
		return GrayLog()
	}
}
