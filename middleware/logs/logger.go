package logs

import (
	"time"

	"aos/pkg/setting"
	"aos/pkg/tool"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		tool.NewUniqueIDAsync()
		uuid := tool.GetUID()

		logger := setting.GrayLog()
		logger = logger.WithField("X-Response-ID", uuid)
		c.Set("logger", logger)
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		logger.Infof("interface url:"+c.Request.URL.Path+"，times:", latency)
	}
}
