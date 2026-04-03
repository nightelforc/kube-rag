package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"kube-rag/internal/handler"
	"kube-rag/pkg/logger"
)

func main() {
	// 初始化日志
	logger.InitLogger()

	// 初始化配置
	initConfig()

	// Gin 引擎
	r := gin.Default()

	// 健康检查（K8s 必须）
	r.GET("/health", handler.Health)
	r.GET("/ready", handler.Ready)

	// 业务路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("/upload", handler.UploadFile) // 文件上传
	}

	// 启动
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}
	logger.Info("🚀 KubeRAG 启动成功，端口：" + port)
	_ = r.Run(":" + port)
}

// initConfig 读取配置文件
func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")

	if err := viper.ReadInConfig(); err != nil {
		logger.Warn("未找到配置文件，使用默认配置")
	}
}
