package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//gets the current directory of the go executable
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}
	return dir
}

func logError(caller string, err error) {
	if err != nil {
		log.Println(caller, err)
	}
}

func response(w http.ResponseWriter, req *http.Request) {
	thing := req.URL.Query().Get("thing")
	file, err := ioutil.ReadFile(getCurrentDirectory() + "/" + thing + ".json")
	logError("response writer", err)
	fmt.Fprintf(w, string(file[:]))
}

func main() {

	http.HandleFunc("/mockapi", response)

	err := http.ListenAndServe(":8080", nil)
	logError("main", err)

}
