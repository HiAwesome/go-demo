// Package github provides a Go API for the Github issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

/**
 * 有关 JSON 的 Github 接口练习
 *
 * @author moqi
 * On 2021/1/3 20:50
 */
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// ~/Code/go-demo/ch04/05/github(main ✗) go run github.go repo:golang/go json decoder
// 204 issues:
// #42571     dsnet encoding/json: clarify Decoder.InputOffset semantics
// #33416   bserdar encoding/json: This CL adds Decoder.InternKeys
// #43401  opennota proposal: encoding/csv: add Reader.InputOffset method
// #34647 babolivie encoding/json: fix byte counter increments when using d
// #32779       rsc encoding/json: memoize strings during decode
// #36225     dsnet encoding/json: the Decoder.Decode API lends itself to m
// #40128  rogpeppe proposal: encoding/json: garbage-free reading of tokens
// #31701    lr1980 encoding/json: second decode after error impossible
// #11046     kurin encoding/json: Decoder internally buffers full input
// #40982   Segflow encoding/json: use different error type for unknown fie
// #40127  rogpeppe encoding/json: add Encoder.EncodeToken method
// #29035    jaswdr proposal: encoding/json: add error var to compare  the
// #40983   Segflow encoding/json: return a different error type for unknow
// #41144 alvaroale encoding/json: Unmarshaler breaks DisallowUnknownFields
// #34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
// #5901        rsc encoding/json: allow per-Encoder/per-Decoder registrati
// #29750  jacoelho cmd/vet: stdmethods check gets confused if run on a pac
// #33835     Qhesz encoding/json: unmarshalling null into non-nullable gol
// #28923     mvdan encoding/json: speed up the decoding scanner
// #28189     adnsv encoding/json: confusing errors when unmarshaling custo
// #16212 josharian encoding/json: do all reflect work before decoding
// #6647    btracey x/tools/cmd/godoc: display type kind of each named type
// #34564  mdempsky go/internal/gcimporter: single source of truth for deco
// #14750 cyberphon encoding/json: parser ignores the case of member names
// #43474 davechene encoding/json: Decode does not error on invalid leading
// #33854     Qhesz encoding/json: unmarshal option to treat omitted fields
// #30301     zelch encoding/xml: option to treat unknown fields as an erro
// #22533      ibrt proposal: encoding/json: preserve unknown fields
// #26946    deuill encoding/json: clarify what happens when unmarshaling i
// #28143    arp242 proposal: encoding/json: add "readonly" tag
func main() {

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

}

// SearchIssues queries the Github issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// Chapter 5 presents defer, which makes this simpler.
	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	// 使用基于流式的解码器 json.Decoder
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		_ = resp.Body.Close()
		return nil, err
	}

	_ = resp.Body.Close()
	return &result, nil
}
