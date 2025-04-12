package main

import (
	"fmt"
	"net/http"
	"github.com/RyonMcAuley/goServer/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := ":8080"
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on %s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
