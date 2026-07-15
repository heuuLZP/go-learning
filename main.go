package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Go")
}

func main() {
	http.HandleFunc("GET /hello", helloHandler)

	log.Println("服务已启动：http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
