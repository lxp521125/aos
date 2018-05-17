package main

import (
	_ "aos/docs"
	"aos/routers"
	"fmt"
	"net/http"

	"aos/pkg/setting"

	"aos/pkg/utils"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

type ResponseObject struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// TODO: 处理HTTP响应包括错误的公共方法
func Dump(c *gin.Context, err error, object interface{}) {
	responseObject := ResponseObject{
		1,
		err.Error(),
		object,
	}
	if nil != err {
		c.JSON(http.StatusOK, responseObject)
	} else {
		c.JSON(http.StatusOK, responseObject)
	}
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
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt,
		syscall.SIGTERM, syscall.SIGINT,
		syscall.SIGHUP, syscall.SIGABRT)

	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	// init log
	setting.LoadConfig()
	// init db
	if err := utils.InitEngine(); err != nil {
		panic(err)
		os.Exit(0)
	}
	handle := routers.InitRouter()
	server := http.Server{
		Addr:           endPoint,
		Handler:        handle,
		WriteTimeout:   setting.WriteTimeout,
		ReadTimeout:    setting.ReadTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Server err: %v", err)
		}
	}()
	<-exit
	if err := server.Close(); err != nil {
		fmt.Printf("server close error %s", err)
	}
	fmt.Println("server close safe")
}
