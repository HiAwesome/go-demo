//
// @author moqi
// On 2021/1/24 21:36:45
package main

import (
	"fmt"
	"log"
	"moqi.com/go/ch12/07/params"
	"net/http"
)

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(resp, "Search: %+v\n", data)
}

// ~/Code/go-demo/shell(main ✗) ./search &
// [1] 47010
// ~/Code/go-demo/shell(main ✗) ./fetch 'http://localhost:12345/search'
// Search: {Labels:[] MaxResults:10 Exact:false}
// ~/Code/go-demo/shell(main ✗) ./fetch 'http://localhost:12345/search?l=golang&l=programming'
// Search: {Labels:[golang programming] MaxResults:10 Exact:false}
// ~/Code/go-demo/shell(main ✗) ./fetch 'http://localhost:12345/search?l=golang&l=programming&max=100'
// Search: {Labels:[golang programming] MaxResults:100 Exact:false}
// ~/Code/go-demo/shell(main ✗) ./fetch 'http://localhost:12345/search?x=true&l=golang&l=programming'
// Search: {Labels:[golang programming] MaxResults:10 Exact:true}
// ~/Code/go-demo/shell(main ✗) ./fetch 'http://localhost:12345/search?l=golang&l=programming&x=123'
// x: strconv.ParseBool: parsing "123": invalid syntax
// ~/Code/go-demo/shell(main ✗) ./fetch 'http://localhost:12345/search?l=golang&l=programming&max=lots'
// max: strconv.ParseInt: parsing "lots": invalid syntax
