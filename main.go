package main

import (
	_ "aos/docs"
	"aos/routers"
	"fmt"
	"log"
	"net/http"

	"aos/pkg/setting"

	"aos/pkg/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type ResponseObject struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// TODO:
// 输出HTTP处理日志
// 配置权限、用户状态等对象容器
// 输出RequestID等处理调用链路
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func preSigUsr1() {
	fmt.Println("endless.PRE_SIGNAL ....")
}
func postSigUsr1() {
	fmt.Println("endless.POST_SIGNAL ... ")
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService 127.0.0.1:6001

// @license.name MIT
// @license.url 127.0.0.1:6001

// @BasePath /v1
func main() {

	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20

	// init log
	setting.LoadConfig()
	// init db
	if err := utils.InitEngine(); err != nil {
		fmt.Println("数据库连接异常：", err)
		// panic(err)
		os.Exit(0)
	}

	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	handle := routers.InitRouter()
	server := http.Server{
		Addr:    endPoint,
		Handler: handle,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
