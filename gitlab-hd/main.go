package main

import (
	logger "github.com/sirupsen/logrus"
	"gitlab-hd/config/db"
	"gitlab-hd/config/gin"
	"gitlab-hd/config/log"
	vc "gitlab-hd/config/viper"
)

func main() {
	// 初始化日志
	log.InitLogConfig()
	logger.Info("===================================================================================")
	logger.Info("Starting Application")

	// 读取本地配置文件
	vc.InitLocalConfigFile()
	// 初始化数据库
	db.InitDB()

	// 初始化Gin
	router := gin.InitGinConfig()
	// 注册router
	gin.InitApiRouter(router)
	//gin server启动
	gin.RunGin(router)
	//优雅关闭Gin
	gin.CloseGin(router)

	//关闭db
	if err := db.Close(); err != nil {
		logger.Fatal("DB关闭异常:", err)
	}
}