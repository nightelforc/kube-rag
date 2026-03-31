package main

import (
	"kube-rag/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	zap.L().Info("Kube-RAG Server Started")
}
