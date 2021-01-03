package main

import (
	"encoding/json"
	"fmt"
	"log"
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
