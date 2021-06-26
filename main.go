package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8500"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
