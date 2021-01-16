// +build ignore

// The thumbnail command produces thumbnails of JPEG files
// whose names are provided on each line of the standard input.
//
// The "+build ignore" tag (see p.295) excludes this file from the
// thumbnail package, but it can be compiled as a command and run like
// this:
//
// Run with:
//   $ go run $GOPATH/src/gopl.io/ch8/thumbnail/main.go
//   foo.jpeg
//   ^D
//

// moqiFIXME: 2021/1/16 : 这里 IDEA 会显示
//  main.go doesn't match to target system. File will be ignored by build tool
//  如果按照 https://youtrack.jetbrains.com/issue/GO-4859 中的解决方案
//  在 Preferences -> Languages & Frameworks -> Go -> Build Tags & Vendoring -> Custom tags
//  中新增了 ignore 选项，则会此文件编译报错: Multiple packages in directory.
// 	在没有找到更好的方法之前，目前暂时由它提示，不处理。
// @author moqi
// On 2021/1/16 15:25:56
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"moqi.com/go/ch08/05/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}

