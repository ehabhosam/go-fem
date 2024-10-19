package main

import "net/http"

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from my first go server"))
}

func main() {
	server := http.NewServeMux()

	// a handler for a specific route
	server.HandleFunc("/hello", handleHello)

	// server a folder as static files
	fs := http.FileServer(http.Dir("./public"))

	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)

	if err != nil {
		panic(err)
	}
}
