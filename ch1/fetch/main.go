// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//////exercise1.7 函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中， 使用这个函数代替掉例子中的ioutil.ReadAll 来拷贝响应结构体到os.Stdout, 避免申请一个缓冲区来存储。 记得处理io.Copy返回结果中的错误。
		if _, err := io.Copy((os.Stdout), resp); err != nil {
			log.Fatal(err)
		}

		////// Original
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)
	}
}

//!-
