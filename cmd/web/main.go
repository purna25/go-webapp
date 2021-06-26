package main

import (
	"fmt"
	"net/http"

	"github.com/purna25/go-webapp/pkg/handlers"
)

const portNumber = ":8500"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
