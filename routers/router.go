package routers

import (
	"aos/bindService"
	"aos/middleware/logs"
	"aos/middleware/panicHandle"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "aos/docs"

	"aos/pkg/setting"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://xxx.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	r.Use(logs.Logger())
	r.Use(gin.Recovery())
	r.Use(panicHandle.CatchError())

	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	containerInstance := container.GetContainer()
	c1 := r.Group("/v1")

	//test
	c1.GET("/graylog", containerInstance.TestApi.TestGraylog)

	//static
	r.Static("/public", "./public")
	return r
}
