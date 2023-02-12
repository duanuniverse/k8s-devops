package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"k8s-hd/config"
	"k8s-hd/controller"
	"k8s-hd/service"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//初始化gin对象
	r := gin.Default()
	//初始化k8s client
	service.K8s.Init()
	//初始化路由
	controller.Router.InitApiRouter(r)
	//启动gin server
	srv := &http.Server{
		Addr: config.ListenAddr,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen: %s\n", err)
		}
	}()
	//优雅关闭server
	//声明一个系统信号的channel，并监听塔；如果没有信号就一直阻塞；如果有信号就继续执行
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	//设置ctx(连接)超时时间，超过设置时间如果还有连接就强制关闭所以连接
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	//cancel用于释放ctx
	defer cancel()
	//关闭gin
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Gin Server关闭异常：", err)
	}
	logger.Info("Gin Server退出成功")
}