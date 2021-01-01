package main

import (
	"github-proj/controller"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {

	godotenv.Load(".env")
	http.HandleFunc("/repo", controller.CreateRepo)
	http.ListenAndServe(":8080", nil)
}
