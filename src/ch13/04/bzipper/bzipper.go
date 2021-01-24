//
// @author moqi
// On 2021/1/25 00:14:01
package main

import (
	"io"
	"log"
	bzip "moqi.com/go/ch13/04/bzip2"
	"os"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}

// ~ wc -c < /usr/share/dict/words
// 2493109
// ~ sha256sum < /usr/share/dict/words
// 69d1a74371b98676d3312727115b13fefaf57d561561b6ee347d13b1eea0d0c9  -
// ~/Code/go-demo/src/ch13/04/bzipper(main ✗) go run bzipper.go < /usr/share/dict/words | bunzip2 | sha256sum
// 69d1a74371b98676d3312727115b13fefaf57d561561b6ee347d13b1eea0d0c9  -
