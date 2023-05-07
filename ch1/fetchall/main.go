// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
// goroutine 是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递。
// main函数本身也运行在一个goroutine 中，而go function 则表示创建一个新的 goroutine, 并在这个新的goroutine中执行这个函数。
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	//用make函数创建了一个传递string类型参数的channel
	ch := make(chan string)
	//对每一个命令行参数，
	for _, url := range os.Args[1:] {
		// go这个关键字来创建一个goroutine，//并且让函数在这个goroutine异步执行http.get方法。
		go fetch(url, ch) // start a goroutine
	}
	
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // 会把响应的body内容拷贝到ioutil.Discard 输出流中。
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
