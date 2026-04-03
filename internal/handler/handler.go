package handler

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"kube-rag/pkg/logger"
	"kube-rag/pkg/response"
)

// Health 健康检查
func Health(c *gin.Context) {
	response.Success(c, gin.H{
		"service": "kube-rag",
		"status":  "running",
	})
}

// Ready 就绪检查（K8s 用）
func Ready(c *gin.Context) {
	response.Success(c, gin.H{"ready": true})
}

// UploadFile 文件上传
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Error("上传文件失败：", zap.Error(err))
		response.Fail(c, "文件获取失败")
		return
	}

	// 保存目录
	savePath := "./uploads/"
	fileName := filepath.Base(file.Filename)
	fullPath := savePath + fileName

	// 创建目录
	_ = c.SaveUploadedFile(file, fullPath)

	logger.Info("文件上传成功：", zap.String("fileName", fileName))
	response.Success(c, gin.H{
		"filename": fileName,
		"path":     fullPath,
	})
}
