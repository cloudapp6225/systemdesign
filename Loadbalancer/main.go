package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("A http server is running on port 80")
	http.HandleFunc("/", server1HealthCheck)
	http.ListenAndServe(":80", nil)
}

func server1HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Server 1 is healthy")
}
