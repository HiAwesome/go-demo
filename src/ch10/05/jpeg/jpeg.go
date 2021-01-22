//
// @author moqi
// On 2021/1/22 22:42:53
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	// 如果没有这一行语句，程序依然可以编译和运行，但是它将不能正确识别和解码 png 格式的图像
	// 为了执行 image/png/reader.go 的 init 方法
	_ "image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
