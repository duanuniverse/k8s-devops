package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"k8s-hd/service"
	"net/http"
)

var Secret secret

type secret struct {}
//获取secret列表，支持过滤、排序、分页
func(s *secret) GetSecrets(ctx *gin.Context) {
	params := new(struct {
		FilterName  string `form:"filter_name"`
		Namespace   string `form:"namespace"`
		Page        int    `form:"page"`
		Limit       int    `form:"limit"`
		Cluster     string `form:"cluster"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Secret.GetSecrets(client, params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "获取Secret列表成功",
		"data": data,
	})
}

//获取secret详情
func(s *secret) GetSecretDetail(ctx *gin.Context) {
	params := new(struct {
		SecretName    string `form:"secret_name"`
		Namespace     string `form:"namespace"`
		Cluster       string `form:"cluster"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Secret.GetSecretDetail(client, params.SecretName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "获取Secret详情成功",
		"data": data,
	})
}

//删除secret
func(s *secret) DeleteSecret(ctx *gin.Context) {
	params := new(struct{
		SecretName   string  `json:"secret_name"`
		Namespace    string  `json:"namespace"`
		Cluster      string  `json:"cluster"`
	})
	//DELETE请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	err = service.Secret.DeleteSecret(client, params.SecretName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "删除Secret成功",
		"data": nil,
	})
}

//更新secret
func(s *secret) UpdateSecret(ctx *gin.Context) {
	params := new(struct{
		Namespace       string  `json:"namespace"`
		Content         string  `json:"content"`
		Cluster         string  `json:"cluster"`
	})
	//PUT请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	err = service.Secret.UpdateSecret(client, params.Namespace, params.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新Secret成功",
		"data": nil,
	})
}