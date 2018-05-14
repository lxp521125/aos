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
	c1.GET("/secret/:access_key", containerInstance.TestApi.GetS)
	c1.GET("/dbtest", containerInstance.TestApi.GetDbTest)
	c1.GET("/servicetest", containerInstance.TestApi.GetServiceTest)
	c1.GET("/graylog", containerInstance.TestApi.TestGraylog)


	// video api
	c1.POST("/video/batch", containerInstance.VideoApi.Batch)
	c1.GET("/ping", containerInstance.VideoApi.Ping)

	// lecture note api
	c1.POST("/lecture/batch", containerInstance.LectureNoteApi.Batch)

	//static
	r.Static("/public", "./public")
	return r
}
