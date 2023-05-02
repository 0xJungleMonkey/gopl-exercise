// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
// //// 在每个包的包声明前添加注释
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	//////声明变量， 初始化为zero value ""
	var s, sep string
	//////Go语言只有for循环
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

}

//!-
