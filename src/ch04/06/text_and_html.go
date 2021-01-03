package main

import (
	"fmt"
	"html/template"
	"log"
	"moqi.com/go/ch04/05/github"
	"os"
	"time"
)

/**
 * 文本和 HTML 模板
 *
 * @author moqi
 * On 2021/1/3 21:10
 */
const temple = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(temple))

func main() {

	f1()

	f2()

}

// ~/Code/go-demo/src/ch04/06(main ✗) go run text_and_html.go repo:golang/go is:open json decoder
// &{<nil> 0xc000026480 0xc000146200 0xc000012230}
//
// 54 issues:
// ----------------------------------------
// Number: 42571
// User:   dsnet
// Title:  encoding/json: clarify Decoder.InputOffset semantics
// Age:    51 days
// ----------------------------------------
// Number: 33416
// User:   bserdar
// Title:  encoding/json: This CL adds Decoder.InternKeys
// Age:    520 days
func f2() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func f1() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(temple)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(report)
	fmt.Println()
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
