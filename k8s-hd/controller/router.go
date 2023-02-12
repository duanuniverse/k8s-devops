package controller

import (
	"github.com/gin-gonic/gin"
)

//实例化router对象，可使用该对象点出首字母大写的方法(跨包调用)
var Router router

//定义router结构体
type router struct{}

//初始化路由，创建测试API接口
func(r *router) InitApiRouter(router *gin.Engine) {
	router.
		//pod操作
		GET("/api/k8s/pods", Pod.GetPods).
		GET("/api/k8s/pod/detail", Pod.GetPodDetail).
		DELETE("/api/k8s/pod/del", Pod.DeletePod).
		PUT("/api/k8s/pod/update", Pod.UpdatePod).
		GET("/api/k8s/pod/container", Pod.GetPodContainer).
		GET("/api/k8s/pod/log", Pod.GetPodLog).
		GET("/api/k8s/pod/numnp", Pod.GetPodNumPerNp)
}