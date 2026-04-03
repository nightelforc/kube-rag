FROM golang:1.25.6-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o kube-rag ./cmd/server

# 运行阶段（极小体积）
FROM scratch
COPY --from=builder /app/kube-rag /kube-rag
COPY --from=builder /app/configs /configs
COPY --from=builder /app/uploads /uploads
EXPOSE 8080
ENTRYPOINT ["/kube-rag"]