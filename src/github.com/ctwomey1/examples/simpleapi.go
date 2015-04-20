package main

import (
	"fmt"
	"net/http"
)

func response(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "EXPLOSIONS!")
}

func main() {
	http.HandleFunc("/mockapi", response)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
