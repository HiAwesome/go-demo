package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*
 * Web 服务 V2
 *
 * @author moqi
 * On 12/11/20 19:20
 */

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handlerV2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handlerV2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, _ *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
