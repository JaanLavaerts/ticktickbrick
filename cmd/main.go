package main

import (
	"fmt"
	"net/http"

	"github.com/JaanLavaerts/ticktickbrick/internal/handlers"
)

func main() {
	http.HandleFunc("/ping", handlers.Ping)

	fmt.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
