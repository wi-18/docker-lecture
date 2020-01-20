package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := 80

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		hostname, _ := os.Hostname()
		_, _ = writer.Write([]byte(hostname))
	})

	fmt.Println(fmt.Sprintf("listetning on port %d", port))
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
