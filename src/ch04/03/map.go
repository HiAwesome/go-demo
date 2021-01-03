package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

/**
 * Map
 *
 * @author moqi
 * On 2021/1/3 13:20
 */
func main() {

	f1()

	f2()

	// depub()

	// charCount()

	f3()

}

func f3() {
	addEdge("1", "10")
	addEdge("1", "2")
	addEdge("2", "10")
	b1 := hasEdge("1", "2")
	b2 := hasEdge("10", "1")
	fmt.Println(b1, b2)
	fmt.Println()
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

// 统计输入中每个 Unicode 码点出现的次数
func charCount() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utfLen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utfLen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utfLen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

func depub() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "depub: %v\n", err)
		os.Exit(1)
	}
}

func f2() {
	b1 := equals(map[string]int{"A": 0}, map[string]int{"B": 42})
	fmt.Println(b1)
	fmt.Println()
}

// compare two map
func equals(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

func f1() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println("ages = ", ages)
	fmt.Println(ages["alice"])

	delete(ages, "alice")
	fmt.Println(ages)

	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages)
	ages["bob1"] += 1
	ages["bob2"]++
	fmt.Println(ages)

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	fmt.Println()

	// var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	fmt.Println()

	var nilMap map[string]int
	fmt.Println(nilMap == nil)
	fmt.Println(len(nilMap) == 0)
	// panic: assignment to entry in nil map
	// nilMap["carol"] = 21
	fmt.Println(nilMap)

	age, ok := ages["bob"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("\"bob\" is not a key in this map; age == 0")
	}
	fmt.Println()
}
