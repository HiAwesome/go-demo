package main

import (
	"html/template"
	"log"
	"os"
)

/**
 * 自动处理两种 HTML
 *
 * @author moqi
 * On 2021/1/3 23:42
 */
func main() {

	const temple = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(temple))

	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}

	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

}
