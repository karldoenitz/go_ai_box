package main

import (
	"net/http"
	"fmt"
	"../handlers"
)

func main()  {
	fmt.Printf("Server is running ...")
	http.HandleFunc("/lean/home", handlers.HomeHandler)
	err := http.ListenAndServe("127.0.0.1:8908", nil)
	if err != nil {
		fmt.Printf("error is: %s", err)
	}
}
