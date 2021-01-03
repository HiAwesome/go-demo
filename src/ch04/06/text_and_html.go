package main

import (
	"fmt"
	"html/template"
	"log"
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

	/*result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
	*/
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
