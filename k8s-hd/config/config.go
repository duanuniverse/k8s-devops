package config

import "time"

const (
	//监听地址
	ListenAddr = "0.0.0.0:9090"
	WsAddr = "0.0.0.0:8081"
	//kubeconfig路径
	//Kubeconfig = "F:\\goproject\\config"  windows路径
	Kubeconfigs = `{"TEST-1":".\\config\\k8s-config","TEST-2":".\\config\\k8s-config"}`
	//pod日志tail显示行数
	PodLogTailLine = 2000
	//登录账号密码
	AdminUser = "admin"
	AdminPwd = "123123"

	//数据库配置
	DbType = "mysql"
	DbHost = "192.168.3.179"
	DbPort = 3306
	DbName = "k8s-devops"
	DbUser = "root"
	DbPwd = "fkR/CCI7T2Um"
	//打印mysql debug sql日志
	LogMode = false
	//连接池配置
	MaxIdleConns = 10 //最大空闲连接
	MaxOpenConns = 100 //最大连接数
	MaxLifeTime = 30 * time.Second //最大生存时间
	//helm配置
	UploadPath = "/Users/adoo/chart"
)
