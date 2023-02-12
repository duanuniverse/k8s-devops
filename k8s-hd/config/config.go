package config

const (
	//监听地址
	ListenAddr = "0.0.0.0:9090"
	Kubeconfigs = `{"TEST-1":".\\config\\k8s-config","TEST-2":".\\config\\k8s-config"}`
	//pod日志tail显示行数
	PodLogTailLine = 2000
)