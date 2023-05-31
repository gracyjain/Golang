package main

import (
	"fmt"
	"net/http"
	"time"
)

var tym = time.Now().Format(time.Stamp)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, tym)

}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe("", nil)
}
