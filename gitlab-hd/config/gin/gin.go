package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab-hd/config/log"
	err "gitlab-hd/exception"
	"gitlab-hd/a_gitlab/controller"
	"gitlab-hd/middle"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// InitGinConfig 初始化Gin
func InitGinConfig() *gin.Engine {
	logger.Info("初始化 gin……")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 入口日志打印
	router.Use(log.LoggerAccess)
	// 统一异常处理
	router.Use(err.ErrHandle)
	// 跨域处理
	router.Use(middle.Cors())
	//jwt token验证
	//router.Use(middle.JWTAuth())
	// token校验
	//router.Use(token.TokenVerify)
	logger.Info("Gin: 初始化完成……")
	return router
}

//设置路由
func InitApiRouter(router *gin.Engine) {
	// 健康检测
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userGroup := router.Group("project/")
	{
		userGroup.
			GET("/:id", controller.Project.FindProjectByID)
	}
	// 健康检测
	userGroup = router.Group("test/")
	{
		userGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	/*router.
		GET("/api/gitlab/project/getlist", project.GetAllGitProject)*/
}

//定义http server值，保内调用
func HttpServer(router *gin.Engine) (srv *http.Server, listenAddr string){
	listenAddr = viper.GetString("server.ListenAddr")
	srv = &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}
	return srv, listenAddr
}

// RunGin 启动Gin
func RunGin(router *gin.Engine) {
	srv, listenAddr := HttpServer(router)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()
	logger.Info(fmt.Sprintf("Service started on ListenAddr(s): %s", listenAddr))
}

// Close Gin 关闭Gin
func CloseGin(router *gin.Engine) {
	//等待中断信号，优雅关闭所有server及DB
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	//设置ctx超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//cancel用于释放ctx
	defer cancel()

	//关闭gin server
	srv, _ := HttpServer(router)
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Gin Server关闭异常:", err)
	}
	logger.Println("Gin Server退出成功")
}