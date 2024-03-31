package main

import (
	"fmt"
	"log"
	"mahai/internal"
	"net/http"
	"time"
)

func main() {
	log.Println("Running on http://localhost:8080")
	mux := http.NewServeMux()
	static := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/", static)
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello, world"))
	})
	start := time.Now()
	_ = internal.NewGame()
	end := time.Since(start)
	fmt.Print(end)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
