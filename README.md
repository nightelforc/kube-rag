# kube-rag

基于 Go + K8s + 大模型构建的云原生多模态 RAG 知识库平台

## 技术栈

- Go 1.25+
- go-zero
- Docker & Kubernetes
- Kafka

## 功能

- 多模态文档解析与存储
- 向量检索
- 智能问答
- AI Agent

# 目录结构

```
kube-rag/
├── cmd/
│   └── server/        # 服务启动入口
│       └── main.go
├── configs/           # 配置文件
│   └── config.yaml
├── internal/          # 内部业务（对外不可见）
│   ├── config/        # 配置结构体
│   ├── handler/       # HTTP 接口层
│   ├── service/       # 业务逻辑层
│   └── pkg/           # 内部工具
├── pkg/               # 公共工具（可复用）
│   ├── logger/        # 日志
│   └── response/      # 统一返回
├── uploads/           # 文件上传目录
├── Dockerfile
├── go.mod
└── README.md
```