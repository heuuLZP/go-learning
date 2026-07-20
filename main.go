package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{
	{ID: 1, Title: "看懂 JSON 响应", Done: false},
	{ID: 2, Title: "增加第二项", Done: false},
}

var nextTodoID = 3

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

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "请求 JSON 无效", http.StatusBadRequest)
		return
	}

	todo.ID = nextTodoID
	nextTodoID++
	todos = append(todos, todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		log.Printf("写入 JSON 响应失败: %v", err)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Todo ID 无效", http.StatusBadRequest)
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(todo); err != nil {
				log.Printf("写入 JSON 响应失败: %v", err)
			}
			return
		}
	}

	http.Error(w, "Todo 不存在", http.StatusNotFound)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Todo ID 无效", http.StatusBadRequest)
		return
	}

	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "请求 JSON 无效", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(updatedTodo.Title) == "" {
		http.Error(w, "title 不能为空", http.StatusBadRequest)
		return
	}

	for index := range todos {
		if todos[index].ID == id {
			updatedTodo.ID = id
			todos[index] = updatedTodo

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(updatedTodo); err != nil {
				log.Printf("写入 JSON 响应失败: %v", err)
			}
			return
		}
	}

	http.Error(w, "Todo 不存在", http.StatusNotFound)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Todo ID 无效", http.StatusBadRequest)
		return
	}

	for index := range todos {
		if todos[index].ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Todo 不存在", http.StatusNotFound)
}

func main() {
	http.HandleFunc("GET /hello", helloHandler)
	http.HandleFunc("GET /ping", pingHandler)
	http.HandleFunc("GET /todos", todosHandler)
	http.HandleFunc("POST /todos", createTodoHandler)
	http.HandleFunc("GET /todos/{id}", todoHandler)
	http.HandleFunc("PUT /todos/{id}", updateTodoHandler)
	http.HandleFunc("DELETE /todos/{id}", deleteTodoHandler)

	log.Println("服务已启动：http://localhost:8080/hello")
	log.Println("第七关准备：先 POST 创建练习 Todo，再 DELETE /todos/{id}")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
