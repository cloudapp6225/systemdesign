package main

import (
	"fmt"
	"io"
	"net/http"

	roundrobin "github.com/vk-NEU7/systemdesign/Loadbalancer/RoundRobin"
)

func main() {
	mux := http.NewServeMux()

	go func() {
		startBackendServer(":8081", "server1")
	}()

	go func() {
		startBackendServer(":8082", "server2")
	}()

	fmt.Println("A http server is running on port 80")

	mux.HandleFunc("/", handleRequest)
	http.ListenAndServe(":80", mux)

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	serverList := []string{"server1", "server2"}

	rr := roundrobin.NewRoundRobin()
	serName := rr.RoundRobin(serverList)
	fmt.Println("Redirecting to server: ", serName)
	port := ""
	if serName == "server1" {
		port = ":8081"
	} else {
		port = ":8082"
	}

	// perform curl on the server and return the response
	resp, err := http.Get("http://localhost" + port)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Response from server: %s", string(body))

}

func startBackendServer(addr string, serverName string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Server ", serverName, " is running on ", addr)
		w.WriteHeader(http.StatusOK)
		message := fmt.Sprintf("Server %s is responding from %s", serverName, addr)
		io.WriteString(w, message)
	})

	http.ListenAndServe(addr, mux)
}
