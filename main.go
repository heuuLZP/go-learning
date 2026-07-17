package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Go First")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "pong")
}

func main() {
	http.HandleFunc("GET /hello", helloHandler)
	http.HandleFunc("GET /ping", pingHandler)

	log.Println("服务已启动：http://localhost:8080/hello")
	log.Println("第二关入口：http://localhost:8080/ping")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
