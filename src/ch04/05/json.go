package main

import (
	"encoding/json"
	"fmt"
	"log"
	"moqi.com/go/ch04/05/github"
	"os"
)

/**
 * JSON
 *
 * @author moqi
 * On 2021/1/3 20:31
 */
// 只有导出的结构体成员才会被编码，这也是我们为什么选择用大写字母开头的成员名称
type Movie struct {
	Title string
	// Year 和 Color 成员后面的字符串面值是结构体成员 Tag
	// Tag 表示生成 JSON 时的名字
	Year int `json:"released"`
	// Color 后面的 omitempty 表示当此成员为空或者零值或者否定值时不生成该 JSON 字段
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main() {

	f1()

	f2()

	f3()

	testGithubSearch()

}

// ~/Code/go-demo/ch04/05/github(main ✗) go run json.go repo:golang/go json decoder
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
func testGithubSearch() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func f3() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling filed: %s", err)
	}
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
	fmt.Println()
}

func f2() {
	// 生成容易阅读的形式
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling filed: %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println()
}

func f1() {
	// 将一个 Go 语言中类似 movies 的结构体 slice 转为 JSON 的过程叫编组（marshaling）
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling filed: %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println()
}
