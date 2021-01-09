package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

/**
 * 处理错误
 * log 包中的所有函数会为没有换行符的字符串增加换行符
 *
 * @author moqi
 * On 2021/1/9 19:16:30
 */
func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "usage: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]

	if err := WaitForServer(url); err != nil {
		// _, _ = fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		// os.Exit(1)
		// 调用 log.Fatalf 可以更简洁的达到上文两行代码相同的效果。
		// log 中的所有函数，都默认会在错误信息之前输出时间信息。
		// 且 log.Fatalf 默认调用 os.Exit(1)
		log.Fatalf("Site is down: %v\n", err)
	}
}

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err != nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
