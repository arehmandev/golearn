package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// A static web server serving a current path
	fmt.Println("Server starting on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
