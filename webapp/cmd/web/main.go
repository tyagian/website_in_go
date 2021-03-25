package main

import (
	"fmt"
	"net/http"
	"website_in_go/webapp/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	// https://golang.org/pkg/net/http/
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//log.Fatal(http.ListenAndServe("portNumber", nil))
	_ = http.ListenAndServe(portNumber, nil)
}
