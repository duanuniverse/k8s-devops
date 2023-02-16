package main

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"go_web_test/biz/controller"
	"go_web_test/biz/dao"
	"go_web_test/config/db"
	"go_web_test/config/gin"
	"go_web_test/config/log"
	vc "go_web_test/config/viper"
	_ "net/url"
)

func main() {
	initComponents()
}

// 初始化服务所有组件
func initComponents() {
	// 初始化日志
	log.InitLogConfig()
	logger.Info("===================================================================================")
	logger.Info("Starting Application")
	// 读取本地配置文件
	vc.InitLocalConfigFile()
	// 初始化url配置
	//url.InitUrlConfig()
	// 初始化Mysql
	db.InitDbConfig()
	// 自动生成表
	autoMigrate()
	// 初始化缓存
	//cache.InitBigCacheConfig()
	// 初始化Redis
	//redis.InitRedisConfig()
	// 初始化HttpClient连接池
	//http.InitHttpClientConfig()

	// 初始化token
	//token.InitTokenConfig()

	// 初始化Gin
	router := gin.InitGinConfig()

	// 注册Api
	// 用户api
	controller.UserApi(router)

	// 启动Gin
	gin.RunGin(router)
}

// 自动生成表
func autoMigrate() {
	err := db.DB.AutoMigrate(dao.User{})
	if err != nil {
		_ = fmt.Errorf("自动生成user表失败")
		panic(err)
	}

}
