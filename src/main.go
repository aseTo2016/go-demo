package main

import (
	"net/http"
	"os"
	"fmt"
)

func Server(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("ok"))
	return
}

func main() {
	listenServer := os.Getenv("LISTEN_ADDRESS")
	if len(listenServer) == 0 {
		listenServer = "127.0.0.1:8080"
	}
	server := http.NewServeMux()
	server.HandleFunc("/", Server)
	fmt.Printf("start listen server %s", listenServer)
	err := http.ListenAndServe(listenServer, server)
	if err != nil {
		fmt.Printf("listern server failed, %s", err.Error())
	}
}
