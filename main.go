package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "✅ KubeRAG 云原生AI多模态知识库已启动\n")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status":"ok"}`)
	})

	fmt.Println("KubeRAG server started on :8080")
	http.ListenAndServe(":8080", nil)
}
