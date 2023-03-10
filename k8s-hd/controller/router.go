package controller

import (
	"github.com/gin-gonic/gin"
)

var Router router

type router struct{}

func(r *router) InitApiRouter(router *gin.Engine) {
	router.
		//登录
		POST("/api/login", Login.Auth).
		//集群
		GET("/api/k8s/clusters", Cluster.GetClusters).
		//事件
		GET("/api/k8s/events", Event.GetList).
		//allRes
		GET("/api/k8s/allres", AllRes.GetAllNum).
		//工作流
		//GET("/api/k8s/workflows", Workflow.GetList).
		//GET("/api/k8s/workflow/detail", Workflow.GetById).
		//POST("/api/k8s/workflow/create", Workflow.Create).
		//DELETE("/api/k8s/workflow/del", Workflow.DelById).
		//pod操作
		GET("/api/k8s/pods", Pod.GetPods).
		GET("/api/k8s/pod/detail", Pod.GetPodDetail).
		DELETE("/api/k8s/pod/del", Pod.DeletePod).
		PUT("/api/k8s/pod/update", Pod.UpdatePod).
		GET("/api/k8s/pod/container", Pod.GetPodContainer).
		GET("/api/k8s/pod/log", Pod.GetPodLog).
		GET("/api/k8s/pod/numnp", Pod.GetPodNumPerNp).
		//deployment操作
		GET("/api/k8s/deployments", Deployment.GetDeployments).
		GET("/api/k8s/deployment/detail", Deployment.GetDeploymentDetail).
		PUT("/api/k8s/deployment/scale", Deployment.ScaleDeployment).
		DELETE("/api/k8s/deployment/del", Deployment.DeleteDeployment).
		PUT("/api/k8s/deployment/restart", Deployment.RestartDeployment).
		PUT("/api/k8s/deployment/update", Deployment.UpdateDeployment).
		GET("/api/k8s/deployment/numnp", Deployment.GetDeployNumPerNp).
		POST("/api/k8s/deployment/create", Deployment.CreateDeployment).
		//daemonset操作
		GET("/api/k8s/daemonsets", DaemonSet.GetDaemonSets).
		GET("/api/k8s/daemonset/detail", DaemonSet.GetDaemonSetDetail).
		DELETE("/api/k8s/daemonset/del", DaemonSet.DeleteDaemonSet).
		PUT("/api/k8s/daemonset/update", DaemonSet.UpdateDaemonSet).
		//statefulset操作
		GET("/api/k8s/statefulsets", StatefulSet.GetStatefulSets).
		GET("/api/k8s/statefulset/detail", StatefulSet.GetStatefulSetDetail).
		DELETE("/api/k8s/statefulset/del", StatefulSet.DeleteStatefulSet).
		PUT("/api/k8s/statefulset/update", StatefulSet.UpdateStatefulSet).
		//service操作
		GET("/api/k8s/services", Servicev1.GetServices).
		GET("/api/k8s/service/detail", Servicev1.GetServiceDetail).
		DELETE("/api/k8s/service/del", Servicev1.DeleteService).
		PUT("/api/k8s/service/update", Servicev1.UpdateService).
		POST("/api/k8s/service/create", Servicev1.CreateService).
		//ingress操作
		GET("/api/k8s/ingresses", Ingress.GetIngresses).
		GET("/api/k8s/ingress/detail", Ingress.GetIngressDetail).
		DELETE("/api/k8s/ingress/del", Ingress.DeleteIngress).
		PUT("/api/k8s/ingress/update", Ingress.UpdateIngress).
		POST("/api/k8s/ingress/create", Ingress.CreateIngress).
		//configmap操作
		GET("/api/k8s/configmaps", ConfigMap.GetConfigMaps).
		GET("/api/k8s/configmap/detail", ConfigMap.GetConfigMapDetail).
		DELETE("/api/k8s/configmap/del", ConfigMap.DeleteConfigMap).
		PUT("/api/k8s/configmap/update", ConfigMap.UpdateConfigMap).
		//sercret操作
		GET("/api/k8s/secrets", Secret.GetSecrets).
		GET("/api/k8s/secret/detail", Secret.GetSecretDetail).
		DELETE("/api/k8s/secret/del", Secret.DeleteSecret).
		PUT("/api/k8s/secret/update", Secret.UpdateSecret).
		//pvc操作
		GET("/api/k8s/pvcs", Pvc.GetPvcs).
		GET("/api/k8s/pvc/detail", Pvc.GetPvcDetail).
		DELETE("/api/k8s/pvc/del", Pvc.DeletePvc).
		PUT("/api/k8s/pvc/update", Pvc.UpdatePvc).
		//node操作
		GET("/api/k8s/nodes", Node.GetNodes).
		GET("/api/k8s/node/detail", Node.GetNodeDetail).
		//namespace操作
		GET("/api/k8s/namespaces", Namespace.GetNamespaces).
		GET("/api/k8s/namespace/detail", Namespace.GetNamespaceDetail).
		DELETE("/api/k8s/namespace/del", Namespace.DeleteNamespace).
		//pv操作
		GET("/api/k8s/pvs", Pv.GetPvs).
		GET("/api/k8s/pv/detail", Pv.GetPvDetail).
		DELETE("/api/k8s/pv/del", Pv.DeletePv).
		//helm应用商店
		GET("/api/helmstore/releases", HelmStore.ListReleases).
		GET("/api/helmstore/release/detail", HelmStore.DetailRelease).
		POST("/api/helmstore/release/install", HelmStore.InstallRelease).
		DELETE("/api/helmstore/release/uninstall", HelmStore.UninstallRelease).
		GET("/api/helmstore/charts", HelmStore.ListCharts).
		POST("/api/helmstore/chart/add", HelmStore.AddChart).
		PUT("/api/helmstore/chart/update", HelmStore.UpdateChart).
		DELETE("/api/helmstore/chart/del", HelmStore.DeleteChart).
		POST("/api/helmstore/chartfile/upload", HelmStore.UploadChartFile).
		DELETE("/api/helmstore/chartfile/del", HelmStore.DeleteChartFile)
}