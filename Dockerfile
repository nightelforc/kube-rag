FROM golang:1.25.6-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o kube-rag ./cmd/server

FROM scratch
COPY --from=builder /app/kube-rag /kube-rag
EXPOSE 8080
CMD ["/kube-rag"]