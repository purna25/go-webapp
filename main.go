package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World")
		n, err := fmt.Fprintf(w, "<h1>Hello World</h1>")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8500", nil)
}
