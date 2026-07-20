package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{
	{ID: 1, Title: "看懂 JSON 响应", Done: false},
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Go First")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "pong")
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		log.Printf("写入 JSON 响应失败: %v", err)
	}
}

func main() {
	http.HandleFunc("GET /hello", helloHandler)
	http.HandleFunc("GET /ping", pingHandler)
	http.HandleFunc("GET /todos", todosHandler)

	log.Println("服务已启动：http://localhost:8080/hello")
	log.Println("第三关入口：http://localhost:8080/todos")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
