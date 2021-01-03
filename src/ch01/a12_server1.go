package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
 * Web 服务 V1
 *
 * @author moqi
 * On 12/11/20 17:50
 */

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
