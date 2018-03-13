package main

import (
	"net/http"
	"fmt"
)

func responseHelloWorld(responseWriter http.ResponseWriter, request *http.Request)  {
	request.ParseForm()
	contentLength := request.Header.Get("Content-Length")
	responseWriter.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(responseWriter, `{"value": "%s"}`, contentLength)
}

func main()  {
	fmt.Printf("Server is running ...")
	http.HandleFunc("/hello-world", responseHelloWorld)
	err := http.ListenAndServe("127.0.0.1:8908", nil)
	if err != nil {
		fmt.Printf("error is: %s", err)
	}
}
