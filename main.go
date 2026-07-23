package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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

const dsn = "go_learning:go_learning@tcp(127.0.0.1:3306)/go_learning?parseTime=true"

func openDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

var nextTodoID = 3

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Go First")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "pong")
}

func todosHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.QueryContext(r.Context(),
			"SELECT id, title, done FROM todos ORDER BY id")
		if err != nil {
			http.Error(w, "查询 Todo 失败", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		result := make([]Todo, 0)
		for rows.Next() {
			var todo Todo
			if err := rows.Scan(&todo.ID, &todo.Title, &todo.Done); err != nil {
				http.Error(w, "读取 Todo 失败", http.StatusInternalServerError)
				return
			}
			result = append(result, todo)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, "读取 Todo 失败", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func createTodoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, "请求 JSON 无效", http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(todo.Title) == "" {
			http.Error(w, "title 不能为空", http.StatusBadRequest)
			return
		}

		result, err := db.ExecContext(r.Context(),
			"INSERT INTO todos (title, done) VALUES (?, ?)",
			todo.Title, todo.Done)
		if err != nil {
			http.Error(w, "创建 Todo 失败", http.StatusInternalServerError)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, "读取新 Todo ID 失败", http.StatusInternalServerError)
			return
		}
		todo.ID = int(id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			log.Printf("写入 JSON 响应失败: %v", err)
		}
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
	db, err := openDB()
	if err != nil {
		log.Fatal("连接 MySQL 失败：", err)
	}
	defer db.Close()

	http.HandleFunc("GET /hello", helloHandler)
	http.HandleFunc("GET /ping", pingHandler)
	http.HandleFunc("GET /todos", todosHandler(db))
	http.HandleFunc("POST /todos", createTodoHandler(db))
	http.HandleFunc("GET /todos/{id}", todoHandler)
	http.HandleFunc("PUT /todos/{id}", updateTodoHandler)
	http.HandleFunc("DELETE /todos/{id}", deleteTodoHandler)

	log.Println("服务已启动：http://localhost:8080/hello")
	log.Println("第七关准备：先 POST 创建练习 Todo，再 DELETE /todos/{id}")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
